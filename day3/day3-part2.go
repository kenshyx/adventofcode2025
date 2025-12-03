package day3

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

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

func SolutionPart2() int {
	file, _ := os.Open("day3/input.txt")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	reader := bufio.NewReader(file)
	totalPower := 0
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			break
		}

		subString := extractMaxDigits(11, line, "")
		result, _ := strconv.Atoi(subString)
		totalPower += result
	}
	return totalPower
}
