package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var commands = []string{}
var commandS string

type operation func(a, b int) int

var (
	add operation = func(a, b int) int {
		return a + b
	}

	mul operation = func(a, b int) int {
		return a * b
	}
)

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
		commandS = commandS + line + "\n"
		commands = append(commands, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file line-by-line:", err)
	}

	parsedIntCmdArrO, parseOprCmdArrO := parseCommandsForPartOne(commands)

	partOneSol := partOne(parsedIntCmdArrO, parseOprCmdArrO)

	partTwoSol := solve(commands)
	fmt.Println(partTwoSol, partOneSol)

}

func parseCommandsForPartOne(rawCommands []string) ([][]int, []string) {
	var inputsArr [][]int
	var operatorArr []string

	for _, rawCmd := range rawCommands {
		rawInputs := strings.Split(rawCmd, " ")
		var row []int
		for j := 0; j < len(rawInputs); j++ {

			if len(rawInputs[j]) > 0 {
				rawInt, err := strconv.Atoi(rawInputs[j])

				if err != nil {
					operatorArr = append(operatorArr, rawInputs[j])
				} else {
					row = append(row, rawInt)
				}

			}
		}
		if len(row) > 0 {
			inputsArr = append(inputsArr, row)
		}
	}

	return inputsArr, operatorArr
}

func partOne(intCmdArr [][]int, oprCmdArr []string) int {
	result := 0

	for i := 0; i < len(intCmdArr[0]); i++ {

		out := 0
		if oprCmdArr[i] == "*" {
			out = intCmdArr[0][i] * intCmdArr[1][i] * intCmdArr[2][i] * intCmdArr[3][i]

		} else {
			out = intCmdArr[0][i] + intCmdArr[1][i] + intCmdArr[2][i] + intCmdArr[3][i]
		}
		result += out
	}
	return result
}

func reduce(op operation, numbers []int) int {
	if len(numbers) == 0 {
		return 0 // Or panic, depending on expected behavior for empty slice
	}
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = op(result, numbers[i])
	}
	return result
}

func solve(lines []string) int {
	if len(lines) == 0 {
		return 0
	}
	width := 0
	for _, line := range lines {
		trimmedLine := strings.TrimRightFunc(line, unicode.IsSpace)
		if len(trimmedLine) > width {
			width = len(trimmedLine)
		}
	}

	if width == 0 {
		return 0
	}

	grid := make([]string, len(lines))
	h := len(grid)
	for i, line := range lines {
		trimmedLine := strings.TrimRightFunc(line, unicode.IsSpace)
		grid[i] = trimmedLine + strings.Repeat(" ", width-len(trimmedLine))
	}

	sep := make([]bool, width)
	for c := 0; c < width; c++ {
		isSeparator := true
		for r := 0; r < h; r++ {
			if grid[r][c] != ' ' {
				isSeparator = false
				break
			}
		}
		sep[c] = isSeparator
	}

	type problemRange struct {
		start int
		end   int
	}
	problemRanges := []problemRange{}
	inProblem := false
	start := 0
	for c := 0; c < width; c++ {
		if !sep[c] {
			if !inProblem {
				inProblem = true
				start = c
			}
		} else {
			if inProblem {
				inProblem = false
				problemRanges = append(problemRanges, problemRange{start, c - 1})
			}
		}
	}
	if inProblem {
		problemRanges = append(problemRanges, problemRange{start, width - 1})
	}

	grandTotal := 0
	for _, pr := range problemRanges {
		cStart := pr.start
		cEnd := pr.end

		opRow := -1
		for r := 0; r < h; r++ {
			segment := grid[r][cStart : cEnd+1]
			if strings.Contains(segment, "+") || strings.Contains(segment, "*") {
				opRow = r
				break
			}
		}

		if opRow == -1 {
			continue
		}

		opSegment := grid[opRow][cStart : cEnd+1]

		var op operation

		plusIdx := strings.Index(opSegment, "+")
		mulIdx := strings.Index(opSegment, "*")

		if plusIdx != -1 && mulIdx != -1 {
			if plusIdx < mulIdx {
				op = add
			} else {
				op = mul
			}
		} else if plusIdx != -1 {
			op = add
		} else if mulIdx != -1 {
			op = mul
		} else {
			continue
		}

		numbers := []int{}
		for c := cStart; c <= cEnd; c++ {
			digits := []rune{}
			for r := 0; r < opRow; r++ {
				ch := rune(grid[r][c])
				if unicode.IsDigit(ch) {
					digits = append(digits, ch)
				}
			}

			if len(digits) > 0 {
				num, err := strconv.Atoi(string(digits))
				if err == nil {
					numbers = append(numbers, num)
				}
			}
		}

		if len(numbers) == 0 {
			continue
		}

		result := reduce(op, numbers)
		grandTotal += result
	}

	return grandTotal
}
