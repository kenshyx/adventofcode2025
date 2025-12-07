package day7

import (
	"strings"

	"github.com/kenshyx/adventofcode2025/utils"
)

const StartChar = "S"
const SplitterChar = "^"
const SpaceChar = "."
const BeamChar = "|"

var matrix [][]string
var totalSplits = 0

func parseRow(row int, pos []int) int {
	if row == len(matrix) {
		return totalSplits
	}
	var newPos []int
	for _, v := range pos {
		if matrix[row][v] == SplitterChar {
			leftPos := v - 1
			rightPos := v + 1
			if leftPos >= 0 {
				if matrix[row][leftPos] == SpaceChar {
					matrix[row][leftPos] = BeamChar
					newPos = append(newPos, leftPos)
				}
			}
			if rightPos < len(matrix[row]) {
				if matrix[row][rightPos] == SpaceChar {
					matrix[row][rightPos] = BeamChar
					newPos = append(newPos, rightPos)
				}
			}

			if leftPos >= 0 || rightPos < len(matrix[row]) {
				totalSplits++
			}
		} else if matrix[row][v] == SpaceChar {
			matrix[row][v] = BeamChar
			newPos = append(newPos, v)
		}
	}

	return parseRow(row+1, newPos)
}

func SolutionPart1() int {
	reader, resp := utils.FetchInput("https://adventofcode.com/2025/day/7/input")
	if resp != nil {
		defer resp.Body.Close()
	}
	startPos := -1
	for {
		line, e := reader.ReadString('\n')
		if e != nil {
			break
		}
		line = strings.TrimSpace(line)
		lineChars := strings.Split(line, "")
		if startPos == -1 {
			for i, v := range lineChars {
				if v == StartChar {
					startPos = i
				}
			}
		}
		matrix = append(matrix, lineChars)
	}

	for k, ll := range matrix {
		if ll[startPos] == StartChar {
			parseRow(k+1, []int{startPos})
			break
		}
	}
	return totalSplits
}
