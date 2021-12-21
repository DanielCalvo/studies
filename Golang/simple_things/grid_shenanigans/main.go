package main

import (
	"fmt"
)

//
func main() {
	b := createGrid(10, 10)
	b = populateGrid(b)
	//y  x
	b[6][2] = "Z"
	b[5][2] = "Z"
	b[4][2] = "Z"
	//Let's identify all the neighbours of this cell!
	//printGrid(b)

	//iterating over the grid
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			//fmt.Print(b[i][j], "  ")
			fmt.Println(identifyNeighbours(b, i, j))
			//identifyAliveNeighbours
			//apply the logic for number of neighbours to see what happens with the cell
		}
		//fmt.Println()
	}

	//Iterate over all the grid
	//count the neighbours a given cell has
	//finish

}
func identifyAliveNeighbours() {

}

func ChangeCoordinates(g [][]string, coords []Coordinates, s string) [][]string {
	for _, coord := range coords {
		g[coord.y][coord.x] = s
	}
	return g
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

//Ruh-roh, what if the neighbours are out of bound? You need to get the dimensions of the Grid to avoid out of bounds problems!
//To handle out of bounds: https://stackoverflow.com/questions/57676736/how-to-catch-the-slice-bounds-out-of-range-error-and-write-a-handle-for-it
func identifyNeighbours(g [][]string, y, x int) []Coordinates { //idk what I'm gonna return
	//This ended up looking mega ugly, could probably define those as constants or something
	var coords []Coordinates

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
	return coords
}
