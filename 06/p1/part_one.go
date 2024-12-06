package main

import (
	"fmt"
	"os"
	"strings"
)

type coords struct {
	y int
	x int
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	strmap := strings.Fields(string(data))
	runeMap := make([][]rune, len(strmap))
	for i := 0; i < len(strmap); i++ {
		runeMap[i] = []rune(strmap[i])
	}

	var location coords
	for i := 0; i < len(runeMap); i++ {
		for j := 0; j < len(runeMap[i]); j++ {
			if runeMap[i][j] == '^' {
				location.y = i
				location.x = j
			}
		}
	}

	for inBounds(location.y, location.x, &runeMap) {
		ydir, xdir := getDirection(runeMap[location.y][location.x])
		runeMap[location.y][location.x] = '@'
		for inBounds(location.y+ydir, location.x+xdir, &runeMap) &&
			runeMap[location.y+ydir][location.x+xdir] == '#' {
			ydir, xdir = rotateNinety(ydir, xdir)
		}
		location.y += ydir
		location.x += xdir
		if inBounds(location.y+ydir, location.x+xdir, &runeMap) {
			runeMap[location.y][location.x] = getGuard(ydir, xdir)
		}
	}

	sum := 0
	for i := 0; i < len(runeMap); i++ {
		for j := 0; j < len(runeMap[i]); j++ {
			if runeMap[i][j] == '@' {
				sum++
			}
		}
	}
	fmt.Println(sum)

}

func getDirection(guard rune) (y int, x int) {
	switch guard {
	case '^':
		return -1, 0
	case '>':
		return 0, 1
	case 'v':
		return 1, 0
	case '<':
		return 0, -1
	}
	return 0, 0
}

func getGuard(y int, x int) rune {
	switch {
	case y == -1 && x == 0:
		return '^'
	case y == 0 && x == 1:
		return '>'
	case y == 1 && x == 0:
		return 'v'
	case y == 0 && x == -1:
		return '<'
	}
	return '^'
}

func rotateNinety(y int, x int) (yOut int, xOut int) {
	switch {
	case y == -1 && x == 0:
		return 0, 1
	case y == 0 && x == 1:
		return 1, 0
	case y == 1 && x == 0:
		return 0, -1
	case y == 0 && x == -1:
		return -1, 0
	}
	return 0, 0
}

func inBounds(y int, x int, runeMap *[][]rune) bool {
	rMap := *runeMap
	return y > -1 && y < len(rMap) && x > -1 && x < len(rMap[0])
}
