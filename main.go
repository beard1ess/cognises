package main

import (
	"./cmd"

	"github.com/sirupsen/logrus"
)

type template interface{}

var log = logrus.New()

func main() {
	cmd.Execute()
}
