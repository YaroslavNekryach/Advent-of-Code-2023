package day4

import (
	"advent-of-code/help"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const day = 4

type Card struct {
	id   int
	wins []int
	nums []int
}

func (card Card) GetWins() []int {
	wins := make([]int, 0)
	for _, win := range card.wins {
		for _, num := range card.nums {
			if win == num {
				wins = append(wins, num)
				break
			}
		}
	}
	return wins
}

type Input = []Card

func Run() {
	input := Parse(help.GetInput(day))
	part1Result := Part1(input)
	part2Result := Part2(input)

	fmt.Println("Part1", part1Result)
	fmt.Println("Part2", part2Result)
}

func Part1(input Input) string {
	result := 0
	for _, card := range input {
		winCount := len(card.GetWins())
		if winCount > 0 {
			result += int(math.Pow(2.0, float64(winCount-1)))
		}
	}
	return strconv.Itoa(result)
}

func Part2(input Input) string {
	result := 0
	cardMap := make(map[int]int)
	for i := 0; i < len(input); i++ {
		cardMap[i] = 1
	}

	for i, card := range input {
		winCount := len(card.GetWins())
		for v := 0; v < winCount; v++ {
			cardMap[i+v+1] += cardMap[i]
		}
	}

	for i := 0; i < len(cardMap); i++ {
		result += cardMap[i]
	}

	return strconv.Itoa(result)
}

func Parse(input string) Input {
	cards := make([]Card, 0)
	for _, line := range strings.Split(input, "\n") {
		numReg := regexp.MustCompile(`\d+`)

		cardSplit := strings.Split(line, ": ")
		idString, allNumbersSting := cardSplit[0], cardSplit[1]
		numbersSplit := strings.Split(allNumbersSting, " | ")
		winsString, numsString := numbersSplit[0], numbersSplit[1]

		idMatch := numReg.FindAllString(idString, 1)
		id, _ := strconv.Atoi(idMatch[0])
		winsMatch := numReg.FindAllString(winsString, -1)
		numsMatch := numReg.FindAllString(numsString, -1)

		wins := make([]int, 0)
		nums := make([]int, 0)

		for _, winString := range winsMatch {
			win, _ := strconv.Atoi(winString)
			wins = append(wins, win)
		}

		for _, numString := range numsMatch {
			num, _ := strconv.Atoi(numString)
			nums = append(nums, num)
		}
		cards = append(cards, Card{id, wins, nums})
	}

	return cards
}
