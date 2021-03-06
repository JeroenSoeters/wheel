package main

import "strings"

func helpFunc() string {
	helpText :=
		`Usage: wheel [OPTIONS] COMMAND [arg...]

A tool for managing distributed systems.

Options:
  -d, --dry-run    Dry run is just a simulation, it will not make any changes to your system
  -v, --version    Print version information and quit

Commands:
  describe		  Describe this wheel system
  init            Create a new wheel system
  services        Mangage services
  environments    List environments
  run	          Run services

Run 'wheel COMMAND --help' for more information on a command.`

	return strings.TrimSpace(helpText)
}
