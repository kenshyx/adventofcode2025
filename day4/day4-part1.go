package day4

import (
	"bufio"
	"log"
	"os"
	"strings"
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

func SolutionPart1() int {
	file, _ := os.Open("day4/input.txt")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	reader := bufio.NewReader(file)
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
	for i, j := range grid {
		for k, v := range j {
			neighbors := extractMatchingNeighbors(grid, i, k)
			if len(neighbors) < 4 && v == rollOfPaper {
				totalRolls++
			}
		}
	}

	return totalRolls
}
