package main

import (
	"fmt"

	"github.com/kenshyx/adventofcode2025/day1"
	"github.com/kenshyx/adventofcode2025/day2"
	"github.com/kenshyx/adventofcode2025/day3"
	"github.com/kenshyx/adventofcode2025/day4"
	"github.com/kenshyx/adventofcode2025/day5"
)

func main() {
	fmt.Println("Solution day one:", day1.Solution())
	fmt.Println("Solution day two:", day2.Solution())
	fmt.Println("Solutions for day three:", day3.SolutionPart1(), day3.SolutionPart2())
	fmt.Println("Solutions for day four:", day4.SolutionPart1(), day4.SolutionPart2())
	fmt.Println("Solutions for day five:", day5.SolutionPart1(), day5.SolutionPart2())
}
