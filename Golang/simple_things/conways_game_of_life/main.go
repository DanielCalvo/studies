package main

import (
	"fmt"
	"math/rand"
	"time"
)

// I need to get rid of this struct!
type Coordinates struct {
	x int
	y int
}

func main() {
	b := createGrid(50)
	b = PopulateGrid(b)

	for {
		printGrid(b)
		b = ApplyRules(b)
		time.Sleep(time.Millisecond * 100)
		fmt.Print("\033[H\033[2J")
	}

}

// Populates cells in a grid with a 50/50 chance of being either alive or dead
func PopulateGrid(g [][]string) [][]string {
	for _, s := range g {
		for i := 0; i < len(s); i++ {
			s[i] = GetRandomString(1)
		}
	}
	return g
}

func GetRandomString(n int) string {
	var letters = []rune(" *") //x == dead, Z == alive
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func ApplyRules(g [][]string) [][]string { //Hmm, maybe this returns the next turn of the game?
	var alive bool
	var liveNeighbourCounter int
	newGrid := createGrid(len(g))

	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {
			liveNeighbourCounter = 0
			if g[y][x] == "*" {
				alive = true
			} else {
				alive = false
			}

			neighbours := identifyNeighbours(g, y, x)

			for _, neighbour := range neighbours {
				if g[neighbour.y][neighbour.x] == "*" {
					liveNeighbourCounter++
				}
			}

			if alive == true && (liveNeighbourCounter == 2 || liveNeighbourCounter == 3) { //This cell is alive and remains alive!
				newGrid[y][x] = "*"
			} else if alive == true { //This is alive, but dies :(
				newGrid[y][x] = " "
			}
			if alive == false && liveNeighbourCounter == 3 { //This cell is dead, but comes to life!
				newGrid[y][x] = "*"
			} else if alive == false { //This cell is dead and remains dead
				newGrid[y][x] = " "
			}
		}
	}
	return newGrid
}

func isOutofBoundsMatrix(s [][]string, c Coordinates) bool {
	if c.y < 0 || c.y > len(s)-1 {
		return true
	}
	//For now I'm assuming all slices inside the slice have equal length (this matrix is either a square or a rectangle)
	if c.x < 0 || c.x > len(s[0])-1 {
		return true
	}
	return false
}

func identifyNeighbours(g [][]string, y, x int) []Coordinates {
	//This ended up looking mega ugly, could probably define those as constants or something
	var coords []Coordinates

	if !isOutofBoundsMatrix(g, Coordinates{x: x - 1, y: y + 1}) { //top left
		coords = append(coords, Coordinates{x: x - 1, y: y + 1})
	}
	if !isOutofBoundsMatrix(g, Coordinates{x: x + 1, y: y + 1}) { //top right
		coords = append(coords, Coordinates{x: x + 1, y: y + 1})
	}
	if !isOutofBoundsMatrix(g, Coordinates{x: x - 1, y: y - 1}) { //bottom left
		coords = append(coords, Coordinates{x: x - 1, y: y - 1})
	}
	if !isOutofBoundsMatrix(g, Coordinates{x: x + 1, y: y - 1}) { //bottom right
		coords = append(coords, Coordinates{x: x + 1, y: y - 1})
	}
	if !isOutofBoundsMatrix(g, Coordinates{x: x - 1, y: y}) { //left
		coords = append(coords, Coordinates{x: x - 1, y: y})
	}
	if !isOutofBoundsMatrix(g, Coordinates{x: x + 1, y: y}) { //right
		coords = append(coords, Coordinates{x: x + 1, y: y})
	}
	if !isOutofBoundsMatrix(g, Coordinates{x: x, y: y + 1}) { //top
		coords = append(coords, Coordinates{x: x, y: y + 1})
	}
	if !isOutofBoundsMatrix(g, Coordinates{x: x, y: y - 1}) { //bottom
		coords = append(coords, Coordinates{x: x, y: y - 1})
	}

	return coords
}

func printGrid(g [][]string) {
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			fmt.Print(g[i][j], "  ")
		}
		fmt.Println()
	}
}

// returns a square grid
func createGrid(size int) [][]string {
	g := make([][]string, size)
	for i := 0; i < size; i++ {
		g[i] = make([]string, size)
	}
	return g
}
