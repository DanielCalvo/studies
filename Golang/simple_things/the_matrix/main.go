package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	b := createGrid(50)
	b = PopulateGrid(b)

	for {
		fmt.Print("\033[H\033[2J")
		b = moveGridDown(b)
		b = randomizeFirstLine(b)
		printGrid(b)
		time.Sleep(100 * time.Millisecond)
	}
}

func printGrid(g [][]string) {
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			fmt.Print(g[i][j], "  ")
		}
		fmt.Println()
	}
}

func createGrid(size int) [][]string {
	g := make([][]string, size)
	for i := 0; i < size; i++ {
		g[i] = make([]string, size)
	}
	return g
}

func PopulateGrid(g [][]string) [][]string {
	for _, s := range g {
		for i := 0; i < len(s); i++ {
			s[i] = GetRandomString(1)
		}
	}
	return g
}

func GetRandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

// If you don't use copy() here and instead do `g[i] = g[i-1]` you run into some shallow copy problems and have 2 indexes pointing to the same thing
func moveGridDown(g [][]string) [][]string {
	for i := len(g) - 1; i > 0; i-- {
		copy(g[i], g[i-1])
	}
	return g
}

func randomizeFirstLine(g [][]string) [][]string {
	for i := 0; i < len(g[0]); i++ {
		g[0][i] = GetRandomString(1)
	}
	return g
}
