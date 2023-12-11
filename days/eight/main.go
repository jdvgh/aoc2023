package main

import (
	"sync"
)

type DayeightInput struct {
	a []string
}

type InputMap struct {
	instructions  []int
	mappings      map[string]MapContent
	start         string
	goal          string
	startingNodes []string
}
type MapContent struct {
	current string
	left    string
	right   string
}

const (
	L   int    = 1
	R   int    = 2
	AAA string = "AAA"
	ZZZ string = "ZZZ"
)

func parseMap(lines []string, partTwo bool) InputMap {
	instructions := parseInstructions(lines[0])

	mappings := parseMappings(lines[2:])

	start := parseMapping(lines[2]).current
	goal := parseMapping(lines[len(lines)-1]).current
	startingNodes := getStartingNodes(mappings)
	if !partTwo {

		start = AAA
		goal = ZZZ
	}

	inputMap := InputMap{
		instructions:  instructions,
		mappings:      mappings,
		startingNodes: startingNodes,
		start:         start,
		goal:          goal,
	}
	return inputMap

}
func getStartingNodes(mappings map[string]MapContent) []string {
	var startingNodes []string
	for key := range mappings {
		if string(key[2]) == "A" {
			startingNodes = append(startingNodes, key)
		}
	}
	return startingNodes
}

func parseInstructions(line string) []int {
	var instructions []int
	for _, char := range line {
		switch string(char) {
		case "L":
			instructions = append(instructions, L)
		case "R":
			instructions = append(instructions, R)

		}

	}
	return instructions

}
func parseMappings(lines []string) map[string]MapContent {
	mapContent := MapContent{}
	mapping := make(map[string]MapContent)
	for _, line := range lines {
		mapContent = parseMapping(line)
		mapping[mapContent.current] = mapContent
	}
	return mapping

}
func parseMapping(line string) MapContent {
	var mapContent MapContent

	mapContent.current = line[0:3]
	mapContent.left = line[7:10]
	mapContent.right = line[12:15]

	return mapContent
}

func (m InputMap) calculateSteps(partTwo bool) int {
	steps := 0
	found := false
	currentMapping := m.mappings[m.start]
	lenInstructions := len(m.instructions)
	targetMapping := ""
	if !partTwo {
		for !found {
			currentInstructions := m.instructions[steps%lenInstructions]
			if currentInstructions == L {
				targetMapping = currentMapping.left
			} else {
				targetMapping = currentMapping.right
			}
			steps = steps + 1
			currentMapping = m.mappings[targetMapping]
			if targetMapping == m.goal {
				found = true
			}
		}
	} else {
		var currentNodes []string
		for _, startNode := range m.startingNodes {
			currentNodes = append(currentNodes, startNode)
		}
		var wg sync.WaitGroup
		var solutionIntervals []int
		solutionIntervals = make([]int, len(currentNodes))
		for index, current := range currentNodes {
			wg.Add(1)
			innerIndex := index
			innerCurrent := current
			go func() {
				defer wg.Done()
				innerSteps := 0
				innerFound := false
				for !innerFound {
					currentInstructions := m.instructions[innerSteps%lenInstructions]
					innerSteps = innerSteps + 1
					currentInnerMapping := m.mappings[innerCurrent]
					if currentInstructions == L {
						innerCurrent = currentInnerMapping.left
					} else {
						innerCurrent = currentInnerMapping.right
					}
					if string(innerCurrent[2]) == "Z" {
						solutionIntervals[innerIndex] = innerSteps
						innerFound = true
					}
				}
			}()
		}
		wg.Wait()
		steps = LCM(solutionIntervals[0], solutionIntervals[1], solutionIntervals[2:]...)
	}

	return steps

}

// greatest common divisor (GCD) via Euclidean algorithm
// taken from : https://go.dev/play/p/SmzvkDjYlb
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
func Dayeight(in DayeightInput) int {
	partTwo := false
	inputMap := parseMap(in.a, partTwo)
	res := inputMap.calculateSteps(partTwo)
	return res
}
func DayeightPartTwo(in DayeightInput) int {
	partTwo := true
	inputMap := parseMap(in.a, partTwo)
	res := inputMap.calculateSteps(partTwo)
	return res
}
func main() {
}
