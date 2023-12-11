package main

import (
	"log"
	"strconv"
	"strings"
)

type DaynineInput struct {
	a []string
}

type HistoryLines struct {
	histories  []History
	historySum int
}

type History struct {
	history               []int
	extraPolatedHistories [][]int
	nextValue             int
}

func parseHistoryLines(lines []string) HistoryLines {
	var historyLines HistoryLines
	for _, line := range lines {
		history := parseHistoryLine(line)
		historyLines.histories = append(historyLines.histories, history)
	}
	return historyLines
}

func parseHistoryLine(line string) History {
	var history History
	for _, numbers := range strings.Split(strings.TrimSpace(line), " ") {
		res, err := strconv.Atoi(numbers)
		if err != nil {
			log.Fatalf("Could not parse history number %v\n", numbers)
		}
		history.history = append(history.history, res)
	}
	history.extraPolatedHistories = make([][]int, 1)
	history.extraPolatedHistories[0] = history.history
	return history
}

func (hl *HistoryLines) extraPolateHistory(partTwo bool) {
	sum := 0
	for _, history := range hl.histories {

		history.extraPolateHistory(partTwo)
		sum = sum + history.nextValue

	}

	hl.historySum = sum

}
func (h *History) extraPolateHistory(partTwo bool) {
	// To do this, start by making a new sequence from the difference at each step of your history.
	allZeroes := false
	for i := 0; i < len(h.extraPolatedHistories) && !allZeroes; i++ {
		newLine := make([]int, 0)
		allZeroes = true
		for j := 1; j < len(h.extraPolatedHistories[i]); j++ {
			difference := h.extraPolatedHistories[i][j] - h.extraPolatedHistories[i][j-1]
			newLine = append(newLine, difference)
			if difference != 0 {
				allZeroes = false
			}
		}
		h.extraPolatedHistories = append(h.extraPolatedHistories, newLine)
	}

	h.extraPolatedHistories[len(h.extraPolatedHistories)-1] = append(h.extraPolatedHistories[len(h.extraPolatedHistories)-1], 0)
	// Once all of the values in your latest sequence are zeroes, you can extrapolate what the next value of the original history should be.
	for i := len(h.extraPolatedHistories) - 2; i >= 0; i-- {
		result := h.extraPolatedHistories[i+1][len(h.extraPolatedHistories[i+1])-1] + h.extraPolatedHistories[i][len(h.extraPolatedHistories[i])-1]
		h.extraPolatedHistories[i] = append(h.extraPolatedHistories[i], result)
	}
	h.nextValue = h.extraPolatedHistories[0][len(h.extraPolatedHistories[0])-1]

	if partTwo {
		h.extraPolatedHistories[len(h.extraPolatedHistories)-1] = append([]int{0}, h.extraPolatedHistories[len(h.extraPolatedHistories)-1]...)
		// Once all of the values in your latest sequence are zeroes, you can extrapolate what the next value of the original history should be.
		for i := len(h.extraPolatedHistories) - 2; i >= 0; i-- {
			result := h.extraPolatedHistories[i][0] - h.extraPolatedHistories[i+1][0]
			h.extraPolatedHistories[i] = append([]int{result}, h.extraPolatedHistories[i]...)
		}
        h.nextValue = h.extraPolatedHistories[0][0]
	}

}
func Daynine(in DaynineInput) int {
	partTwo := false
	historyLines := parseHistoryLines(in.a)
	historyLines.extraPolateHistory(partTwo)
	return historyLines.historySum
}
func DayninePartTwo(in DaynineInput) int {
	partTwo := true
	historyLines := parseHistoryLines(in.a)
	historyLines.extraPolateHistory(partTwo)
	return historyLines.historySum
}
func main() {
}
