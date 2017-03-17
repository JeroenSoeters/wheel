package config

import (
	"testing"
)

func Test_LoadValidConfig(t *testing.T) {
	c, err := LoadConfig("test-fixtures/config.wheel")
	if err != nil {
		t.Fatalf("Error loading config: %v", err)
	}
	expected := Config{"test-project", "some-keypair"}
	if c != expected {
		t.Fatalf("Expected: %+v\nActual: %+v", expected, c)
	}
}

func Test_FileNotFound(t *testing.T) {
}
