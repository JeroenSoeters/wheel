package wheel

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestInitConfig(t *testing.T) {
	project := "my-test-project"
	err := InitConfig(Config{
		ProjectName: project,
	})
	if err != nil {
		t.Errorf("Error initializing configuration: %v", err)
	}

	f, err := ioutil.ReadFile(".wheel/config")
	if err != nil {
		t.Errorf("Failed to open config file: %v", err)
	}

	if strings.Index(string(f), project) == -1 {
		t.Errorf("Config did not contain project name %s", project)
	}

	clean()
}

func clean() {
	_ = os.RemoveAll(".wheel")
}
