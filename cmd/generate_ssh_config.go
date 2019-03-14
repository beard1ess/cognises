package cmd

import (
	"strings"

	aws "../providers/aws"
	s "../ssh_generator"
	"github.com/spf13/cobra"
)

// generateSshConfigCmd represents the generateSshConfig command
var generateSSHConfigCmd = &cobra.Command{
	Use:   "ssh-config",
	Short: "Generate an OpenSSH configuration file from cloud servers.",
	Run: func(cmd *cobra.Command, args []string) {
		var servers []s.Server
		// iterate providers
		for _, p := range c.Providers {
			switch p {
			case "aws":
				servers = append(servers, getAWS()...)
			default:

			}
		}
		s.RenderTemplate(servers)
	},
}

func init() {
	rootCmd.AddCommand(generateSSHConfigCmd)
}

func getAWS() []s.Server {

	var ret []s.Server
	config := c.SSHGenerateConfig

	for _, instance := range aws.GetAllEC2Instances() {
		var i s.Server
		host := ""
		i.Port = config.General.Port
		for _, tag := range instance.Tags {
			if *tag.Key == "Name" {
				if strings.Contains(*tag.Value, " ") {
					host = strings.Replace(*tag.Value, " ", "_", -1)
				} else {
					host = *tag.Value
				}
				if strings.Contains(*tag.Value, ".") {
					host = host + " " + strings.Split(*tag.Value, ".")[0]
				}
			}
		}

		// TODO: Some way to detect user?
		i.User = config.General.User

		i.Host = host + " " + *instance.InstanceId

		if config.PreferPublic {
			if instance.PublicIpAddress != nil {
				i.Hostname = *instance.PublicIpAddress
			} else {
				i.Hostname = *instance.PrivateIpAddress
			}
		} else {
			i.Hostname = *instance.PrivateIpAddress
		}

		if config.General.ProxyCommand != "" {
			i.ProxyCommand = config.General.ProxyCommand
		}
		if config.PrivateKeys.UseInstanceKeypair == true {
			i.Identityfile = config.PrivateKeys.KeyPath + *instance.KeyName + ".pem"
		} else {
			i.Identityfile = config.PrivateKeys.CustomKeypair
		}

		ret = append(ret, i)
	}
	return ret
}
