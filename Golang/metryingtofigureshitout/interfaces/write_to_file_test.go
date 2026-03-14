package main

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

type errWriter struct {
	err error
}

//you dont use table tests here, they become useful when many test cases follow the same structure with different inputs!

// for interface based code, you can fake behaviour by creating a type that implements the interface
func (e errWriter) Write(p []byte) (int, error) {
	return 0, e.err //so it always fails, how handy to test errors!
}

func TestWriter(t *testing.T) {
	var buf bytes.Buffer //bytes.Buffer also satisfies io.Writer, so its useful for testing writer based code!
	s := "hello world"
	err := write(s, &buf)
	if err != nil {
		t.Fatalf("write could not write to buffer, got: %v", err)
	}
	if s != buf.String() {
		t.Errorf("expected buffer to contain %v, got: %v", s, buf.String())
	}
}

// error path tests check that failures are properly returned!
func TestWriterError(t *testing.T) {
	erroredWriter := errWriter{err: errors.New("fail")}
	err := write("hi", erroredWriter)
	if !errors.Is(err, erroredWriter.err) { //errors.Is is nice to compare errors in test!
		t.Errorf("expected write to error with %v, got %v", erroredWriter.err, err)
	}
}

// not every dependency should be faked, if you have code that creates files, test creating files then (aka call the os functions, dont fake the test)
func TestFileWiring(t *testing.T) {
	path := filepath.Join(t.TempDir(), "file.txt") //t.TempDir gives each test an isolated and unique temp dir for testing (avoiding collisions with other tests) and it cleans up that dir after the test is done -- this is a much better idea than using os.Tempdir() which does not have these features!
	f, err := fileWiring(path)
	if err != nil {
		/*
			on test error messages:
			- the go test runner tells you which test failed, so you dont need to put that on the error message
			- saying which tested function failed with which input seems to be a good starting point!
		*/
		t.Fatalf("fileWiring(%q) returned error: %v", path, err)
	}
	defer f.Close()
	_, err = os.Stat(path) //codex suggests: a testing principle to follow is to use the least powerful operation that proves the thing you care about (when asked if I should use os.Openfile() here)
	if err != nil {
		t.Fatalf("Expected file to have been created by fileWiring at %q, got error: %v", path, err)
	}
}

func TestFileWiringError(t *testing.T) {
	//in this case we need to test an invalid file creation,
	path := filepath.Join(t.TempDir(), "path-that-doesnt-exist", "file.txt")
	_, err := fileWiring(path)
	if err == nil {
		t.Errorf("expected fileWiring to error creating file on directory that did not exist")
	}
}
