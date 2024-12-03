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
		mulIndex := strings.Index(strdata, "mul(")
		dontIndex := strings.Index(strdata, "don't()")
		if mulIndex < dontIndex && mulIndex != -1 {
			_, strdata, _ = strings.Cut(strdata, "mul(")
			if len(strdata) == 0 {
				break
			}
			if unicode.IsSpace(rune(strdata[0])) {
				continue
			}
		} else if dontIndex < mulIndex && dontIndex != 1 {
			_, strdata, _ = strings.Cut(strdata, "don't()")
			_, strdata, isMore = strings.Cut(strdata, "do()")
			continue
		} else if dontIndex == -1 && mulIndex == -1 {
			break
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
