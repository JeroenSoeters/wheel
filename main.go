package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
)

func main() {
	args := os.Args[1:]

	if len(os.Args) == 0 {
		helpFunc()
		os.Exit(2)
	}

	c := cli.NewCLI("wheel", "0.0.1")
	c.Args = args
	c.Commands = Commands

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
