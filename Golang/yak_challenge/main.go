package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

func SaveChunkSorted(lines []int, filePath string) error {

	//If you have this: {1,6,8,2,4}
	//After this sort you'll have: {8,6,4,2,1}
	sort.Sort(sort.Reverse(sort.IntSlice(lines)))

	f, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return err
	}

	defer f.Close()
	for _, num := range lines {
		_, err = f.WriteString(fmt.Sprintf("%d\n", num))
		if err != nil {
			return err
		}
	}
	return nil
}

type HappySorter []Happylist

func (a HappySorter) Len() int           { return len(a) }
func (a HappySorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a HappySorter) Less(i, j int) bool { return a[i].CurrentValue < a[j].CurrentValue }

type Happylist struct {
	filePath     string
	CurrentValue int
	File         *os.File
	Scanner      *bufio.Scanner
	IsScanning   bool
}

func (m *Happylist) Initiate(filePath string) error {
	var err error
	m.File, err = os.Open(filePath)
	if err != nil {
		return err
	}
	m.Scanner = bufio.NewScanner(m.File)
	m.GetNextValue()

	return nil
}

func (m *Happylist) GetNextValue() error {
	var err error
	m.IsScanning = m.Scanner.Scan()
	if !m.IsScanning {
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

//Thanks wikipedia, very cool:
//https://en.wikipedia.org/wiki/External_sorting
//https://en.wikipedia.org/wiki/Merge_algorithm#K-way_merging

//make the upper limit a variable
//use int64s for the argument and all else! Your generator has an int64...
//I think I should put Yak references on the code in hopes they think I'm cool...

func main() {

	filePath := flag.String("filepath", "", "Filesystem path to an unsorted list")
	num := flag.Int("num", 10, "The amount of results to display")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("You must provide a filepath to an unsorted list")
		fmt.Println("Ex: go run main.go -filepath=/tmp/unsorted_list.txt")
		os.Exit(1)
	}

	if *num <= 0 {
		fmt.Println("ERROR: Number of top results must be bigger than 0")
		os.Exit(1)
	}

	if *num >= 30000000 {
		fmt.Println("ERROR: Maximum number of top results must be less or equal than 30000000")
		os.Exit(1)
	}

	file, err := os.Open(*filePath)

	if os.IsNotExist(err) {
		fmt.Println("ERROR: input file does not exist.")
		os.Exit(1)
	}

	if os.IsPermission(err) {
		fmt.Println("ERROR: input file is not readable.")
		os.Exit(1)
	}

	//Maybe we get some other error other than IsNotExist or IsPermission...
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	//General variable set up:
	tmpWorkDir := "/tmp/split_tmp/"
	linesPerTmpFile := 500000

	//Part 1: Sorting
	if _, err := os.Stat(tmpWorkDir); os.IsNotExist(err) {
		os.Mkdir(tmpWorkDir, 0755)
	}

	lineCounter := 0
	scanner := bufio.NewScanner(file)
	var intSlice []int
	fileNameCounter := 1

	for scanner.Scan() {
		lineCounter++
		Int, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err) //I think one of the requirements goes here
		}
		intSlice = append(intSlice, Int)

		if lineCounter == linesPerTmpFile {
			err = SaveChunkSorted(intSlice, tmpWorkDir+strconv.Itoa(fileNameCounter)+".txt")
			if err != nil {
				fmt.Println(err)
			}

			lineCounter = 0
			intSlice = nil
			fileNameCounter++
		}
	}
	// This will almost never reach lineCounter == linesPerTmpFile on the last loop iteration, so we save the leftovers:
	err = SaveChunkSorted(intSlice, tmpWorkDir+strconv.Itoa(fileNameCounter)+".txt")
	if err != nil {
		fmt.Println(err)
	}

	//Part 2: Merging
	var happylists []Happylist

	files, err := ioutil.ReadDir(tmpWorkDir)
	if err != nil {
		panic(err)
	}
	//throw error if no files to resize?
	for _, v := range files {
		h := Happylist{}
		h.Initiate(tmpWorkDir + v.Name())
		happylists = append(happylists, h)
	}

	f, err := os.Create("/tmp/sorted.txt")
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}
	defer f.Close()

	for {
		sort.Sort(sort.Reverse(HappySorter(happylists)))
		_, err = f.WriteString(fmt.Sprintf("%d\n", happylists[0].CurrentValue))

		happylists[0].GetNextValue()

		//if a list has no more elements to scan, we remove it from our list of lists
		if happylists[0].IsScanning == false {
			_, happylists = happylists[0], happylists[1:]
		}

		if len(happylists) == 0 {
			break
		}
	}
	//Hey, don't forget cleanup!
}
