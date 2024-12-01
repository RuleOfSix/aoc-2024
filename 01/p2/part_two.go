package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	strdata := strings.Fields(string(data))
	slice1 := make([]int, len(strdata)/2)
	slice2 := make([]int, len(strdata)/2)
	for i := 0; i < len(strdata); i++ {
		if i%2 == 0 {
			val64, _ := strconv.ParseInt(strdata[i], 10, 32)
			slice1[i/2] = int(val64)
		} else {
			val64, _ := strconv.ParseInt(strdata[i], 10, 32)
			slice2[i/2] = int(val64)
		}
	}
	var sum int64 = 0
	for i := 0; i < len(slice1); i++ {
		var count int64 = 0
		for j := 0; j < len(slice2); j++ {
			if slice1[i] == slice2[j] {
				count++
			}
		}
		sum += int64(slice1[i]) * count
	}
	fmt.Println(sum)
}
