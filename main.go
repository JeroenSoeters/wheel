package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		helpFunc()
		os.Exit(0)
	}

	fmt.Println("WARNING: This is a pretoyype, no functionality included!")
}

func helpFunc() {
	helpText :=
		`Usage: wheel [OPTIONS] COMMAND [arg...]

A tool for managing distributed systems.

Options:
  -d, --dry-run    Dry run is just a simulation, it will not make any changes to your system
  -v, --version    Print version information and quit

Commands:
  init            Create a new wheel system
  services        Mangage services
  environments    List environments
  run	          Run services

Run 'wheel COMMAND --help' for more information on a command.`

	fmt.Println(helpText)
}
