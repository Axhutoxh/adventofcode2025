package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	grid     [][]rune
	rows     int
	cols     int
	startCol int
	visited  map[[2]int]bool
	count    int
	memo     map[[2]int]int
)

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
	partTwoSol := partTwo(commands)

	fmt.Println(partOneSol, partTwoSol)

}

func partOne(cmd []string) int {

	rows = len(cmd)
	grid = make([][]rune, rows)

	for r, line := range cmd {
		grid[r] = []rune(line)
		if r == 0 {
			cols = len(grid[r])

			for c, char := range grid[r] {
				if char == 'S' {
					startCol = c
					break
				}
			}
		}
	}

	visited = make(map[[2]int]bool)
	count = 0

	traverse(1, startCol)

	return count
}

func traverse(row, col int) {
	if row < 0 || row >= rows || col < 0 || col >= cols {
		return
	}

	pos := [2]int{row, col}

	if visited[pos] {
		return
	}

	visited[pos] = true
	if row < rows && col < cols {
		currentCell := grid[row][col]

		if currentCell == '^' {
			count++
			traverse(row+1, col-1)
			traverse(row+1, col+1)
		} else if currentCell == '.' {
			traverse(row+1, col)
		}
	}
}

func partTwo(cmd []string) int {

	rows = len(cmd)

	cols = len(cmd[0])
	grid = make([][]rune, rows)

	startCol := -1

	for r, line := range cmd {

		grid[r] = []rune(line)

		if r == 0 {
			for c, char := range grid[r] {
				if char == 'S' {
					startCol = c
					break
				}
			}
		}
	}

	if startCol == -1 {
		fmt.Println("Error: 'S' not found in the first row.")

	}

	memo = make(map[[2]int]int)

	totalPaths := traversePartSecond(1, startCol)

	return totalPaths
}

func traversePartSecond(row, col int) int {
	if row < 0 || row >= rows || col < 0 || col >= cols {
		return 1
	}

	key := [2]int{row, col}
	if result, found := memo[key]; found {
		return result
	}

	currentCell := grid[row][col]
	result := 0

	if currentCell == '^' {
		leftPath := traversePartSecond(row+1, col-1)
		rightPath := traversePartSecond(row+1, col+1)
		result = leftPath + rightPath

	} else if currentCell == '.' {

		result = traversePartSecond(row+1, col)

	} else {

		result = 0
	}

	memo[key] = result
	return result
}
