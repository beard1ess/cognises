package awsprovider

import (
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// declare ec2 service instance
var sess = session.Must(session.NewSession())
var svc = ec2.New(sess, aws.NewConfig().WithRegion("us-east-2"))

type functionToIterate func(string) interface{}

func check(err error) {
	if err != nil {
		log.Error(err)
	}
}

func getRegions() []string {
	var regionNames []string
	input := &ec2.DescribeRegionsInput{}
	regions, err := svc.DescribeRegions(input)
	check(err)
	for _, r := range regions.Regions {
		regionNames = append(regionNames, *r.RegionName)
	}
	return regionNames
}

// provide a functionToIterate function and return channel
func iterateRegions(iter functionToIterate) <-chan interface{} {
	var wg sync.WaitGroup
	allRegions := getRegions()
	wg.Add(len(allRegions))
	returns := make(chan interface{}, len(allRegions))
	for i := range allRegions {
		go func(r string, w *sync.WaitGroup) {
			defer w.Done()
			returns <- iter(r)
		}(allRegions[i], &wg)
	}
	wg.Wait()
	close(returns)
	return returns
}

func concatReservations(res []*ec2.Reservation) []*ec2.Instance {
	var ret []*ec2.Instance
	for _, res := range res {
		ret = append(ret, res.Instances...)
	}
	return ret
}

func getInstances(region string) interface{} {
	var instances []*ec2.Instance
	lsvc := ec2.New(sess, aws.NewConfig().WithRegion(region))
	input := &ec2.DescribeInstancesInput{}
	instanceOutput, err := lsvc.DescribeInstances(input)
	check(err)
	instances = append(instances, concatReservations(instanceOutput.Reservations)...)
	pageToken := instanceOutput.NextToken

	for pageToken != nil {
		input := &ec2.DescribeInstancesInput{}
		instanceOutput, err := lsvc.DescribeInstances(input)
		pageToken = instanceOutput.NextToken
		check(err)
		instances = append(instances, concatReservations(instanceOutput.Reservations)...)
	}

	return instances
}

// GetAllEC2Instances return slice of all EC2 instances in all regions
func GetAllEC2Instances() []*ec2.Instance {
	var ret []*ec2.Instance

	for instances := range iterateRegions(getInstances) {
		ret = append(ret, instances.([]*ec2.Instance)...)
	}

	return ret
}
