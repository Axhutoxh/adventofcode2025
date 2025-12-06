package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// The input list of commands from the previous interaction
var commands = []string{}

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
		commands = append(commands, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file line-by-line:", err)
	}

	partOneSol := partOne(commands)

	fmt.Println(" Part 1 solution ", partOneSol)

	partTwoSol := partTwo(commands)
	fmt.Printf(" Part 2 solution %d\n", partTwoSol)

}

func partOne(parsedCommands []string) (totalJoltage int) {
	calculateMaxJoltage := 0
	for _, cmd := range parsedCommands {
		maxNumber := 0
		secondMaxNumber := 0
		maxTwoDigit := 0

		for i, str := range cmd {
			input, err := strconv.Atoi(string(str))
			if err != nil {
				fmt.Printf("Error converting input for command '%s': %v\n", str, err)
				continue
			}
			if input > maxNumber && i != len(cmd)-1 {

				maxNumber = input
				secondMaxNumber = 0
			} else if input > secondMaxNumber {
				secondMaxNumber = input
			}

			// 	fmt.Printf(" Index: %d Input: %d maxnumber: %d secondMaxNumber: %d \n", i, input, maxNumber, secondMaxNumber)
			//
		}
		maxTwoDigit = maxNumber*10 + secondMaxNumber
		calculateMaxJoltage += maxTwoDigit

	}
	return calculateMaxJoltage
}

func partTwo(parsedCommands []string) (totalJoltage int) {
	sumMaxNumber := 0
	for _, cmd := range parsedCommands {
		numberList := []int{}
		strList := strings.Split(cmd, "")

		for _, str := range strList {
			input, err := strconv.Atoi(str)
			if err != nil {
				fmt.Printf("Error converting input for command '%s': %v\n", str, err)
				continue
			}
			numberList = append(numberList, input)
		}

		idx := 0
		number := ""
		for noLeft := 11; noLeft >= 0; noLeft-- {
			index, res := findMax(numberList, noLeft, idx)
			idx = index
			number += strconv.Itoa(res)

		}
		parsedNumber, err := strconv.Atoi(number)
		if err != nil {
			fmt.Printf("Error converting number '%s': %v\n", number, err)
			continue
		}
		sumMaxNumber += parsedNumber

	}
	return sumMaxNumber

}

func findMax(arr []int, numberLeft int, idx int) (int, int) {
	newIdx := 0
	largest := 0
	for i := idx; i < len(arr)-numberLeft; i++ {
		if arr[i] > largest {
			largest = arr[i]
			newIdx = i + 1
		}
	}
	return newIdx, largest
}
