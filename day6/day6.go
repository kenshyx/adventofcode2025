package day6

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kenshyx/adventofcode2025/utils"
)

var ops = map[string]func(int, int) int{
	"+": func(a, b int) int { return a + b },
	"*": func(a, b int) int { return a * b },
}
var matrix [][]int
var operators []string

func GetSolution(authenticatedR *utils.UrlWithAuth) utils.Solution {
	reader, resp := authenticatedR.FetchInput()
	if resp != nil {
		defer resp.Body.Close()
	}
	lineNumber := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		columns := strings.Fields(line)
		values := make([]int, 0)
		for _, col := range columns {
			vv, ee := strconv.Atoi(col)
			if ee != nil {
				_, ok := ops[col]
				if !ok {
					fmt.Printf("error parsing column %s\n", col)
					continue
				}
				operators = append(operators, col)
				continue
			}
			values = append(values, vv)
		}
		if len(values) != 0 {
			matrix = append(matrix, values)
		}

		lineNumber++
	}
	total := 0
	for j := 0; j < len(matrix[0]); j++ {
		subTotal := 0
		if operators[j] == "*" {
			subTotal = 1
		}
		for jk := 0; jk < len(matrix); jk++ {
			subTotal = ops[operators[j]](subTotal, matrix[jk][j])

		}
		total += subTotal
	}

	return utils.Solution{
		Part1: total,
		Part2: SolutionPart2(authenticatedR),
	}
}
