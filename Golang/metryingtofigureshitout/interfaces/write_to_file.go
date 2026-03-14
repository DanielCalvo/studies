package main

import (
	"io"
	"os"
	"path/filepath"
)

// right lets do this
// when i run my program normally, i want to write lines to a file
// but when i run tests, i want to write to an in memory writer, so we "fake" a file, so to speak
// in our imaginary world files are expensive
// so in this way we explore the concept of how we would also fake connecting to a db, for example, and i want to use an interface for this!
// but in this case i suspect you dont need to create a new interface, just using the writer one would be sufficient
func main() {
	text := "weeeee"
	//keeping the file path logic outside of the function makes it much easier to test (splitting set up from behaviour makes code easier to test)
	//in this program, file creation is set up, writing the bytes is the behaviour
	path := filepath.Join(os.TempDir(), "file.txt")
	f, err := fileWiring(path)
	if err != nil {
		panic(err)
	}
	//only defer closing after checking the error, trying to call f.Close on a nil file would panic!
	defer f.Close() //the code that owns/acquires the resouce is usually responsible for closing it
	err = write(text, f)
	if err != nil {
		panic(err)
	}
}

// using os.WriteFile here was an idea, but it hides the interface (writer) for testing, so that would've forced me to test writing to a file and i wouldnt have been able to test the writer as well
func fileWiring(path string) (*os.File, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// so lets do the write to file first
func write(s string, w io.Writer) error {
	b := []byte(s)
	_, err := w.Write(b)
	return err
}
