package wheel

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"text/template"
)

var project = flag.String("project", "", "name of the project")
var keyPair = flag.String("key-pair", "", "name of the aws key pair")

type Config struct {
	ProjectName string
}

func InitWheel() {
	flag.Parse()
	InitConfig(Config{
		ProjectName: *project,
	})
}

func InitConfig(config Config) error {
	// Wheel config template
	const configTemplate = `
project {
	name = "{{.ProjectName}}"
}
`

	err := os.Mkdir(".wheel", 0777)
	if err != nil {
		fmt.Println("Error creating .wheel directory")
		return err
	}
	fmt.Println("Created .wheel directory")

	f, err := os.Create(".wheel/config")
	if err != nil {
		fmt.Println("Error creating config file")
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	t := template.Must(template.New("config").Parse(configTemplate))
	err = t.Execute(w, config)

	w.Flush()
	fmt.Println("Created wheel config")

	return err
}
