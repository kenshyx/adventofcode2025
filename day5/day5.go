package day5

import (
	"sort"
	"strconv"
	"strings"

	"github.com/kenshyx/adventofcode2025/utils"
)

type Interval = [][2]int

var intervals Interval

func GetSolution(authenticatedR *utils.UrlWithAuth) utils.Solution {
	reader, resp := authenticatedR.FetchInput()
	if resp != nil {
		defer resp.Body.Close()
	}
	startCheck := false
	counter := 0
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			break
		}
		if strings.Contains(line, "-") {
			interval := strings.SplitN(line, "-", 2)
			start, _ := strconv.Atoi(interval[0])
			stop, _ := strconv.Atoi(interval[1])
			intervals = append(intervals, [2]int{start, stop})
		}
		if line == "" {
			startCheck = true
			continue
		}
		if startCheck {
			ingredient, _ := strconv.Atoi(line)
			for _, v := range intervals {
				if ingredient >= v[0] && ingredient <= v[1] {
					counter++
					break
				}
			}
		}
	}
	totalFreshIngredients := 0

	var merged Interval
	current := intervals[0]

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 1; i < len(intervals); i++ {
		next := intervals[i]

		if next[0] <= current[1] {
			if next[1] > current[1] {
				current[1] = next[1]
			}
		} else {
			merged = append(merged, current)
			current = next
		}
	}
	merged = append(merged, current)

	for _, v := range merged {
		totalFreshIngredients += v[1] - v[0] + 1
	}

	return utils.Solution{
		Part1: counter,
		Part2: totalFreshIngredients,
	}
}
