package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Cell struct {
	x     int
	y     int
	value string
}

//Let's display a big X on screen on a 20x20 grid!
//But first, let's initialize a bunch of cells

func main() {

	//cell := Cell{
	//	x:     0,
	//	y:     0,
	//	value: "X",
	//}

	//cells := make([]Cell, 6)
	//cells = populateCells(cells)
	//fmt.Println(cells)
	//print20x20()
	//a := createSlice(5)
	//fmt.Println(a)

	b := createGrid(60, 60) //only accepts x == y, breaks otherwise, fix it later!
	b = populateGrid(b)
	b[0][0] = "a"
	b[1][0] = "b"
	b[2][0] = "c"
	printGrid(b)
	b = moveGridDown(b)
	b = randomizeFirstLine(b)
	printGrid(b)

	//
	for {
		fmt.Print("\033[H\033[2J")
		b = moveGridDown(b)
		b = randomizeFirstLine(b)
		printGrid(b)
		time.Sleep(100 * time.Millisecond)
	}
}

func randomizeFirstLine(g [][]string) [][]string {
	for i := 0; i < len(g[0]); i++ {
		g[0][i] = getRandomString(1)
	}

	return g
}

//YOU STOPPED HERE: https://go.dev/blog/slices-intro
//There's something wrong in this function. Changing the first line changes actually the first two lines.

//Let's just one row for now, you can parametrize later
//If you don't use copy() here and instead do `g[i] = g[i-1]` you run into some shallow copy problems and have 2 indexes pointing to the same thing
func moveGridDown(g [][]string) [][]string {
	for i := len(g) - 1; i > 0; i-- {
		copy(g[i], g[i-1])
	}
	return g
}

func randomizeMatrixStyle(g [][]string) [][]string {
	//move the entire grid down one row
	//generate a row of random characters
	//insert that row in the top of the grid
	return g
}

func randomizeGrid(g [][]string) [][]string {
	for _, s := range g {
		for i := 0; i < len(s); i++ {
			s[i] = getRandomString(1)
		}
	}
	return g
}

func getRandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func populateCells(cell []Cell) []Cell {
	//What does this number even mean? What do you call it?
	num := len(cell) / 2 //But then we can only handle an even ammount of cells? uh-oh we need to handle this
	var cs []Cell
	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			c := Cell{
				x:     i,
				y:     j,
				value: "a",
			}
			cs = append(cs, c)
		}
	}
	return cs
}

func printGrid(g [][]string) {
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			fmt.Print(g[i][j], "  ")
		}
		fmt.Println()
	}
}

//Initialize and return a bi-dimensional slice. Nested slices? I did not expect this to be this unwieldy!
func createGrid(x, y int) [][]string {
	g := make([][]string, x)
	for i := 0; i < y; i++ {
		g[i] = make([]string, y)
	}
	return g
}

func populateGrid(g [][]string) [][]string {
	for _, s := range g {
		for i := 0; i < len(s); i++ {
			s[i] = "x"
		}
	}
	return g
}

func createSlice(x int) []string {
	s := make([]string, x)
	for i := 0; i < x; i++ {
		s[i] = "x"
		i++
	}
	return s
}
