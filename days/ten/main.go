package main

import (
	"log"
	"sync"
)

type DaytenInput struct {
	a []string
}

type Board struct {
	tiles           [][]int
	tilesNoStartPos [][]int
	original        []string
	distances       [][]int
	startingPos     StartPos
    maxDistance int
}
type StartPos struct {
	i int
	j int
}

const (
	PIPE_VERTICAL       int       = 1
	PIPE_HORIZONTAL     int       = 2
	BEND_L              int       = 3
	BEND_J              int       = 4
	BEND_7              int       = 5
	BEND_F              int       = 6
	GROUND              int       = 7
	STARTING_POS        int       = 8
	SYM_PIPE_VERTICAL   string    = "|"
	SYM_PIPE_HORIZONTAL string    = "-"
	SYM_BEND_L          string    = "L"
	SYM_BEND_J          string    = "J"
	SYM_BEND_7          string    = "7"
	SYM_BEND_F          string    = "F"
	SYM_GROUND          string    = "."
	SYM_STARTING_POS    string    = "S"
	LEFT                Direction = "LEFT"
	DOWN                Direction = "DOWN"
	RIGHT               Direction = "RIGHT"
	UP                  Direction = "UP"
)

var directions = []Direction{LEFT, DOWN, UP, RIGHT}

type Direction string

func CanMove(direction Direction, from int, to int) bool {
	switch direction {
	case LEFT:
		return canMoveLeft(from, to)
	case DOWN:
		return canMoveDown(from, to)
	case UP:
		return canMoveUp(from, to)
	case RIGHT:
		return canMoveRight(from, to)
	default:
		log.Fatalf("Could not parse direction %v", direction)
		return false

	}
}
func canMoveLeft(from, to int) bool {
	fromPipeHorizontal := from == PIPE_HORIZONTAL
	fromBendJ := from == BEND_J
	fromBend7 := from == BEND_7
	fromAllowed := fromPipeHorizontal || fromBendJ || fromBend7
	fromStartingPos := from == STARTING_POS
	fromAllowed = fromAllowed || fromStartingPos
	toBendF := to == BEND_F
	toBendL := to == BEND_L
	toPipeHorizontal := to == PIPE_HORIZONTAL
	toAllowed := toBendF || toBendL || toPipeHorizontal
	return toAllowed && fromAllowed
}
func canMoveRight(from, to int) bool {
	fromPipeHorizontal := from == PIPE_HORIZONTAL
	fromBendF := from == BEND_F
	fromBendL := from == BEND_L
	fromAllowed := fromPipeHorizontal || fromBendF || fromBendL
	fromStartingPos := from == STARTING_POS
	fromAllowed = fromAllowed || fromStartingPos
	toBendJ := to == BEND_J
	toBend7 := to == BEND_7
	toPipeHorizontal := to == PIPE_HORIZONTAL
	toAllowed := toBendJ || toBend7 || toPipeHorizontal
	return toAllowed && fromAllowed
}
func canMoveUp(from, to int) bool {
	fromBendJ := from == BEND_J
	fromBendL := from == BEND_L
	fromPipeVertical := from == PIPE_VERTICAL
	fromAllowed := fromPipeVertical || fromBendJ || fromBendL
	fromStartingPos := from == STARTING_POS
	fromAllowed = fromAllowed || fromStartingPos
	toBend7 := to == BEND_7
	toBendF := to == BEND_F
	toPipeVertical := to == PIPE_VERTICAL
	toAllowed := toBendF || toBend7 || toPipeVertical
	return toAllowed && fromAllowed
}
func canMoveDown(from, to int) bool {
	fromPipeVertical := from == PIPE_VERTICAL
	fromBend7 := from == BEND_7
	fromBendF := from == BEND_F
	fromAllowed := fromPipeVertical || fromBendF || fromBend7
	fromStartingPos := from == STARTING_POS
	fromAllowed = fromAllowed || fromStartingPos
	toBendJ := to == BEND_J
	toBendL := to == BEND_L
	toPipeVertical := to == PIPE_VERTICAL
	toAllowed := toBendJ || toBendL || toPipeVertical
	return toAllowed && fromAllowed
}

func parseLine(line string) []int {
	newLine := make([]int, len(line))
	for index, char := range line {
		switch string(char) {
		case SYM_PIPE_VERTICAL:
			newLine[index] = PIPE_VERTICAL
		case SYM_PIPE_HORIZONTAL:
			newLine[index] = PIPE_HORIZONTAL
		case SYM_BEND_L:
			newLine[index] = BEND_L
		case SYM_BEND_J:
			newLine[index] = BEND_J
		case SYM_BEND_7:
			newLine[index] = BEND_7
		case SYM_BEND_F:
			newLine[index] = BEND_F
		case SYM_GROUND:
			newLine[index] = GROUND
		case SYM_STARTING_POS:
			newLine[index] = STARTING_POS
		default:
			log.Fatalf("Could not parse %v symbol in game input\n", char)
		}

	}
	return newLine

}
func parseLines(lines []string) [][]int {
	tiles := make([][]int, len(lines))
	for index, line := range lines {
		tiles[index] = parseLine(line)
	}
	return tiles
}

func parseInput(lines []string) Board {
	var board Board
	board.tiles = parseLines(lines)
	board.original = lines
	board.distances = make([][]int, len(lines))
	for i := 0; i < len(board.distances); i++ {
		board.distances[i] = make([]int, len(lines[i]))
	}
	return board
}

func (b *Board) determineStartPosTile() {
	tilesNoStartPos := make([][]int, len(b.tiles))
	for i := 0; i < len(b.tiles); i++ {
		newLine := make([]int, len(b.tiles[i]))
		for j := 0; j < len(b.tiles[i]); j++ {
			newLine[j] = b.tiles[i][j]
			b.distances[i][j] = -2
			if b.tiles[i][j] == STARTING_POS {
				b.startingPos = StartPos{i: i, j: j}
				b.distances[i][j] = 0
			}
			if b.tiles[i][j] == GROUND {
				b.distances[i][j] = -1
			}
		}
		tilesNoStartPos[i] = newLine
	}
	b.tilesNoStartPos = tilesNoStartPos
}

func walkIntoDirection(direction Direction, currentPos StartPos, wg *sync.WaitGroup, b *Board) {
	defer wg.Done()
	currentSymbol := b.tiles[currentPos.i][currentPos.j]
	currentDistance := b.distances[currentPos.i][currentPos.j]
	newI := currentPos.i
	newJ := currentPos.j
	var availableDirections []Direction
	switch direction {
	case LEFT:
		availableDirections = []Direction{LEFT, UP, DOWN}
		newJ = currentPos.j - 1
		if newJ >= 0 && CanMove(LEFT, currentSymbol, b.tiles[newI][newJ]) {
			// log.Printf("Trying to move %v from %v to %v\n", direction, string(b.original[currentPos.i][currentPos.j]), string(b.original[newI][newJ]))
			targetDistance := b.distances[newI][newJ]
			if targetDistance > currentDistance+1 || targetDistance < 0 {
				b.distances[newI][newJ] = currentDistance + 1
				for _, newDirection := range availableDirections {
					newPos := StartPos{i: newI, j: newJ}
					wg.Add(1)
					go walkIntoDirection(newDirection, newPos, wg, b)
				}
			}
		}
	case RIGHT:
		availableDirections = []Direction{RIGHT, UP, DOWN}
		newJ = currentPos.j + 1
		if newJ < len(b.distances[currentPos.i]) && CanMove(RIGHT, currentSymbol, b.tiles[newI][newJ]) {
			// log.Printf("Trying to move %v from %v to %v\n", direction, string(b.original[currentPos.i][currentPos.j]), string(b.original[newI][newJ]))
			targetDistance := b.distances[newI][newJ]
			if targetDistance > currentDistance+1 || targetDistance < 0 {
				b.distances[newI][newJ] = currentDistance + 1
				for _, newDirection := range availableDirections {
					newPos := StartPos{i: newI, j: newJ}
					wg.Add(1)
					go walkIntoDirection(newDirection, newPos, wg, b)
				}
			}
		}
	case UP:
		availableDirections = []Direction{RIGHT, LEFT, UP}
		newI = currentPos.i - 1
		if newI >= 0 && CanMove(UP, currentSymbol, b.tiles[newI][newJ]) {
			// log.Printf("Trying to move %v from %v to %v\n", direction, string(b.original[currentPos.i][currentPos.j]), string(b.original[newI][newJ]))
			targetDistance := b.distances[newI][newJ]
			if targetDistance > currentDistance+1 || targetDistance < 0 {
				b.distances[newI][newJ] = currentDistance + 1
				for _, newDirection := range availableDirections {
					newPos := StartPos{i: newI, j: newJ}
					wg.Add(1)
					go walkIntoDirection(newDirection, newPos, wg, b)
				}
			}
		}
	case DOWN:
		availableDirections = []Direction{RIGHT, LEFT, DOWN}
		newI = currentPos.i + 1
		if newI < len(b.distances) && CanMove(DOWN, currentSymbol, b.tiles[newI][newJ]) {
			// log.Printf("Trying to move %v from %v to %v\n", direction, string(b.original[currentPos.i][currentPos.j]), string(b.original[newI][newJ]))
			targetDistance := b.distances[newI][newJ]
			if targetDistance > currentDistance+1 || targetDistance < 0 {
				b.distances[newI][newJ] = currentDistance + 1
				for _, newDirection := range availableDirections {
					newPos := StartPos{i: newI, j: newJ}
					wg.Add(1)
					go walkIntoDirection(newDirection, newPos, wg, b)
				}
			}
		}
	}
}
func (b *Board) determineDistances() {

	var wg sync.WaitGroup
	currentPos := b.startingPos
    maxDistance := -2
	for _, direction := range directions {
		wg.Add(1)
		innerDirection := direction
		go walkIntoDirection(innerDirection, currentPos, &wg, b)
	}
	wg.Wait()
	for i := 0; i < len(b.distances); i++ {
		for j := 0; j < len(b.distances[i]); j++ {
            if b.distances[i][j] > maxDistance{
                maxDistance = b.distances[i][j]
            }
		}
	}

	// for _, direction := range directions {
	// 	wg.Add(1)
	// 	innerDirection := direction
	// 	go walkIntoDirection(innerDirection, currentPos, &wg, b)
	// 	wg.Wait()
	// }
        b.maxDistance = maxDistance

}

func Dayten(in DaytenInput) int {
	board := parseInput(in.a)
	board.determineStartPosTile()
	board.determineDistances()
	// log.Printf("Distances: %v\n", board.distances)
	// log.Printf("Parsed Board %v \n", board)
	return board.maxDistance
}
func main() {
}
