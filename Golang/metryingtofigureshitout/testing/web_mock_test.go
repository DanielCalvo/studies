package main

import "testing"

func myString() string {
	return "Hello world!"
}

func TestMyString(t *testing.T) {
	if myString() != "Hello world!" {
		t.Errorf("Didn't get hello world!")
	}
}
