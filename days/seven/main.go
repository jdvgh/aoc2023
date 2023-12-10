package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

type DaysevenInput struct {
	a []string
}

type Game struct {
	inputs        []Input
	totalWinnings int
}

type Input struct {
	hand      Hand
	bid       int
	rank      int
	winAmount int
}

type Hand struct {
	cards         []int
	cardsUnparsed string
	cardAmountMap map[int]int
	handType      HandType
	elaborateRank int
	isPartTwo     bool
}

type isHandTypeFunc func(h Hand) bool

type HandType struct {
	name       HandTypeName
	rank       int
	isHandType isHandTypeFunc
}

type HandTypeName string

const (
	FIVE_OF_A_KIND  HandTypeName = "FIVE_OF_A_KIND"
	FOUR_OF_A_KIND  HandTypeName = "FOUR_OF_A_KIND"
	FULL_HOUSE      HandTypeName = "FULL_HOUSE"
	THREE_OF_A_KIND HandTypeName = "THREE_OF_A_KIND"
	TWO_PAIR        HandTypeName = "TWO_PAIR"
	ONE_PAIR        HandTypeName = "ONE_PAIR"
	HIGH_CARD       HandTypeName = "HIGH_CARD"
	JOKER           int          = 1
)

var HAND_TYPES = map[HandTypeName]HandType{
	FIVE_OF_A_KIND: HandType{FIVE_OF_A_KIND, 7, func(h Hand) bool {
		maxValue := 0
		maxValueNonJoker := 0
		for key, value := range h.cardAmountMap {
			if value > maxValue {
				maxValue = value
			}
			if value > maxValueNonJoker && key != JOKER {
				maxValueNonJoker = value
			}
		}
		if h.isPartTwo {
			jokerAmount, ok := h.cardAmountMap[JOKER]
			if ok {
				maxValue = maxValueNonJoker + jokerAmount
			}
		}
		return maxValue == 5
	}},
	FOUR_OF_A_KIND: HandType{FOUR_OF_A_KIND, 6, func(h Hand) bool {
		maxValue := 0
		maxValueNonJoker := 0
		for key, value := range h.cardAmountMap {
			if value > maxValue {
				maxValue = value
			}
			if value > maxValueNonJoker && key != JOKER {
				maxValueNonJoker = value
			}
		}
		if h.isPartTwo {
			jokerAmount, ok := h.cardAmountMap[JOKER]
			if ok {
				maxValue = maxValueNonJoker + jokerAmount
			}
		}
		return maxValue == 4
	}},
	FULL_HOUSE: HandType{FULL_HOUSE, 5, func(h Hand) bool {
		maxValue := 0
		minValue := 999
		maxValueNonJoker := 0
		minValueNonJoker := 999
		for key, value := range h.cardAmountMap {
			if value < minValue {
				minValue = value
			}
			if value > maxValue {
				maxValue = value
			}
			if value > maxValueNonJoker && key != JOKER {
				maxValueNonJoker = value
			}
			if value < minValueNonJoker && key != JOKER {
				minValueNonJoker = value
			}
		}
		if h.isPartTwo {
			jokerAmount, ok := h.cardAmountMap[JOKER]
			if ok {
				maxValue = maxValueNonJoker + jokerAmount
				minValue = minValueNonJoker
			}
		}
		return maxValue == 3 && minValue == 2
	}},

	THREE_OF_A_KIND: HandType{THREE_OF_A_KIND, 4, func(h Hand) bool {
		maxValue := 0
		minValue := 999
		maxValueNonJoker := 0
		minValueNonJoker := 999
		for key, value := range h.cardAmountMap {
			if value < minValue {
				minValue = value
			}
			if value > maxValue {
				maxValue = value
			}
			if value > maxValueNonJoker && key != JOKER {
				maxValueNonJoker = value
			}
			if value < minValueNonJoker && key != JOKER {
				minValueNonJoker = value
			}
		}
		if h.isPartTwo {
			jokerAmount, ok := h.cardAmountMap[JOKER]
			if ok {
				maxValue = maxValueNonJoker + jokerAmount
				minValue = minValueNonJoker
			}
		}
		return maxValue == 3 && minValue == 1
	}},

	TWO_PAIR: HandType{TWO_PAIR, 3, func(h Hand) bool {
		pairOneFound := false
		pairTwoFound := false
		singleFound := false
		maxValue := 0
		maxValueNonJoker := 0
		for key, value := range h.cardAmountMap {
			if value > maxValue {
				maxValue = value
			}
			if value > maxValueNonJoker && key != JOKER {
				maxValueNonJoker = value
			}
			if value == 1 {
				singleFound = true
			} else {
				if value == 2 && !pairOneFound {
					pairOneFound = true

				} else if value == 2 && !pairTwoFound {
					pairTwoFound = true
				}
			}
		}
		if h.isPartTwo {
			jokerAmount, ok := h.cardAmountMap[JOKER]
			if ok {
				maxValue = maxValueNonJoker + jokerAmount
			}
		}
		return singleFound && pairOneFound && pairTwoFound && maxValue == 2
	}},

	ONE_PAIR: HandType{ONE_PAIR, 2, func(h Hand) bool {
		onePairFound := false
		maxValue := 0
		maxValueNonJoker := 0
		for key, value := range h.cardAmountMap {
			if value > maxValue {
				maxValue = value
			}
			if value > maxValueNonJoker && key != JOKER {
				maxValueNonJoker = value
			}
			if value == 2 {
				if !onePairFound {
					onePairFound = true
				} else {
					return false
				}
			} else if value > 2 {
				return false
			}
		}
		if h.isPartTwo {
			jokerAmount, ok := h.cardAmountMap[JOKER]
			if ok {
				maxValue = maxValueNonJoker + jokerAmount
			}
			if !onePairFound && maxValue == 2 {
				onePairFound = true
			}
		}

		return onePairFound && maxValue == 2

	}},

	HIGH_CARD: HandType{HIGH_CARD, 1, func(h Hand) bool {
		maxValue := 0
		for _, value := range h.cardAmountMap {
			if value > maxValue {
				maxValue = value
			}
		}
		if h.isPartTwo {
			jokerAmount, ok := h.cardAmountMap[JOKER]
			if ok {
				for key, value := range h.cardAmountMap {
					if value == maxValue && key != JOKER {
						maxValue = maxValue + jokerAmount
					}
				}
			}
		}

		return maxValue == 1
	}},
}

func parseGame(lines []string, partTwo bool) Game {
	var game Game
	for _, line := range lines {
		input := parseInput(strings.TrimSpace(line), partTwo)
		game.inputs = append(game.inputs, input)
	}
	game.evaluateWinners()
	return game
}
func (g *Game) evaluateWinners() {
	rank := 1
	slices.SortFunc(g.inputs, func(a Input, b Input) int {
		if a.hand.elaborateRank > b.hand.elaborateRank {
			return 1
		} else if a.hand.elaborateRank < b.hand.elaborateRank {
			return -1
		}
		return 0
	})
	totalWinningsPost := 0
	rank = 1
	gamesOrder := ""
	for _, input := range g.inputs {
		input.rank = rank
		gamesOrder = gamesOrder + "\n" + fmt.Sprintf("%v", input.hand.elaborateRank) + "-" + input.hand.cardsUnparsed
		totalWinningsPost = totalWinningsPost + input.bid*input.rank
		rank = rank + 1
	}
	g.totalWinnings = totalWinningsPost
}
func parseInput(line string, partTwo bool) Input {
	inputSplit := strings.Split(line, " ")
	hand := parseHand(inputSplit[0], partTwo)
	bids, err := strconv.Atoi(inputSplit[1])
	if err != nil {
		log.Fatalf("Could not parse bid %v\n", inputSplit[1])
	}
	return Input{hand: hand, bid: bids}

}
func parseHand(cardString string, partTwo bool) Hand {
	var hand Hand
	hand.isPartTwo = partTwo
	hand.cardAmountMap = make(map[int]int)
	elaborateRank := ""
	hand.cardsUnparsed = cardString
	maxNonJoker := 0
	for _, numberChar := range cardString {
		res, err := strconv.Atoi(string(numberChar))
		if err != nil {
			switch string(numberChar) {
			case "A":
				res = 14
			case "K":
				res = 13
			case "Q":
				res = 12
			case "J":
				if partTwo {
					res = 1
				} else {
					res = 11
				}
			case "T":
				res = 10
			default:
				log.Fatalf("Could not parse card %v\n", numberChar)
			}
		}
		hand.cards = append(hand.cards, res)
		old, ok := hand.cardAmountMap[res]
		if !ok {
			old = 0
		}
		hand.cardAmountMap[res] = old + 1
		if old+1 > maxNonJoker && res != JOKER {
			maxNonJoker = res
		}
		elaborateRank = fmt.Sprintf("%v%02d", elaborateRank, res)
	}

	for _, handType := range HAND_TYPES {
		if handType.isHandType(hand) {
			hand.handType = handType
			tempRank, err := strconv.Atoi(strings.TrimSpace(fmt.Sprintf("%2d%v", handType.rank, elaborateRank)))
			if err != nil {
				log.Fatalf("Could not pare elaborate rank from %v\n", fmt.Sprintf("%2d%v", handType.rank, elaborateRank))
			}
			hand.elaborateRank = tempRank
			return hand
		}
	}
	return hand
}

func Dayseven(in DaysevenInput) int {
	game := parseGame(in.a, false)
	return game.totalWinnings
}
func DaysevenPartTwo(in DaysevenInput) int {
	game := parseGame(in.a, true)
	return game.totalWinnings
}
func main() {
}
