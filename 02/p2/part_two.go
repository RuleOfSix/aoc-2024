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
		curReport := strings.Split(reports[i], " ")
		intReport := make([]int64, len(curReport))
		for j := 0; j < len(curReport); j++ {
			intReport[j], _ = strconv.ParseInt(curReport[j], 10, 64)
		}
		overall_safe := false
		for j := -1; j < len(intReport); j++ {
			safe := true
			subReport := remove(intReport, j)
			if i == 0 {
				fmt.Println(subReport)
			}
			var ascending bool
			var descending bool
			if len(subReport) > 1 {
				ascending = subReport[0] < subReport[1]
				descending = !ascending
			}
			for k := 0; k < len(subReport)-1; k++ {
				ascending = ascending && subReport[k] < subReport[k+1]
				descending = descending && subReport[k] > subReport[k+1]
				if math.Abs(float64(subReport[k]-subReport[k+1])) > 3 {
					safe = false
					break
				}
				if subReport[k] == subReport[k+1] {
					safe = false
					break
				}
			}

			if !descending && !ascending {
				safe = false
			}
			overall_safe = overall_safe || safe
		}
		if i == 0 {
			fmt.Println(overall_safe)
		}
		if overall_safe {
			sum++
		}
	}
	fmt.Println(sum)
}

func remove(slice []int64, index int) []int64 {
	if index == -1 {
		return slice
	}
	output := make([]int64, 0)
	output = append(output, slice[:index]...)
	output = append(output, slice[index+1:]...)
	return output
}
