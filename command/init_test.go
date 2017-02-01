package command

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

type MockCloudProvider struct {
	BuildEnvironmentProvisioned bool
}

func (m *MockCloudProvider) ProvisionBuildEnvironment() error {
	m.BuildEnvironmentProvisioned = true
	return nil
}

func TestInitConfig(t *testing.T) {
	clean()

	args := []string{
		"-project-name", "training-wheels",
		"-key-pair", "dcos-bootstrap",
	}

	provider := MockCloudProvider{}
	c := &InitCommand{
		Provider: &provider,
	}
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

	if !provider.BuildEnvironmentProvisioned {
		t.Fatal("ProvisionBuildEnvironment was not called")
	}

	clean()
}

func clean() {
	_ = os.RemoveAll(".wheel")
}
