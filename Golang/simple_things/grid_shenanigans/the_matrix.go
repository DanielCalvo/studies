package main

//
//import (
//	"fmt"
//	"time"
//)
//

//
////Let's display a big X on screen on a 20x20 grid!
////But first, let's initialize a bunch of cells
//
//func main() {
//
//	//cell := Cell{
//	//	x:     0,
//	//	y:     0,
//	//	value: "X",
//	//}
//
//	//cells := make([]Cell, 6)
//	//cells = populateCells(cells)
//	//fmt.Println(cells)
//	//print20x20()
//	//a := createSlice(5)
//	//fmt.Println(a)
//
//	b := createGrid(60, 60) //only accepts x == y, breaks otherwise, fix it later!
//	b = populateGrid(b)
//	b[0][0] = "a"
//	b[1][0] = "b"
//	b[2][0] = "c"
//	printGrid(b)
//	b = moveGridDown(b)
//	b = randomizeFirstLine(b)
//	printGrid(b)
//
//
//	for {
//		fmt.Print("\033[H\033[2J")
//		b = moveGridDown(b)
//		b = randomizeFirstLine(b)
//		printGrid(b)
//		time.Sleep(100 * time.Millisecond)
//	}
//}
