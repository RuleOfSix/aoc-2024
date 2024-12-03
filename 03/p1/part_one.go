package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	strdata := string(data)
	var isMore bool = true
	var sum int
	for isMore {
		_, strdata, isMore = strings.Cut(strdata, "mul(")
		if !isMore || unicode.IsSpace(rune(strdata[0])) {
			continue
		}
		var x int
		var y int
		n, err := fmt.Sscanf(strdata, "%d,%d)", &x, &y)
		if err != nil || n < 2 {
			continue
		}
		sum += x * y
	}
	fmt.Println(sum)
}
