package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// The input list of inputs from the previous interaction
var inputs = []string{}

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

	parseCommands := parseCommands(inputs)

	partOneSol := partOne(parseCommands)

	fmt.Println(" Part 1 solution ", partOneSol)

	partTwoSol := partTwo(parseCommands)

	fmt.Println(" Part 2 solution ", partTwoSol)

}

func parseCommands(rawCommands []string) []string {
	var results []string

	for _, rawCmd := range rawCommands {
		subStr := strings.Split(rawCmd, ",")

		for _, str := range subStr {

			if len(str) > 0 {
				results = append(results, str)
			}

		}
	}

	return results
}

func partOne(parsedCommands []string) (finalZeroCounter int) {
	result := 0

	for _, cmd := range parsedCommands {
		args := strings.Split(cmd, "-")
		startingPoint, err := strconv.Atoi(args[0])
		endingPoint, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error converting to int:", err)
			continue
		}

		for i := startingPoint; i <= endingPoint; i++ {
			nts := strconv.Itoa(i)
			length := len(nts)

			if len(nts)%2 == 0 {
				midpoint := length / 2

				// Slice the byte string directly
				firstHalf := nts[:midpoint]
				secondHalf := nts[midpoint:]

				if firstHalf == secondHalf {
					result += i
				}
			}
		}

	}

	return result
}

func partTwo(parsedCommands []string) (finalZeroCounter int) {
	result := 0

	for _, cmd := range parsedCommands {
		args := strings.Split(cmd, "-")
		startingPoint, err := strconv.Atoi(args[0])
		endingPoint, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error converting to int:", err)
			continue
		}

		for i := startingPoint; i <= endingPoint; i++ {
			nts := strconv.Itoa(i)
			lenNts := len(nts)

			for pattern := 1; pattern <= lenNts/2; pattern++ {
				if lenNts%pattern != 0 {
					continue
				}
				repeated := true
				patternStr := nts[0:pattern]
				for j := pattern; j < lenNts; j += pattern {
					if nts[j:j+pattern] != patternStr {
						repeated = false
						break
					}
				}
				if repeated {
					result += i
					break
				}
			}

			// 4174379265
			//13

		}

	}

	return result
}
