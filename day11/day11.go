package day11

import (
	"fmt"
	"strings"

	"github.com/kenshyx/adventofcode2025/utils"
)

const (
	START = "you"
	OUT   = "out"
	DAC   = "dac"
	FFT   = "fft"
	SVR   = "svr"
)

type Devices map[string][]string

type state struct {
	device string
	dFound bool
	fFound bool
}

func (devices Devices) calculatePaths(device string) int {
	if device == OUT {
		return 1
	}
	sum := 0
	for _, d := range devices[device] {
		sum += devices.calculatePaths(d)
	}
	return sum
}

func (devices Devices) calculateDacFftPaths(device string, dFound bool, fFound bool, memo map[state]int) int {
	key := state{device, dFound, fFound}
	if val, ok := memo[key]; ok {
		return val
	}

	if device == OUT {
		if dFound && fFound {
			memo[key] = 1
		} else {
			memo[key] = 0
		}
		return memo[key]
	}

	sum := 0
	for _, d := range devices[device] {
		dacFound := dFound || d == DAC
		fftFound := fFound || d == FFT
		sum += devices.calculateDacFftPaths(d, dacFound, fftFound, memo)
	}
	memo[key] = sum
	return sum
}

func GetSolution(authenticatedR *utils.UrlWithAuth) utils.Solution {
	reader, resp := authenticatedR.FetchInput()
	if resp != nil {
		defer resp.Body.Close()
	}
	devices := make(Devices)

	for {
		line, e := reader.ReadString('\n')
		if e != nil {
			break
		}
		line = strings.TrimSpace(line)
		segments := strings.Split(line, ":")
		if len(segments) != 2 {
			fmt.Println("Invalid input", line)
			continue
		}
		devices[segments[0]] = strings.Fields(segments[1])

	}

	totalP1 := devices.calculatePaths(START)
	memo := make(map[state]int)
	totalP2 := devices.calculateDacFftPaths(SVR, false, false, memo)

	return utils.Solution{
		Part1: totalP1,
		Part2: totalP2,
	}
}
