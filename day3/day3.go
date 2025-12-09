package day3

import (
	"log"
	"strconv"
	"strings"

	"github.com/kenshyx/adventofcode2025/utils"
)

func GetSolution(authenticatedR *utils.UrlWithAuth) utils.Solution {
	reader, resp := authenticatedR.FetchInput()
	if resp != nil {
		defer resp.Body.Close()
	}
	totalPower := 0
	totalPower2 := 0
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			break
		}
		firstDigit := 0
		firstDigitPosition := 0
		lastDigit := 0
		subString := extractMaxDigits(11, line, "")
		result2, _ := strconv.Atoi(subString)
		totalPower2 += result2
		for k, char := range line[:len(line)-1] {
			d := int(char - '0')
			if d > firstDigit {
				firstDigit = d
				firstDigitPosition = k
			}
		}
		if firstDigitPosition+1 > len(line) {
			log.Println("out of range")
			continue
		}
		for _, char := range line[firstDigitPosition+1:] {
			d := int(char - '0')
			if d > lastDigit {
				lastDigit = d
			}

		}
		strNumber := strconv.Itoa(firstDigit) + strconv.Itoa(lastDigit)
		result, _ := strconv.Atoi(strNumber)
		totalPower += result
	}
	return utils.Solution{
		Part1: totalPower,
		Part2: totalPower2,
	}
}

func extractMaxDigits(maxLength int, slice string, substring string) string {

	digit := 0
	digitPos := 0
	substr := substring
	for k, char := range slice[:len(slice)-maxLength] {
		d := int(char - '0')
		if d > digit {
			digit = d
			digitPos = k
		}
	}
	substr += strconv.Itoa(digit)
	if maxLength == 0 {
		return substr
	} else {
		maxLength--
		return extractMaxDigits(maxLength, slice[digitPos+1:], substr)
	}
}
