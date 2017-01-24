package wheel

import "testing"

const temmplateFile = "/Users/jeroensoeters/dev/gocode/src/github.com/JeroenSoeters/wheel/templates/single-master.cloudformation.json"

func TestReadTemplate(t *testing.T) {
	template, err := readTemplate(templateFile)

	if err != nil {
		t.Error("Error loading cloudformation template file: %v", err)
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

func TestCreateStack(t *testing.T) {
	if err := CreateStack("us-west-2", "dcos-test", map[string]string{
		"KeyName": "dcos-bootstrap",
	}); err != nil {
		t.Errorf("Error creating stack %v", err)
	}
}
