package day4

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func SolutionPart2() int {
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

	var valuesToRemove [][2]int

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

	return totalRolls
}
