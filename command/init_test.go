package command

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestInitConfig(t *testing.T) {
	args := []string{
		"-project-name", "training-wheels",
		"-key-pair", "dcos-bootstrap",
	}

	c := &InitCommand{}
	if code := c.Run(args); code != 0 {
		t.Fatalf("Exited with code: %d", code)
	}

	f, err := ioutil.ReadFile(".wheel/config")
	if err != nil {
		t.Errorf("Failed to open config file: %v", err)
	}

	if strings.Index(string(f), "training-wheels") == -1 {
		t.Errorf("Config did not contain project name %s", "training-wheels")
	}

	clean()
}

func clean() {
	_ = os.RemoveAll(".wheel")
}
