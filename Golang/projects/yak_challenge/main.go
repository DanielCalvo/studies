package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

//A happylist(TM) is a large list, arranged in decreasing order, stored on a file.
//After initiating it with a file path with happylist.Initiate(), you can:
//See it's current value (line) on happylist.CurrentValue
//Advance to it's next value (line) with happylist.GetNextValue()
//See if you have reached the end of this happylist with happylist.IsScanning
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
	err = m.GetNextValue()
	if err != nil {
		return err
	}
	return nil
}

func (m *Happylist) GetNextValue() error {
	var err error
	m.IsScanning = m.Scanner.Scan()
	if m.IsScanning == false { //Scanner.Scan() returns false on end of file
		return nil
	}

	m.CurrentValue, err = strconv.Atoi(m.Scanner.Text())
	if err != nil {
		return err
	}
	return nil
}

//The sort interface is implemented for a slice of happylists.
//This allows us to sort a slice of happylists by CurrentValue
type HappySorter []Happylist

func (a HappySorter) Len() int           { return len(a) }
func (a HappySorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a HappySorter) Less(i, j int) bool { return a[i].CurrentValue < a[j].CurrentValue }

//SaveChunkSorted is used on the first of the problem to save a large unordered list into smaller, ordered lists
func SaveChunkSorted(lines []int, filePath string) error {

	//If you have this: {1,6,8,2,4}
	//After this sort you'll have: {8,6,4,2,1}
	sort.Sort(sort.Reverse(sort.IntSlice(lines)))

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer f.Close()
	for _, num := range lines {
		_, err = f.WriteString(fmt.Sprintf("%d\n", num)) //Free newline on the last element!
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {

	filePath := flag.String("filepath", "", "Filesystem path to an unsorted list")
	resultNum := flag.Int("num", 10, "The amount of results to display")
	tmpWorkDir := flag.String("tmpworkdir", "/tmp/lists/", "Directory in which to store temporary files")
	linesPerFile := flag.Int("linesPerFile", 500000, "Amount of lines per temporary file")
	sortedFilePath := flag.String("sortedfilepath", "/tmp/sorted.txt", "Filesystem path to an unsorted list")

	flag.Parse()

	if *filePath == "" {
		fmt.Println("You must provide at least a filepath to an unsorted list")
		fmt.Println("Ex: go run simple.go -filepath=/tmp/unsorted_list.txt")
		os.Exit(1)
	}

	if *resultNum <= 0 {
		fmt.Println("ERROR: Number of top results must be bigger than 0")
		os.Exit(1)
	}

	if *resultNum > 30000000 {
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

	//Quietly attempt to remove files from previous execution if they were the same as this one
	_ = os.RemoveAll(*tmpWorkDir)
	_ = os.Remove(*sortedFilePath)

	//Part 1: Sorting
	_, err = os.Stat(*tmpWorkDir)

	if os.IsNotExist(err) {
		err = os.Mkdir(*tmpWorkDir, 0755)
		if err != nil {
			fmt.Println("Could not create temporary directory at:", *tmpWorkDir)
		}
	}

	lineCounter := 0
	fileCounter := 0
	totalLineCounter := 0
	scanner := bufio.NewScanner(file)
	var intSlice []int

	for scanner.Scan() {
		totalLineCounter++
		lineCounter++
		Int, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("WARN: Invalid line", totalLineCounter)
			continue
		}
		intSlice = append(intSlice, Int)

		if lineCounter == *linesPerFile {
			err = SaveChunkSorted(intSlice, *tmpWorkDir+strconv.Itoa(fileCounter)+".txt")
			if err != nil {
				fmt.Println("Could not save temporary file at:", *tmpWorkDir+strconv.Itoa(fileCounter)+".txt")
				panic(err)
			}

			lineCounter = 0
			intSlice = nil
			fileCounter++
		}
	}
	//This will almost never reach lineCounter == *linesPerFile on the last loop iteration, so we save the leftovers
	//There could also be an `if lineCounter > 0` or similar logic to avoid creating empty files, but that introduced a bug
	//And I have to get back to my actual job...
	err = SaveChunkSorted(intSlice, *tmpWorkDir+strconv.Itoa(fileCounter)+".txt")
	if err != nil {
		fmt.Println("Could not save temporary file at:", *tmpWorkDir+strconv.Itoa(fileCounter)+".txt")
		panic(err)
	}

	//Part 2: Merging
	//Wikipedia explains this well: https://en.wikipedia.org/wiki/Merge_algorithm#K-way_merging

	//A happylist(TM) is a sorted list in decreasing order with additional functionality (see top of this file)
	var happylists []Happylist

	files, err := ioutil.ReadDir(*tmpWorkDir)
	if err != nil {
		fmt.Println(err)
	}

	//Populate a slice of happylists with initiated happylists
	for _, v := range files {
		h := Happylist{}
		err = h.Initiate(*tmpWorkDir + v.Name())
		if err != nil {
			fmt.Println("Error initiating happylist with argument:", *tmpWorkDir+v.Name())
			panic(err)
		}
		happylists = append(happylists, h)
	}

	//Create the file that will contain the final sorted list:
	finalFile, err := os.Create(*sortedFilePath)
	if err != nil {
		fmt.Println("Could not create file with final sorted list at", *sortedFilePath)
		panic(err)
	}
	defer finalFile.Close()

	//Iterate over our sorted lists until we reach the desired number of results or run out of elements in the happylists(TM)
	results := 0
	for len(happylists) > 0 && results < *resultNum {

		//Sort our slice of sorted lists, the element with the highest CurrentValue will be first
		sort.Sort(sort.Reverse(HappySorter(happylists)))

		//Write the current value of this element to the final sorted list
		_, err = finalFile.WriteString(fmt.Sprintf("%d\n", happylists[0].CurrentValue))
		results++

		//Advance CurrentValue to the next element of that sorted list
		err = happylists[0].GetNextValue()
		if err != nil {
			fmt.Println("Error converting to Integer on", happylists[0].File.Name(), "or got non EOF value from file:", happylists[0].CurrentValue)
			panic(err)
		}

		//If there is no next CurrentValue on that sorted list (scanner at EOF) we have reached its end.
		//remove it from the slice of happylists so we no longer iterate over it
		if happylists[0].IsScanning == false {
			_, happylists = happylists[0], happylists[1:]
		}
	}

	//Part 3: Displaying results
	_, err = finalFile.Seek(0, 0) //Return to the beginning of the file with the final list
	if err != nil {
		fmt.Println("Could not return to the beginning of the file with the final sorted list")
		panic(err)
	}

	//If you had 100 unsorted elements and asked for the top 5, you'll get the top 5.
	//If you had 10 unsorted elements and asked for the top 20, you'll get the top 10, there were no more elements :(
	finalScanner := bufio.NewScanner(finalFile)

	for finalScanner.Scan() {
		fmt.Println(finalScanner.Text())
	}
}
