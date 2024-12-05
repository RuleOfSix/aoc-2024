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
	strdata := string(data)
	slicedata := strings.Split(strdata, "\n\n")
	if len(slicedata) != 2 {
		fmt.Fprintln(os.Stderr, "Error: invalid input format.")
		os.Exit(1)
	}
	rulesLines := strings.Fields(slicedata[0])
	updatesLines := strings.Fields(slicedata[1])

	var rules = make([][2]string, len(rulesLines))
	for i := 0; i < len(rulesLines); i++ {
		ruleNums := strings.Split(rulesLines[i], "|")
		rules[i][0] = ruleNums[0]
		rules[i][1] = ruleNums[1]
	}

	validUpdates := make([]bool, len(updatesLines))
	for i := 0; i < len(validUpdates); i++ {
		validUpdates[i] = true
	}
	sum := 0
	for i := 0; i < len(updatesLines); i++ {
		for j := 0; j < len(rules); j++ {
			indexLow := strings.Index(updatesLines[i], rules[j][0])
			indexHigh := strings.Index(updatesLines[i], rules[j][1])
			if indexLow != -1 && indexHigh != -1 && indexLow > indexHigh {
				validUpdates[i] = false
			}
		}

		if validUpdates[i] {
			updateVals := strings.Split(updatesLines[i], ",")
			mid, _ := strconv.ParseInt(updateVals[len(updateVals)/2], 10, 32)
			sum += int(mid)
		}
	}
	fmt.Println(sum)
}
