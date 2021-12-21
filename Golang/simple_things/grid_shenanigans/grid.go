package main

import (
	"fmt"
	"math/rand"
)

type Coordinates struct {
	x int
	y int
}

func randomizeFirstLine(g [][]string) [][]string {
	for i := 0; i < len(g[0]); i++ {
		g[0][i] = getRandomString(1)
	}
	return g
}

//Let's just one row for now, you can parametrize later
//If you don't use copy() here and instead do `g[i] = g[i-1]` you run into some shallow copy problems and have 2 indexes pointing to the same thing
func moveGridDown(g [][]string) [][]string {
	for i := len(g) - 1; i > 0; i-- {
		copy(g[i], g[i-1])
	}
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
