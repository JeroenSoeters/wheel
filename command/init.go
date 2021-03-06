package command

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/JeroenSoeters/wheel/config"
	"github.com/JeroenSoeters/wheel/wheel"
)

// InitCommand is a cli.Command implementation that initializes a Wheel project
type InitCommand struct {
	Provider wheel.CloudProvider
}

func (c *InitCommand) Run(args []string) int {
	fs := flag.NewFlagSet("init", flag.ContinueOnError)

	var configFile string
	fs.StringVar(&configFile, "config", "", "Wheel configuration file")

	fs.Parse(args)

	if configFile == "" {
		fmt.Printf("Usage...")
		fmt.Printf(c.Help())
		return 2
	}

	config, err := config.LoadConfig(configFile)

	if err != nil {
		fmt.Printf("Error loading config %v", err)
		return 2
	}

	fmt.Print("Config loaded")

	// Create .wheel folder with configuration
	const configTemplate = `
project {
	name = "{{.ProjectName}}"
}
`

	err = os.Mkdir(".wheel", 0777)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error creating .wheel directory")
		return 1
	}
	fmt.Println("Created .wheel directory")

	f, err := os.Create(".wheel/config")
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error creating config file")
		return 1
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	t := template.Must(template.New("config").Parse(configTemplate))
	err = t.Execute(w, config)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error writing to config file")
		return 1
	}
	w.Flush()
	fmt.Println("Created wheel config")

	// Deploy cloudformatin template
	if err = c.Provider.ProvisionBuildEnvironment(config.ProjectName); err != nil {
		fmt.Fprintf(os.Stdout, "Issue provisioning build environment: %v", err)
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
