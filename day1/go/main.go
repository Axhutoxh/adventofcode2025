package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// The input list of commands from the previous interaction
var commands = []string{}

type Command struct {
	Direction string // "L" or "R"
	Distance  int    // The numeric distance
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
		commands = append(commands, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file line-by-line:", err)
	}

	startingPoint := 50

	parsedCommands := parseCommands(commands)

	partOneSol := partOne(startingPoint, parsedCommands)
	fmt.Printf("Number of times at position 0: %d\n", partOneSol)

	partTwoSol := partTwo(startingPoint, parsedCommands)
	fmt.Printf("Final position after all commands: %d\n", partTwoSol)

}

func partTwo(startingPoint int, parsedCommands []Command) (finalPosition int) {
	zeroCounter := 0
	for _, cmd := range parsedCommands {
		for i := 0; i < cmd.Distance; i++ {
			if cmd.Direction == "L" {
				startingPoint = (startingPoint - 1 + 100) % 100
			} else if cmd.Direction == "R" {
				startingPoint = (startingPoint + 1) % 100
			}
			if startingPoint == 0 {
				zeroCounter++
			}
		}
	}

	return zeroCounter
}

func partOne(startingPoint int, parsedCommands []Command) (finalZeroCounter int) {
	zeroCounter := 0
	for _, cmd := range parsedCommands {
		if cmd.Direction == "L" {
			startingPoint = (startingPoint - cmd.Distance + 100) % 100
		} else if cmd.Direction == "R" {
			startingPoint = (startingPoint + cmd.Distance) % 100
		}
		if startingPoint == 0 {
			zeroCounter++
		}
	}
	return zeroCounter
}

func parseCommands(rawCommands []string) []Command {
	var results []Command

	for _, rawCmd := range rawCommands {
		direction := string(rawCmd[0])
		distanceStr := rawCmd[1:]
		distance, err := strconv.Atoi(distanceStr)
		if err != nil {
			fmt.Printf("Error converting distance for command '%s': %v\n", rawCmd, err)
			continue
		}
		results = append(results, Command{
			Direction: direction,
			Distance:  distance,
		})
	}

	return results
}
