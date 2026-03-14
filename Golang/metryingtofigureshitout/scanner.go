package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	input := "one line\ntwo lines\nthree lines"
	//strings new reader returns a *strings.Reader that implements the io.Reader interface from a string
	//it also implements a few more: [io.ReaderAt], [io.ByteReader], [io.ByteScanner], [io.RuneReader], [io.RuneScanner], [io.Seeker], and [io.WriterTo]
	myReader := strings.NewReader(input)

	//a scanner consumes bytes from an io.Reader
	//it splits that byte stream into tokes -- by default tokens are lines
	//you pull tokens with methods like scan, text and bytes
	myScanner := bufio.NewScanner(myReader)

	//common usage pattern
	for myScanner.Scan() { // scan moves to the next token and returns true while you still have tokens to process
		fmt.Println(myScanner.Text()) //returns the current token as string, but you could alse use bytes() here
	}

	//for any code outside of a study environment, you should always check for error:
	if err := myScanner.Err(); err != nil {
		//the scanner errored out. Maybe it encountered a string that was too big for the buffer?
		//note: if a scan errors it breaks out of the loop and stops scanning, it does not go to the next line!
		fmt.Println(err)
	}

	//you can also change the token you're iterating over to be something else, like a word!
	myScanner2 := bufio.NewScanner(strings.NewReader(input))
	myScanner2.Split(bufio.ScanWords)
	for myScanner2.Scan() { // scan moves to the next token and returns true while you still have tokens to process
		fmt.Print(myScanner2.Text(), " ") //returns the current token as string, but you could alse use bytes() here
	}

}

//question here: what if you're reading from something and waiting for input, like getting bytes from the network, how does that go? does scanner scan just quit and you have to launch a new one?
