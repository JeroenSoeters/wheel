package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cf "github.com/crewjam/go-cloudformation"
	dcf "github.com/crewjam/go-cloudformation/deploycfn"
)

func OpenTemplate() *cf.Template {
	file, err := os.Open("/Users/jeroensoeters/dev/gocode/src/github.com/JeroenSoeters/wheel/templates/single-master.cloudformation.json")
	if err != nil {
		fmt.Printf("Error opening CloudFormation template: %v", err)
		os.Exit(1)
	}
	t := cf.Template{}
	json.NewDecoder(bufio.NewReader(file)).Decode(&t)

	return &t
}

func DeployDCOSCluster() error {
	template := OpenTemplate()
	s := session.New(&aws.Config{Region: aws.String("us-west-2")})
	params := map[string]string{
		"KeyName": "dcos-bootstrap",
	}
	input := dcf.DeployInput{s, "dcos-test-cluster", template, params, ""}

	return dcf.Deploy(input)
}
