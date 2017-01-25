package command

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/JeroenSoeters/wheel/aws"
)

// InitCommand is a cli.Command implementation that initializes a Wheel project
type InitCommand struct {
}

// Configuration for Wheel
type Config struct {
	ProjectName string
	KeyPair     string
}

func (c *InitCommand) Run(args []string) int {
	fs := flag.NewFlagSet("init", flag.ContinueOnError)

	config := Config{}

	fs.StringVar(&config.ProjectName, "project-name", "", "project name")
	fs.StringVar(&config.KeyPair, "key-pair", "", "AWS key pair")

	fs.Parse(args)

	if config.ProjectName == "" || config.KeyPair == "" {
		fmt.Printf("Usage...")
		return 2
	}

	const configTemplate = `
project {
	name = "{{.ProjectName}}"
}
`

	err := os.Mkdir(".wheel", 0777)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating .wheel directory")
		return 1
	}
	fmt.Println("Created .wheel directory")

	f, err := os.Create(".wheel/config")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating config file")
		return 1
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	t := template.Must(template.New("config").Parse(configTemplate))
	err = t.Execute(w, config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to config file")
		return 1
	}
	w.Flush()
	fmt.Println("Created wheel config")

	if err := aws.CreateStack("us-west-2", "dcos-test", map[string]string{
		"KeyName": config.KeyPair,
	}); err != nil {
		fmt.Errorf("Error creating stack %v", err)
		return 1
	}

	return 0
}

func (c *InitCommand) Help() string {
	helpText := `
Usage: wheel init [options]	

Options:
	
	-project-name 	Name of the project
	-key-pair		AWS key pair to use
`

	return strings.TrimSpace(helpText)
}

func (c *InitCommand) Synopsis() string {
	return "Initializes Wheel project in the current directory"
}
