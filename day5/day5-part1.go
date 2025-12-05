package day5

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Interval = [][2]int

var intervals Interval

func SolutionPart1() int {
	file, e := os.Open("day5/input.txt")
	if e != nil {
		log.Fatal(e)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	reader := bufio.NewReader(file)
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
	return counter
}
