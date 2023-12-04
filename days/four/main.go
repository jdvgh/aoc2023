package main

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

type DayfourInput struct {
	a []string
}
type Card struct {
	id               int
	winningNumbers   []int
	numbersScratched []int
	hits             int
}

var lut map[int]Card
var lut2 = []int{0, 1, 2, 4, 8, 16, 32, 64, 128, 256, 512}

func parseNumbers(numberString string) []int {
	numberStrings := strings.Split(strings.TrimSpace(numberString), " ")
	var numbers []int
	for _, numberA := range numberStrings {
		if numberA != "" {
			res, err := strconv.Atoi(numberA)
			if err != nil {
				log.Panicf("Could not parse %v as number\n", numberA)
			} else {
				numbers = append(numbers, res)
			}
		}

	}
	slices.Sort(numbers)
	return numbers

}

func parseCardId(line string) int {
	cardStringSplit := strings.Split(strings.TrimSpace(line), " ")
	numberString := cardStringSplit[len(cardStringSplit)-1]
	res, err := strconv.Atoi(numberString)
	if err != nil {
		log.Panicf("Could not parse %v as number\n", numberString)
	}
	return res

}
func parseCard(line string) Card {

	cardSplit := strings.Split(line, "|")
	winningNumbersSplit := strings.Split(cardSplit[0], ":")
	cardId := parseCardId(winningNumbersSplit[0])
	winningNumbers := parseNumbers(winningNumbersSplit[1])
	numbersScratched := parseNumbers(cardSplit[1])

	return Card{
		id:               cardId,
		winningNumbers:   winningNumbers,
		numbersScratched: numbersScratched}

}

func countHits(card Card) int {
	hits := 0
	for i := 0; i < len(card.numbersScratched); i++ {
		_, found := slices.BinarySearch(card.winningNumbers, card.numbersScratched[i])
		if found {
			hits = hits + 1
		}
	}

	return hits

}

func countPoints(hits int) int {
	return lut2[hits]
}

func Dayfour(in DayfourInput) int {
	sum := 0

	for _, line := range in.a {
		card := parseCard(line)
		hits := countHits(card)
		points := countPoints(hits)
		sum = sum + points

	}

	return sum
}
func DayfourPartTwo(in DayfourInput) int {
	lut := make(map[int]Card, len(in.a))
	var cards []Card
	for _, line := range in.a {
		card := parseCard(line)
		hits := countHits(card)
		card.hits = hits
		lut[card.id] = card
		cards = append(cards, card)
	}
	for j := 0; j < len(cards); j++ {
		card := cards[j]
		if card.hits > 0 {
			for i := card.id + 1; i <= card.id+card.hits; i++ {
				cards = append(cards, lut[i])
			}
		}
	}

	return len(cards)
}
func main() {
}
