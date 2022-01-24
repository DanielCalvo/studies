package main

import "testing"

func TestGetYaml(t *testing.T) {
	_, err := GetYaml("./a.yml")
	if err != nil {
		t.Error("Could not GetYaml")
	}
}
