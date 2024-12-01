package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	strarray := strings.Fields(string(data))
	list1 := make([]int, len(strarray)/2)
	list2 := make([]int, len(strarray)/2)
	for i := 0; i < len(strarray); i++ {
		if i%2 == 0 {
			val64, _ := strconv.ParseInt(strarray[i], 10, 32)
			list1[i/2] = int(val64)
		} else {
			val64, _ := strconv.ParseInt(strarray[i], 10, 32)
			list2[i/2] = int(val64)
		}
	}
	sort.Ints(list1)
	sort.Ints(list2)
	var sum = 0
	for i := 0; i < len(list1); i++ {
		sum += int(math.Abs(float64(list1[i] - list2[i])))
	}
	fmt.Println(sum)
}
