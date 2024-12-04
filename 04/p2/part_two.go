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
			if wordSearch[i][j] == 'A' && isCrossMas(wordSearch, i, j) {
				sum++
			}
		}
	}
	fmt.Println(sum)
}

func isCrossMas(wordSearch []string, i int, j int) bool {
	if len(wordSearch[i]) <= j+1 ||
		j < 1 ||
		len(wordSearch) <= i+1 ||
		i < 1 {
		return false
	}

	if (wordSearch[i-1][j-1] == 'M' && wordSearch[i+1][j+1] == 'S') ||
		(wordSearch[i-1][j-1] == 'S' && wordSearch[i+1][j+1] == 'M') {
		if (wordSearch[i-1][j+1] == 'M' && wordSearch[i+1][j-1] == 'S') ||
			(wordSearch[i-1][j+1] == 'S' && wordSearch[i+1][j-1] == 'M') {
			return true
		}
	}

	return false
}
