package wheel

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"io/ioutil"
	"os"
)

const templateFile = "/Users/jeroensoeters/dev/gocode/src/github.com/JeroenSoeters/wheel/templates/single-master.cloudformation.json"

func readTemplate(template string) (string, error) {
	bs, err := ioutil.ReadFile(template)
	if err != nil {
		fmt.Printf("Error opening CloudFormation template: %v", err)
		os.Exit(1)
	}

	return string(bs), nil
}

func CreateStack(region string, name string, parameters map[string]string) error {
	template, err := readTemplate(templateFile)
	if err != nil {
		fmt.Printf("Error loading template: %v", err)
	}

	s := session.New(&aws.Config{Region: aws.String(region)})
	cf := cloudformation.New(s)

	_, err = cf.DescribeStacks(&cloudformation.DescribeStacksInput{
		StackName: aws.String(name),
	})

	if err != nil {
		fmt.Printf("Error describing stack %v", err)
	}

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
