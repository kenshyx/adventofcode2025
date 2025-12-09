package day2

import (
	"log"
	"strconv"
	"strings"

	"github.com/kenshyx/adventofcode2025/utils"
)

func IsMadeOfRepeats(s, sub string) bool {
	if len(sub) == 0 {
		return false
	}
	if len(s)%len(sub) != 0 {
		return false
	}

	for i := 0; i < len(s); i += len(sub) {
		if s[i:i+len(sub)] != sub {
			return false
		}
	}
	return true
}

func GetSolution(authenticatedR *utils.UrlWithAuth) utils.Solution {
	reader, resp := authenticatedR.FetchInput()
	if resp != nil {
		defer resp.Body.Close()
	}
	sum := 0
	sum1 := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		intervals := strings.Split(line, ",")
		for _, interval := range intervals {
			parts := strings.Split(interval, "-")
			if len(parts) != 2 {
				log.Fatal("Invalid interval format", interval)
			}
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			for i := start; i <= end; i++ {
				curr := strconv.Itoa(i)
				mid := len(curr) / 2
				left := curr[:mid]
				right := curr[mid:]
				if left == right {
					sum += i
					sum1 += i
					continue
				}
				for l := 1; l <= mid; l++ {
					if len(curr)%l != 0 {
						continue
					}
					unit := curr[:l]
					if IsMadeOfRepeats(curr, unit) {
						sum += i
						break
					}
				}
			}
		}
	}
	return utils.Solution{
		Part1: sum1,
		Part2: sum,
	}
}
