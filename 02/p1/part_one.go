package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	reports := strings.Split(string(data), "\n")
	sum := 0
	for i := 0; i < len(reports); i++ {
		safe := true
		curReport := strings.Split(reports[i], " ")
		intReport := make([]int64, len(curReport))
		for j := 0; j < len(curReport); j++ {
			intReport[j], _ = strconv.ParseInt(curReport[j], 10, 64)
		}
		var ascending bool
		var descending bool
		if len(intReport) > 1 {
			ascending = intReport[0] < intReport[1]
			descending = !ascending
		}
		for j := 0; j < len(intReport)-1; j++ {
			ascending = ascending && intReport[j] < intReport[j+1]
			descending = descending && intReport[j] > intReport[j+1]
			if math.Abs(float64(intReport[j]-intReport[j+1])) > 3 {
				safe = false
				break
			}
		}
		if !descending && !ascending {
			safe = false
		}
		if safe {
			sum++
		}
	}
	fmt.Println(sum)
}
