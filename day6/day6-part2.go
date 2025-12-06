package day6

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/kenshyx/adventofcode2025/utils"
)

var matrix2 [][]string

func SolutionPart2() int {

	reader, resp := utils.FetchInput("https://adventofcode.com/2025/day/6/input")
	if resp != nil {
		defer resp.Body.Close()
	}
	maxChars := 0
	for {
		line, err := reader.ReadString('\n')
		if strings.TrimSpace(line) == "" {
			break
		}
		ll := strings.Split(line, "")
		matrix2 = append(matrix2, ll)
		if len(ll) > maxChars {
			maxChars = len(ll)
		}
		if err != nil {
			break
		}

	}

	var vals []int
	total := 0
	subtotal := 0
	operator := ""
	var numRe = regexp.MustCompile(`\d+`)
	for q := maxChars - 1; q >= 0; q-- {
		nr := ""
		currentOperator := strings.TrimSpace(matrix2[len(matrix2)-1][q])

		if currentOperator != "" {
			operator = currentOperator
		}
		for i := 0; i < len(matrix2); i++ {
			if q >= len(matrix2[i]) {
				continue
			}
			nr += matrix2[i][q]
		}
		if strings.TrimSpace(nr) == "" {
			if operator == "*" {
				subtotal = 1
			}
			for _, ll := range vals {
				subtotal = ops[operator](subtotal, ll)
			}
			total += subtotal
			subtotal = 0
			vals = vals[:0]
			continue
		}
		match := numRe.FindString(nr)
		if match == "" {
			continue
		}
		val, _ := strconv.Atoi(match)
		vals = append(vals, val)
		// collect the rest
		if q == 0 {
			if operator == "*" {
				subtotal = 1
			}
			for _, ll := range vals {
				subtotal = ops[operator](subtotal, ll)
			}
			total += subtotal
		}
	}
	return total
}
