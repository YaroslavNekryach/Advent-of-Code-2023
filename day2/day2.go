package day2

import (
	"advent-of-code/help"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const day = 2

type Set struct {
	red   int
	green int
	blue  int
}

type Game struct {
	id  int
	set []Set
}

type Input = []Game

func Run() {
	input := Parse(help.GetInput(day))
	part1Result := Part1(input)
	part2Result := Part2(input)

	fmt.Println("Part1", part1Result)
	fmt.Println("Part2", part2Result)
}

func Part1(input Input) string {
	const redLimit = 12
	const greenLimit = 13
	const blueLimit = 14
	result := 0

	for _, game := range input {
		ok := true
		for _, set := range game.set {
			if set.red > redLimit || set.blue > blueLimit || set.green > greenLimit {
				ok = false
				break
			}
		}
		if ok {
			result += game.id
		}
	}

	return strconv.Itoa(result)
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Part2(input Input) string {
	result := 0

	for _, game := range input {
		min := Set{}
		for _, set := range game.set {
			min.red = Max(min.red, set.red)
			min.green = Max(min.green, set.green)
			min.blue = Max(min.blue, set.blue)
		}
		result += min.red * min.green * min.blue
	}

	return strconv.Itoa(result)
}

func Parse(input string) Input {
	result := make(Input, 0)
	gameReg, _ := regexp.Compile(`Game (\d+)`)
	redReg, _ := regexp.Compile(`(\d+) red`)
	greenReg, _ := regexp.Compile(`(\d+) green`)
	blueReg, _ := regexp.Compile(`(\d+) blue`)
	for _, games := range strings.Split(input, "\n") {
		gamesPart := strings.Split(games, ": ")
		game, setsStr := gamesPart[0], gamesPart[1]
		gameMatch := gameReg.FindAllStringSubmatch(game, 1)
		gameId, _ := strconv.Atoi(gameMatch[0][1])
		sets := make([]Set, 0)
		for _, setStr := range strings.Split(setsStr, "; ") {
			redMatch := redReg.FindAllStringSubmatch(setStr, 1)
			greenMatch := greenReg.FindAllStringSubmatch(setStr, 1)
			blueMatch := blueReg.FindAllStringSubmatch(setStr, 1)
			set := Set{}
			if len(redMatch) > 0 {
				set.red, _ = strconv.Atoi(redMatch[0][1])
			}
			if len(greenMatch) > 0 {
				set.green, _ = strconv.Atoi(greenMatch[0][1])
			}
			if len(blueMatch) > 0 {
				set.blue, _ = strconv.Atoi(blueMatch[0][1])
			}
			sets = append(sets, set)
		}
		result = append(result, Game{
			id:  gameId,
			set: sets,
		})
	}
	return result
}
