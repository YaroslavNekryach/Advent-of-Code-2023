package day6

import (
	"advent-of-code/help"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const day = 6

type Race struct {
	time     int
	distance int
}

func (race Race) BetterCount() int {
	D := math.Sqrt(math.Pow(float64(race.time), 2) - 4*float64(race.distance))
	minF := (float64(race.time) - D) / 2
	maxF := (float64(race.time) + D) / 2
	min := int(math.Floor(minF) + 1)
	max := int(math.Ceil(maxF) - 1)

	return (max - min + 1)
}

type Input = []Race

func Run() {
	input := Parse(help.GetInput(day))
	part1Result := Part1(input)
	part2Result := Part2(input)

	fmt.Println("Part1", part1Result)
	fmt.Println("Part2", part2Result)
}

func Part1(input Input) string {
	result := 1
	for _, race := range input {
		result *= race.BetterCount()
	}
	return strconv.Itoa(result)
}

func Part2(input Input) string {
	timeString := ""
	distanceString := ""
	for _, race := range input {
		timeString += strconv.Itoa(race.time)
		distanceString += strconv.Itoa(race.distance)
	}

	time, _ := strconv.Atoi(timeString)
	distance, _ := strconv.Atoi(distanceString)
	race := Race{time, distance}

	return strconv.Itoa(race.BetterCount())
}

func Parse(input string) Input {
	split := strings.Split(input, "\n")
	timeString, distanceString := split[0], split[1]
	numReg := regexp.MustCompile(`\d+`)
	timeMatch := numReg.FindAllString(timeString, -1)
	distanceMatch := numReg.FindAllString(distanceString, -1)

	result := make([]Race, 0)

	for i := 0; i < len(timeMatch); i++ {
		time, _ := strconv.Atoi(timeMatch[i])
		distance, _ := strconv.Atoi(distanceMatch[i])
		result = append(result, Race{time, distance})
	}

	return result
}
