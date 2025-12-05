package day5

import "sort"

func SolutionPart2() int {

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
	return totalFreshIngredients
}
