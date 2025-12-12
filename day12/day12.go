package day12

import (
	"io"
	"strconv"
	"strings"

	"github.com/kenshyx/adventofcode2025/utils"
)

func GetSolution(authenticatedR *utils.UrlWithAuth) utils.Solution {
	reader, resp := authenticatedR.FetchInput()
	if resp != nil {
		defer resp.Body.Close()
	}
	SIZES := make(map[int]int)

	var block []string
	inRegions := false
	ans := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}

		line = strings.TrimRight(line, "\n")

		if line == "" {
			if !inRegions && len(block) > 0 {
				parsePresent(block, SIZES)
				block = nil
			}
		} else {
			if !inRegions && strings.Contains(line, "x") && strings.Contains(line, ": ") {
				inRegions = true
			}

			if inRegions {
				processRegion(line, SIZES, &ans)
			} else {
				block = append(block, line)
			}
		}

		if err == io.EOF {
			break
		}
	}

	return utils.Solution{
		Part1: ans,
	}
}

func parsePresent(lines []string, SIZES map[int]int) {
	nameStr := strings.TrimSuffix(lines[0], ":")
	name, err := strconv.Atoi(nameStr)
	if err != nil {
		panic(err)
	}

	size := 0
	for _, row := range lines[1:] {
		for _, c := range row {
			if c == '#' {
				size++
			}
		}
	}

	SIZES[name] = size
}

func processRegion(line string, SIZES map[int]int, ans *int) {
	left, right, ok := strings.Cut(line, ": ")
	if !ok {
		panic("invalid region line: " + line)
	}

	rc := strings.Split(left, "x")
	R, err := strconv.Atoi(rc[0])
	if err != nil {
		panic(err)
	}
	C, er := strconv.Atoi(rc[1])
	if er != nil {
		panic(er)
	}

	fields := strings.Fields(right)

	totalPresentSize := 0
	for i, f := range fields {
		n, e := strconv.Atoi(f)
		if e != nil {
			panic(e)
		}
		totalPresentSize += n * SIZES[i]
	}

	totalGridSize := R * C

	if float64(totalPresentSize)*1.2 < float64(totalGridSize) {
		*ans++
	}
}
