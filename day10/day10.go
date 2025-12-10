package day10

import (
	"strconv"
	"strings"

	"github.com/kenshyx/adventofcode2025/utils"
)

const (
	SwitchOn = '#'
)

func GetSolution(authenticatedR *utils.UrlWithAuth) utils.Solution {
	reader, resp := authenticatedR.FetchInput()
	if resp != nil {
		defer resp.Body.Close()
	}
	p1 := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		words := strings.Fields(line)
		if len(words) == 0 {
			continue
		}

		goal := words[0]
		if len(goal) < 2 {
			continue
		}
		goal = goal[1 : len(goal)-1]

		goalN := 0
		for i, c := range goal {
			if c == SwitchOn {
				goalN += 1 << i
			}
		}

		buttonWords := words[1 : len(words)-1]
		B := make([]int, 0, len(buttonWords))

		for _, button := range buttonWords {
			if len(button) < 2 {
				B = append(B, 0)
				continue
			}
			inner := button[1 : len(button)-1]
			if strings.TrimSpace(inner) == "" {
				B = append(B, 0)
				continue
			}

			parts := strings.Split(inner, ",")
			mask := 0
			for _, part := range parts {
				part = strings.TrimSpace(part)
				if part == "" {
					continue
				}
				x, er := strconv.Atoi(part)
				if er != nil {
					continue
				}
				mask |= 1 << x
			}
			B = append(B, mask)
		}

		n := len(B)
		score := n

		limit := 1 << n
		for subset := 0; subset < limit; subset++ {
			an := 0
			aScore := 0

			for i := 0; i < n; i++ {
				if (subset>>i)&1 == 1 {
					an ^= B[i]
					aScore++
				}
			}

			if an == goalN && aScore < score {
				score = aScore
			}
		}

		p1 += score
	}

	return utils.Solution{
		Part1: p1,
	}
}
