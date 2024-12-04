package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	wordSearch := strings.Split(string(data), "\n")
	wordSearch = wordSearch[:len(wordSearch)-1]

	var sum int
	for i := 0; i < len(wordSearch); i++ {
		for j := 0; j < len(wordSearch[i]); j++ {
			if wordSearch[i][j] == 'X' {
				sum += countXmases(wordSearch, i, j)
			}
		}
	}
	fmt.Println(sum)
}

func countXmases(wordSearch []string, i int, j int) int {
	var numXmases int
	if len(wordSearch[i]) > j+3 {
		if wordSearch[i][j+1] == 'M' &&
			wordSearch[i][j+2] == 'A' &&
			wordSearch[i][j+3] == 'S' {
			numXmases++
		}
		if len(wordSearch) > i+3 &&
			wordSearch[i+1][j+1] == 'M' &&
			wordSearch[i+2][j+2] == 'A' &&
			wordSearch[i+3][j+3] == 'S' {
			numXmases++
		}
		if i > 2 &&
			wordSearch[i-1][j+1] == 'M' &&
			wordSearch[i-2][j+2] == 'A' &&
			wordSearch[i-3][j+3] == 'S' {
			numXmases++
		}
	}
	if j > 2 {
		if wordSearch[i][j-1] == 'M' &&
			wordSearch[i][j-2] == 'A' &&
			wordSearch[i][j-3] == 'S' {
			numXmases++
		}
		if len(wordSearch) > i+3 &&
			wordSearch[i+1][j-1] == 'M' &&
			wordSearch[i+2][j-2] == 'A' &&
			wordSearch[i+3][j-3] == 'S' {
			numXmases++
		}
		if i > 2 &&
			wordSearch[i-1][j-1] == 'M' &&
			wordSearch[i-2][j-2] == 'A' &&
			wordSearch[i-3][j-3] == 'S' {
			numXmases++
		}
	}
	if i > 2 &&
		wordSearch[i-1][j] == 'M' &&
		wordSearch[i-2][j] == 'A' &&
		wordSearch[i-3][j] == 'S' {
		numXmases++
	}
	if len(wordSearch) > i+3 &&
		wordSearch[i+1][j] == 'M' &&
		wordSearch[i+2][j] == 'A' &&
		wordSearch[i+3][j] == 'S' {
		numXmases++
	}
	return numXmases
}
