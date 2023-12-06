package main

import (
	"log"
	"math"
	"strconv"
	"strings"
)

type DaysixInput struct {
	a []string
}

type SolvedRace struct {
	duration          int
	winningRaceLength int
	fLowerSolution    float64
	fHigherSolution   float64
	lowerSolution     int
	higherSolution    int
	winningTimes      int
}

func (s *SolvedRace) solvePQ() {
	halfDuration := float64((float64(s.duration) / float64(2)))
	inner := float64(halfDuration * halfDuration) - float64(s.winningRaceLength)
	sqrtContent := math.Sqrt(inner)
	s.fLowerSolution = halfDuration - sqrtContent
	s.fHigherSolution = halfDuration + sqrtContent
	s.lowerSolution = int(math.Ceil(s.fLowerSolution + 1.0e-14))
	s.higherSolution = int(math.Floor(s.fHigherSolution - 1.0e-14))
	s.winningTimes = s.higherSolution - s.lowerSolution + 1
}  

func parseInput(lines []string, two bool) int {

	times := strings.TrimSpace(strings.Split(lines[0], ":")[1])
	var timeNumbers []int
	if two {
		times = strings.ReplaceAll(times, " ", "")
	}
	for _, numberString := range strings.Split(times, " ") {
		if number, err := strconv.Atoi(string(numberString)); err == nil {
			timeNumbers = append(timeNumbers, number)
		}
	}
	lengths := strings.TrimSpace(strings.Split(lines[1], ":")[1])
	var winningLengths []int
	if two {
		lengths = strings.ReplaceAll(lengths, " ", "")
	}
	for _, lengthString := range strings.Split(lengths, " ") {
		if number, err := strconv.Atoi(string(lengthString)); err == nil {
			winningLengths = append(winningLengths, number)
		}
	}

	product := 1
	for i := 0; i < len(timeNumbers); i++ {
		solved := &SolvedRace{duration: timeNumbers[i], winningRaceLength: winningLengths[i]}
		solved.solvePQ()
        product = product * solved.winningTimes

	}
    return product

}
func Daysix(in DaysixInput) int {
	return parseInput(in.a, false)
}
func DaysixPartTwo(in DaysixInput) int {
	return parseInput(in.a, true)
}
func main() {
}
