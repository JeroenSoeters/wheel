package main

import "testing"

func TestLoadTemplate(t *testing.T) {
	template := OpenTemplate()
	expected := "5"
	actual := template.Parameters["SlaveInstanceCount"].Default
	if actual != expected {
		t.Error("expected", expected)
		t.Error("actual ", actual)
	}
}

func TestDeployTemplate(t *testing.T) {
	err := DeployDCOSCluster()
	if err != nil {
		t.Errorf("Error deploying cloudformation DCOS stack: %v", err)
	}
}
