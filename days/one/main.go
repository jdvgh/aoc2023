package main

import (
	"strconv"
	"strings"
)

type DayoneInput struct {
	a []string
}

func Dayone(in DayoneInput) int {
	sum := 0
	for _, line := range in.a {
		lower, _ := firstDigitFromLeft(line)

		higher, _ := firstDigitFromRight(line)

		sum = sum + lower*10 + higher
	}
	return sum
}

func firstDigitFromLeft(line string) (digit int, pos int) {

	for i := 0; i < len(line); i++ {
		res, err := strconv.Atoi(string(line[i]))
		if err == nil {
			return res, i
		}
	}
	return 0, len(line)

}

func firstDigitFromRight(line string) (digit int, pos int) {
	for j := len(line) - 1; j >= 0; j-- {
		res, err := strconv.Atoi(string(line[j]))
		if err == nil {
			return res, j
		}
	}
	return 0, 0
}

func DayonePartTwo(in DayoneInput) int {
	sum := 0
	for _, line := range in.a {
		lower, digitPosLeft := firstDigitFromLeft(line)
		digit := 0
		if digitPosLeft > 2 {

			digit = firstSpelledOutDigitFromLeft(line, digitPosLeft)
			if digit >= 0 {
				lower = digit
			}
		}

		higher, digitPosRight := firstDigitFromRight(line)
		if digitPosRight < len(line)-3 {

			digit = firstSpelledOutDigitFromRight(line, digitPosRight)
			if digit >= 0 {
				higher = digit
			}

		}

		sum = sum + lower*10 + higher
	}

	return sum
}

func findSpelledOutDigit(line string) int {

	currentLine := line
	if strings.Contains(currentLine, "one") {
		return 1
	} else if strings.Contains(currentLine, "two") {
		return 2
	} else if strings.Contains(currentLine, "three") {
		return 3
	} else if strings.Contains(currentLine, "four") {
		return 4
	} else if strings.Contains(currentLine, "five") {
		return 5
	} else if strings.Contains(currentLine, "six") {
		return 6
	} else if strings.Contains(currentLine, "seven") {
		return 7
	} else if strings.Contains(currentLine, "eight") {
		return 8
	} else if strings.Contains(currentLine, "nine") {
		return 9
	}
	return -1
}

func firstSpelledOutDigitFromLeft(line string, digitPosLeft int) int {

	for i := 2; i < len(line) && i < digitPosLeft; i++ {
		currentLine := line[0 : i+1]
		res := findSpelledOutDigit(currentLine)
		if res > 0 {
			return res
		}
	}
	return -1
}
func firstSpelledOutDigitFromRight(line string, digitPosRight int) int {

	for j := len(line) - 2; j >= 0 && j > digitPosRight; j-- {
		currentLine := line[j:]
		res := findSpelledOutDigit(currentLine)
		if res > 0 {
			return res
		}
	}
	return -1
}
func main() {
}
