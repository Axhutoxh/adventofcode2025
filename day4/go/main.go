package main

import (
	"bufio"
	"fmt"
	"os"
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

	partOne := partOneSol(makeGrid(commands))
	fmt.Printf(" Part 1 solution %d\n", partOne)

	partTwoSol := partTwoSol(makeGrid(commands))
	fmt.Printf(" Part 2 solution %d\n", partTwoSol)

}

func makeGrid(commands []string) [][]string {
	rows := len(commands)
	grid := make([][]string, rows)

	for i := 0; i < rows; i++ {
		cols := len(commands[i])
		grid[i] = make([]string, cols)
		grid[i] = strings.Split(commands[i], "")

	}

	return grid
}
func partOneSol(grid [][]string) int {
	rowCount := len(grid)
	if rowCount == 0 {
		return 0
	}
	colCount := len(grid[0]) // Assumes a rectangular grid

	rollCount := 0
	// Use the same structure for the new grid
	xGrid := make([][]string, rowCount)

	for i := 0; i < rowCount; i++ {
		xGrid[i] = make([]string, colCount)

		for j := 0; j < colCount; j++ {

			if grid[i][j] == "@" {
				aAdjCounter := 0

				for di := -1; di <= 1; di++ {
					for dj := -1; dj <= 1; dj++ {

						if di == 0 && dj == 0 {
							continue
						}

						ni, nj := i+di, j+dj // Neighbor coordinates

						// Check boundaries
						if ni >= 0 && ni < rowCount && nj >= 0 && nj < colCount {
							// Count adjacent rolls in the original grid
							if grid[ni][nj] == "@" {
								aAdjCounter++
							}
						}
					}
				}

				if aAdjCounter < 4 {
					xGrid[i][j] = "x" // Accessible by forklift
					rollCount++
				} else {
					xGrid[i][j] = "@" // Not accessible
				}
			} else {
				// 3. Keep non-roll spots as '.'
				xGrid[i][j] = "."
			}
		}
	}

	return rollCount
}

func isRoll(grid [][]string, r, c int) bool {
	rowCount := len(grid)
	colCount := len(grid[0]) // Assumes rectangular grid

	if r >= 0 && r < rowCount && c >= 0 && c < colCount {
		return grid[r][c] == "@"
	}
	return false
}

func singleRemovalStep(grid [][]string) (int, [][]string) {
	rowCount := len(grid)
	if rowCount == 0 {
		return 0, grid
	}
	colCount := len(grid[0])

	nextState := make([][]string, rowCount)
	for i := range grid {
		nextState[i] = make([]string, colCount)
		copy(nextState[i], grid[i])
	}

	removedCount := 0

	for r := 0; r < rowCount; r++ {
		for c := 0; c < colCount; c++ {

			if grid[r][c] == "@" {
				adjacentRolls := 0

				for dr := -1; dr <= 1; dr++ {
					for dc := -1; dc <= 1; dc++ {

						if dr == 0 && dc == 0 {
							continue
						}

						if isRoll(grid, r+dr, c+dc) {
							adjacentRolls++
						}
					}
				}

				if adjacentRolls < 4 {

					nextState[r][c] = "."
					removedCount++
				} else {

					nextState[r][c] = "@"
				}
			}

		}
	}

	return removedCount, nextState
}

// The main solver function for Part Two
func partTwoSol(initialGrid [][]string) int {
	currentGrid := initialGrid
	totalRemoved := 0
	step := 1

	for {
		removedInStep, nextGrid := singleRemovalStep(currentGrid)

		if removedInStep == 0 {
			break
		}

		totalRemoved += removedInStep
		currentGrid = nextGrid

		step++
	}

	return totalRemoved
}

// func partOneSol(grid [][]string) int {
// 	rowCount := len(grid)
// 	fmt.Println()
// 	rollCount := 0
// 	xGrid := make([][]string, rowCount)
// 	for i := 0; i < rowCount; i++ {
// 		colCount := len(grid[i])
// 		xGrid[i] = make([]string, colCount)

// 		for j := 0; j < colCount; j++ {
// 			aAdjCounter := 0
// 			if grid[i][j] == "@" {

// 				if i > 0 && j > 0 && grid[i-1][j-1] == "@" {
// 					aAdjCounter++

// 				}
// 				if i > 0 && grid[i-1][j] == "@" {
// 					aAdjCounter++
// 				}
// 				if i > 0 && j < colCount-1 && grid[i-1][j+1] == "@" {
// 					aAdjCounter++
// 				}
// 				if j > 0 && grid[i][j-1] == "@" {
// 					aAdjCounter++
// 				}
// 				if j > 0 && i < rowCount-1 && grid[i+1][j-1] == "@" {
// 					aAdjCounter++
// 				}
// 				if i < rowCount-1 && grid[i+1][j] == "@" {
// 					aAdjCounter++
// 				}
// 				if i < rowCount-1 && j < colCount-1 && grid[i+1][j+1] == "@" {
// 					aAdjCounter++
// 				}
// 				if j < colCount-1 && grid[i][j+1] == "@" {
// 					aAdjCounter++
// 				}

// 				if aAdjCounter > 0 && aAdjCounter < 4 {
// 					xGrid[i][j] = "x"
// 					rollCount++
// 				} else {
// 					xGrid[i][j] = "@"
// 				}
// 			} else {
// 				xGrid[i][j] = "."
// 			}

// 		}
// 		fmt.Println(xGrid[i])
// 		// ..xx.xx@x.
// 		// x@@.@.@.@@
// 		// @@@@@.x.@@
// 		// @.@@@@..@.
// 		// x@.@@@@.@x
// 		// .@@@@@@@.@
// 		// .@.@.@.@@@
// 		// x.@@@.@@@@
// 		// .@@@@@@@@.
// 		// x.x.@@@.x.
// 	}

// 	return rollCount
// }
