package day4

import (
	"strings"

	"github.com/kenshyx/adventofcode2025/utils"
)

var directions = [8][2]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
	{1, 1},
	{1, -1},
	{-1, 1},
	{-1, -1},
}

const rollOfPaper = "@"

func extractMatchingNeighbors(matrix [][]string, row int, column int) []string {
	maxRows := len(matrix)
	cols := len(matrix[0])
	var foundNeighbors []string
	if matrix[row][column] != rollOfPaper {
		return foundNeighbors
	}
	for _, d := range directions {
		nr := row + d[0]
		nc := column + d[1]

		if nr >= 0 && nr < maxRows && nc >= 0 && nc < cols {
			if matrix[nr][nc] != rollOfPaper {
				continue
			}
			foundNeighbors = append(foundNeighbors, matrix[nr][nc])
		}
	}

	return foundNeighbors
}

func GetSolution(authenticatedR *utils.UrlWithAuth) utils.Solution {
	reader, resp := authenticatedR.FetchInput()
	if resp != nil {
		defer resp.Body.Close()
	}
	var grid [][]string
	gridRow := 0
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			break
		}
		grid = append(grid, []string{})
		rowValues := strings.Split(line, "")
		grid[gridRow] = append(grid[gridRow], rowValues...)

		gridRow++
	}

	totalRolls := 0
	var valuesToRemove [][2]int
	var solution utils.Solution
	for {
		found := false
		for i, j := range grid {
			for k, v := range j {
				neighbors := extractMatchingNeighbors(grid, i, k)
				if len(neighbors) < 4 && v == rollOfPaper {
					valuesToRemove = append(valuesToRemove, [2]int{i, k})
					totalRolls++
				}
			}
		}

		if solution.Part1 == 0 {
			solution.Part1 = totalRolls
		}

		if len(valuesToRemove) > 0 {
			found = true
		}

		if !found {
			break
		}

		for _, vv := range valuesToRemove {
			grid[vv[0]][vv[1]] = "x"
		}
		// reset
		valuesToRemove = valuesToRemove[:0]
	}

	solution.Part2 = totalRolls

	return solution
}
