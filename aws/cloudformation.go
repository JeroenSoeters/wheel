package aws

import (
	"fmt"
	"os"

	"github.com/JeroenSoeters/wheel/templates"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

type CloudFormationService struct {
}

func ReadTemplate() (string, error) {
	data, err := templates.Asset("templates/single-master.cloudformation.json")
	if err != nil {
		fmt.Errorf("Error loading CloudFormation single-master template: %v", err)
		os.Exit(1)
	}

	return string(data), nil
}

func (c CloudFormationService) ProvisionBuildEnvironment(name string) error {
	s := session.New(&aws.Config{Region: aws.String("us-west-2")})
	cf := cloudformation.New(s)

	stack := name + "-ci"
	// Kick off stack creation
	if err := CreateStack(cf, stack, map[string]string{"KeyName": "dcos-bootstrap"}); err != nil {
		return fmt.Errorf("Problem creating stack: %v", err)
	}

	// Wait for stack to be completed
	ew, err := NewStackEventWatcher(cf, stack)
	if err != nil {
		fmt.Printf("Failed to create stack event watcher: %v", err)
	} else {
		fmt.Print("Waiting for stack creation to be created. This can take up to 10 minutes..")
		err = ew.Watch()
	}

	fmt.Print("Build environment provisioned..")

	return err
}

func CreateStack(cf *cloudformation.CloudFormation, name string, parameters map[string]string) error {
	template, err := ReadTemplate()
	if err != nil {
		fmt.Printf("Error loading template: %v", err)
	}

	_, _ = cf.DescribeStacks(&cloudformation.DescribeStacksInput{
		StackName: aws.String(name),
	})

	capabilities := []*string{}
	capabilities = append(capabilities, aws.String(cloudformation.CapabilityCapabilityIam))

	params := []*cloudformation.Parameter{}
	for key, value := range parameters {
		params = append(params, &cloudformation.Parameter{
			ParameterKey:   aws.String(key),
			ParameterValue: aws.String(value),
		})
	}

	_, err = cf.CreateStack(&cloudformation.CreateStackInput{
		StackName:    aws.String(name),
		TemplateBody: &template,
		Capabilities: capabilities,
		Parameters:   params,
	})

	return err
}
