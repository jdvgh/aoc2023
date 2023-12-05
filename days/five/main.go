package main

import (
	"slices"
	"strconv"
	"strings"
)

type DayfiveInput struct {
	a []string
}

type Almanac struct {
	seeds []int
	maps  []Mapping
}

type Mapping struct {
	name        string
	conversions []Conversion
}

type Conversion struct {
	destinationRangeStart int
	sourceRangeStart      int
	length                int
}

func (a Almanac) convertSeeds(pairs bool) []int {
	var newSeedNumbers []int
	lowestNumber := 226172555
	if !pairs {
		for _, seed := range a.seeds {
			newSeedNumber := seed
			for _, mapping := range a.maps {
				newSeedNumber = mapping.convertSeed(newSeedNumber)
			}
			newSeedNumbers = append(newSeedNumbers, newSeedNumber)
			if newSeedNumber < lowestNumber {
				lowestNumber = newSeedNumber
			}
		}
	} else {
		if pairs {
			for i := 0; i < len(a.seeds)-1; i = i + 2 {
				start := a.seeds[i]
				length := a.seeds[i+1]
				for j := start; j < start+length; j++ {
					newSeedNumber := j
					for _, mapping := range a.maps {
						newSeedNumber = mapping.convertSeed(newSeedNumber)
					}
					if newSeedNumber < lowestNumber {
						lowestNumber = newSeedNumber
					}
				}
			}
            newSeedNumbers = append(newSeedNumbers, lowestNumber)
		}
	}
	slices.Sort(newSeedNumbers)
	return newSeedNumbers
}

func (m Mapping) convertSeed(seed int) int {
	for _, conversion := range m.conversions {
		newSeedNumber := conversion.convertSeed(seed)
		if newSeedNumber >= 0 {
			return newSeedNumber
		}
	}
	return seed
}

func (c Conversion) convertSeed(seed int) int {
	seedPos := seed - c.sourceRangeStart
	if seedPos < c.length && seed >= c.sourceRangeStart {
		return c.destinationRangeStart + seedPos
	}
	return -1
}

func parseAlmanac(lines []string) Almanac {
	var lineBuffer []string
	var almanac Almanac
	for _, line := range lines[0:2] {
		if strings.Contains(line, "seeds") {
			almanac.seeds = parseSeeds(line)
		}
	}
	for _, line := range lines[2:] {
		if line != "" {
			lineBuffer = append(lineBuffer, line)
		} else {
			mapping := parseMappings(lineBuffer)
			almanac.maps = append(almanac.maps, mapping)
			lineBuffer = []string{}
		}
	}
	mapping := parseMappings(lineBuffer)
	almanac.maps = append(almanac.maps, mapping)
	return almanac

}

func parseMappings(lines []string) Mapping {
	var mappings Mapping
	mappings.name = lines[0]
	for _, line := range lines[1:] {

		mappings.conversions = append(mappings.conversions, parseMappingLine(line))

	}
	return mappings

}

func parseSeeds(line string) []int {
	trimmedLine := strings.TrimSpace(line)
	splitValues := strings.Split(strings.TrimSpace(strings.Split(trimmedLine, ":")[1]), " ")
	var seeds []int
	for _, seed := range splitValues {
		res, err := strconv.Atoi(seed)
		if err != nil {
		} else {
			seeds = append(seeds, res)
		}
	}
	return seeds
}

func parseMappingLine(line string) Conversion {
	trimmedLine := strings.TrimSpace(line)
	splitValues := strings.Split(trimmedLine, " ")

	destinationRangeStart, err := strconv.Atoi(splitValues[0])
	if err != nil {
	}
	sourceRangeStart, err := strconv.Atoi(splitValues[1])
	if err != nil {
	}
	length, err := strconv.Atoi(splitValues[2])
	if err != nil {
	}
	return Conversion{
		destinationRangeStart: destinationRangeStart,

		sourceRangeStart: sourceRangeStart,
		length:           length,
	}
}

func Dayfive(in DayfiveInput) int {
	almanac := parseAlmanac(in.a)
	newSeeds := almanac.convertSeeds(false)
	return newSeeds[0]
}
func DayfivePartTwo(in DayfiveInput) int {
	almanac := parseAlmanac(in.a)
	newSeeds := almanac.convertSeeds(true)
	return newSeeds[0]
}
func main() {
}
