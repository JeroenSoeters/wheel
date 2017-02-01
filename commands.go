package main

import (
	"github.com/JeroenSoeters/wheel/command"
	"github.com/mitchellh/cli"
)

var Commands map[string]cli.CommandFactory

func init() {
	Commands = map[string]cli.CommandFactory{
		"init": func() (cli.Command, error) {
			return &command.InitCommand{
				Provider: AwsClient{}
			}, nil
		},
	}
}
