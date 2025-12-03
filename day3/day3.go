package day3

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func SolutionPart1() int {
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
		firstDigit := 0
		firstDigitPosition := 0
		lastDigit := 0
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
	return totalPower
}
