package main

import "strconv"

type DaythreeInput struct {
	a []string
}

type Numbers struct {
	number int
	startI int
	startJ int
	endI   int
	endJ   int
}

func createNumbers(a []string) []Numbers {
	jLen := len(a[0])
	var numbers []Numbers
	currentNumber := Numbers{startI: -1, startJ: -1}
	for i, line := range a {
		if currentNumber.startI != -1 {
			currentNumber.endI = i - 1
			currentNumber.endJ = jLen - 1
			numbers = append(numbers, currentNumber)
		}
		currentNumber = Numbers{startI: -1, startJ: -1}
		for j, char := range line {
			digit, err := strconv.Atoi(string(char))
			if err != nil {
				if currentNumber.startI != -1 {
					currentNumber.endI = i
					currentNumber.endJ = j - 1
					numbers = append(numbers, currentNumber)
				}
				currentNumber = Numbers{startI: -1, startJ: -1}
			} else {
				if currentNumber.startI == -1 {
					currentNumber.startI = i
					currentNumber.startJ = j
				}
				currentNumber.number = currentNumber.number*10 + digit
			}
		}
	}
	return numbers
}

func DaythreePartTwo(in DaythreeInput) int {

	iLen := len(in.a)
	jLen := len(in.a[0])
	numbers := createNumbers(in.a)
	sum := 0

	for i, line := range in.a {
		for j, char := range line {
			if string(char) == "*" {
				num1 := 0
				num2 := 0
				okay := true
				minI2 := i - 1
				if minI2 < 0 {
					minI2 = i
				}
				maxI2 := i + 1
				if maxI2 == iLen {
					maxI2 = i
				}
				minJ2 := j - 1
				if minJ2 < 0 {
					minJ2 = j
				}
				maxJ2 := j + 1
				if maxJ2 == jLen {
					maxJ2 = j
				}
				for _, number := range numbers {
					if okay {
						if number.startI >= minI2 && number.startI <= maxI2 {
							if number.endJ >= minJ2 && number.startJ <= minJ2 || number.startJ > minJ2 && number.startJ <= maxJ2 {

								if num1 == 0 {
									num1 = number.number
								} else if num2 == 0 {
									num2 = number.number
								} else {
									okay = false
								}
							}

						}

					}

				}
				if okay {
					sum = sum + num1*num2

				}

			}
		}
	}
	return sum
}

func Daythree(in DaythreeInput) int {
	iLen := len(in.a)
	jLen := len(in.a[0])
	numbers := createNumbers(in.a)
	var okayRanges [][]bool
	okayRanges = make([][]bool, iLen)
	for i := 0; i < iLen; i++ {
		okayRanges[i] = make([]bool, jLen)
	}
	for i, line := range in.a {
		for j, char := range line {
			_, err := strconv.Atoi(string(char))
			if err != nil {
				switch string(char) {
				case ".":
				default:
					minI2 := i - 1
					if minI2 < 0 {
						minI2 = i
					}
					maxI2 := i + 1
					if maxI2 == iLen {
						maxI2 = i
					}
					minJ2 := j - 1
					if minJ2 < 0 {
						minJ2 = j
					}
					maxJ2 := j + 1
					if maxJ2 == jLen {
						maxJ2 = j
					}
					for i2 := minI2; i2 <= maxI2; i2++ {
						for j2 := minJ2; j2 <= maxJ2; j2++ {
							okayRanges[i2][j2] = true
						}
					}
				}
			}
		}
	}
	sum := 0
	for _, number := range numbers {
		okay := false
		for j := number.startJ; j <= number.endJ; j++ {
			okay = okay || okayRanges[number.startI][j]
		}

		if okay {
			sum += number.number
		}
	}
	return sum
}
func main() {
}
