package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// The input list of inputs from the previous interaction
var inputs = []string{}

type coordinate struct {
	x int
	y int
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

	parseCommands := parseCommands(inputs)

	partOneSol := calMaxAreaOfRectange(parseCommands)
	fmt.Println(partOneSol)

	// partOneSol := partOne(parseCommands)

	// fmt.Println(" Part 1 solution ", partOneSol)

	// partTwoSol := partTwo(parseCommands)

	// fmt.Println(" Part 2 solution ", partTwoSol)

}

func parseCommands(rawCommands []string) []coordinate {
	var results []coordinate

	for _, rawCmd := range rawCommands {
		subStr := strings.Split(rawCmd, ",")

		xCor, err := strconv.Atoi(subStr[0])
		yCor, err := strconv.Atoi(subStr[1])

		if err != nil {
			fmt.Println("Unable to conver into int because :", err)
		}

		results = append(results, coordinate{
			x: xCor,
			y: yCor,
		})

	}

	return results
}

func calMaxAreaOfRectange(cor []coordinate) int {
	maxArea := 0

	for i := 0; i < len(cor); i++ {
		x1, y1 := cor[i].x, cor[i].y
		for j := 1; j < len(cor); j++ {
			x2, y2 := cor[j].x, cor[j].y

			width := math.Abs(float64(x2-x1)) + 1
			height := math.Abs(float64(y2-y1)) + 1

			area := width * height

			if area > float64(maxArea) {
				maxArea = int(area)
			}

		}
	}
	return maxArea

}
