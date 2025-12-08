package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// The input list of inputs from the previous interaction
var inputs = []string{}

type Range struct {
	Start int
	End   int
}

func main() {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		inputs = append(inputs, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file line-by-line:", err)
	}

	rangeArr, inputArr := parseCommands(inputs)
	rangeList := parseRange(rangeArr)

	partOneSol := partOne(rangeArr, inputArr)

	fmt.Println("part One", partOneSol)

	mergeRangeList := partTwoOptimize(rangeList)
	partTwoSol := CalculateTotalFreshIDs(mergeRangeList)
	fmt.Println("part two", partTwoSol)
}

func parseCommands(rawCommands []string) ([]string, []int) {
	var rangeArr []string
	var inputArr []int
	for _, rawCmd := range rawCommands {

		if rawCmd != "\n" && rawCmd != "" {
			if strings.Contains(rawCmd, "-") {
				rangeArr = append(rangeArr, rawCmd)
			} else {
				intRawValue, err := strconv.Atoi(rawCmd)
				if err != nil {
					fmt.Println("Error converting to int:", err)
					continue
				}

				inputArr = append(inputArr, intRawValue)
			}
		}
	}

	return rangeArr, inputArr
}

func parseRange(rawCommands []string) []Range {
	var results []Range

	for _, itemRange := range rawCommands {
		itemRangeArr := strings.Split(itemRange, "-")
		startingPoint, err := strconv.Atoi(itemRangeArr[0])
		endingPoint, err := strconv.Atoi(itemRangeArr[1])

		if err != nil {
			fmt.Println("Error converting t int", err)
			continue
		}
		results = append(results, Range{
			Start: startingPoint,
			End:   endingPoint,
		})

	}

	return results
}

func partOne(rangeArr []string, inputArr []int) int {
	freshItemCounter := 0

	for _, inp := range inputArr {
		for _, itemRange := range rangeArr {
			itemRangeArr := strings.Split(itemRange, "-")
			startingPoint, err := strconv.Atoi(itemRangeArr[0])
			endingPoint, err := strconv.Atoi(itemRangeArr[1])

			if err != nil {
				fmt.Println("Error converting t int", err)
				continue
			}
			if inp >= startingPoint && inp <= endingPoint {
				freshItemCounter++
				break
			}

		}
	}
	return freshItemCounter
}

func partTwo(rangeArr []string) int {
	freshItemCounter := 0
	freshItemMap := make(map[int]int)

	for _, itemRange := range rangeArr {
		itemRangeArr := strings.Split(itemRange, "-")
		startingPoint, err := strconv.Atoi(itemRangeArr[0])
		endingPoint, err := strconv.Atoi(itemRangeArr[1])

		if err != nil {
			fmt.Println("Error converting t int", err)
			continue
		}

		freshItemMap[startingPoint] = endingPoint

	}
	fmt.Println(freshItemMap)

	return freshItemCounter
}

func partTwoOptimize(ranges []Range) []Range {
	if len(ranges) == 0 {
		return nil
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	var merged []Range
	currentRange := ranges[0]

	for i := 1; i < len(ranges); i++ {
		nextRange := ranges[i]

		if nextRange.Start <= currentRange.End+1 {
			if nextRange.End > currentRange.End {
				currentRange.End = nextRange.End
			}
		} else {
			merged = append(merged, currentRange)
			currentRange = nextRange
		}
	}

	merged = append(merged, currentRange)
	return merged
}

func CalculateTotalFreshIDs(mergedRanges []Range) int {
	total := 0
	for _, r := range mergedRanges {
		total += (r.End - r.Start + 1)
	}
	return total
}
