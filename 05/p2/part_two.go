package main

import (
	"fmt"
	"os"
	"slices"
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

	var rules = make([][2]int, len(rulesLines))
	for i := 0; i < len(rulesLines); i++ {
		ruleNums := strings.Split(rulesLines[i], "|")
		for j := 0; j < 2; j++ {
			val64, _ := strconv.ParseInt(ruleNums[j], 10, 32)
			rules[i][j] = int(val64)
		}
	}

	validUpdates := make([]bool, len(updatesLines))
	for i := 0; i < len(validUpdates); i++ {
		validUpdates[i] = true
	}
	sum := 0
	for i := 0; i < len(updatesLines); i++ {
		curUpdateStrs := strings.Split(updatesLines[i], ",")
		curUpdate := make([]int, len(curUpdateStrs))
		for j := 0; j < len(curUpdateStrs); j++ {
			val64, _ := strconv.ParseInt(curUpdateStrs[j], 10, 32)
			curUpdate[j] = int(val64)
		}
		for j := 0; j < len(rules); j++ {
			indexLow := slices.Index(curUpdate, rules[j][0])
			indexHigh := slices.Index(curUpdate, rules[j][1])
			if indexLow != -1 && indexHigh != -1 && indexLow > indexHigh {
				validUpdates[i] = false
			}
		}

		if !validUpdates[i] {
			relevantRules := [][2]int{}
			for j := 0; j < len(rules); j++ {
				if slices.Contains(curUpdate, rules[j][0]) && slices.Contains(curUpdate, rules[j][1]) {
					relevantRules = append(relevantRules, rules[j])
				}
			}

			newUpdate := []int{}
			for len(newUpdate) < len(curUpdate) {
				for j := 0; j < len(curUpdate); j++ {
					if slices.Contains(newUpdate, curUpdate[j]) {
						continue
					}

					isAddable := true
					for k := 0; k < len(relevantRules); k++ {
						if relevantRules[k][1] == curUpdate[j] && !slices.Contains(newUpdate, relevantRules[k][0]) {
							isAddable = false
						}
					}

					if isAddable {
						newUpdate = append(newUpdate, curUpdate[j])
						break
					}
				}
			}
			sum += newUpdate[len(newUpdate)/2]
		}
	}
	fmt.Println(sum)
}
