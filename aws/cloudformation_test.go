package aws

import (
	"testing"
)

func TestReadTemplate(t *testing.T) {
	template, err := ReadTemplate()
	if err != nil {
		t.Fatalf("Failed to load template: %v", err)
	}

	expectedStart := "{\"Para"
	start := template[:6]
	if start != expectedStart {
		t.Error("Expected start", expectedStart)
		t.Error("Actual start ", start)
	}

	expectedEnd := "n\"}}}"
	end := template[len(template)-5:]
	if end != expectedEnd {
		t.Error("Expected end", expectedEnd)
		t.Error("Actual end ", end)
	}
}

// Commented out as this test is spinning up a full DC/OS cluster. This should move to a functional test at some point

func TestCreateStack(t *testing.T) {
	if err := CreateStack("us-west-2", "dcos-test", map[string]string{
		"KeyName": "dcos-bootstrap",
	}); err != nil {
		t.Errorf("Error creating stack %v", err)
	}
}
