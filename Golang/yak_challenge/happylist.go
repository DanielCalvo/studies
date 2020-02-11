package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strconv"
)

//remember to return errors, not panic inside functions

type Mylist struct {
	filePath     string
	CurrentValue int
	File         *os.File
	Scanner      *bufio.Scanner
}

func (m *Mylist) Initiate(filePath string) error {
	var err error
	m.File, err = os.Open(filePath)
	if err != nil {
		return err
	}
	m.Scanner = bufio.NewScanner(m.File)
	m.GetNextValue()

	return nil
}

func (m *Mylist) GetNextValue() error {
	var err error
	boolsito := m.Scanner.Scan()
	if !boolsito {
		err = errors.New("EOF")
		m.CurrentValue = -1
		return err
	}
	m.CurrentValue, err = strconv.Atoi(m.Scanner.Text())
	if err != nil {
		return err
		fmt.Println("Error", err)
	}
	return nil
}

func main() {
	happylist := Mylist{}
	happylist.Initiate("/tmp/nums/11.txt")
	fmt.Println(happylist.CurrentValue)
	happylist.GetNextValue()
	fmt.Println(happylist.CurrentValue)
	happylist.GetNextValue()
	fmt.Println(happylist.CurrentValue)

}
