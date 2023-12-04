package main

import (
	"log"
	"strconv"
	"strings"
)

type DaytwoInput struct {
	a []string
	b CubeSet
}

type Game struct {
	id           int
	cubeSets     []CubeSet
	possible     bool
	minimumCubes CubeSet
}

type CubeSet struct {
	blue  int
	green int
	red   int
}

func Daytwo(in DaytwoInput) int {
	sum := 0
	for _, line := range in.a {
		game := parseGame(line)
		game.possible = checkGamePossible(game, in.b)
		if game.possible {
			sum += game.id
		}
	}
	return sum
}

func DaytwoPartTwo(in DaytwoInput) int {
	sum := 0
	for _, line := range in.a {
		game := parseGame(line)
		game.minimumCubes = checkMinimumCubes(game)
		sum += game.minimumCubes.blue * game.minimumCubes.green * game.minimumCubes.red
	}
	return sum
}

func checkMinimumCubes(game Game) CubeSet {
	minimumCubeSet := CubeSet{}
	for _, cubeSet := range game.cubeSets {
		if cubeSet.blue > minimumCubeSet.blue {
			minimumCubeSet.blue = cubeSet.blue
		}
		if cubeSet.green > minimumCubeSet.green {
			minimumCubeSet.green = cubeSet.green
		}
		if cubeSet.red > minimumCubeSet.red {
			minimumCubeSet.red = cubeSet.red
		}
	}
	return minimumCubeSet
}
func checkGamePossible(game Game, restrictions CubeSet) bool {
	possible := true
	for _, cubeSet := range game.cubeSets {
		possible = possible && cubeSet.blue <= restrictions.blue && cubeSet.green <= restrictions.green && cubeSet.red <= restrictions.red
		if !possible {
			return possible
		}
	}
	return possible
}

func parseCubes(cubeLine string) CubeSet {
	cubePulls := strings.Split(cubeLine, ",")
	cubeSet := CubeSet{}
	for _, cubePulls := range cubePulls {
		cubePullsTrimmed := strings.TrimSpace(cubePulls)
		cubePullsSplit := strings.Split(cubePullsTrimmed, " ")
		amount, err := strconv.Atoi(cubePullsSplit[0])
		if err != nil {
			log.Fatalf("Could not parse cube amount from %v\n", cubePullsSplit[0])
		}
		colour := cubePullsSplit[1]
		switch colour {
		case "blue":
			cubeSet.blue = amount
		case "green":
			cubeSet.green = amount
		case "red":
			cubeSet.red = amount
		}
	}
	return cubeSet
}
func parseGame(line string) Game {
	gameLineSplit := strings.Split(line, ":")
	gameIdRaw := gameLineSplit[0]
	games := strings.Split(gameLineSplit[1], ";")
	gameId, err := strconv.Atoi(string(strings.Split(gameIdRaw, "Game ")[1]))
	if err != nil {
		log.Fatalf("Could not parse gameId from %v\n", gameIdRaw)
	}
	var cubeSets []CubeSet
	cubeSet := CubeSet{}
	for _, gameLine := range games {
		cubeSet = parseCubes(gameLine)
		cubeSets = append(cubeSets, cubeSet)
	}
	game := Game{
		id:       gameId,
		cubeSets: cubeSets,
	}

	return game
}
func main() {
}
