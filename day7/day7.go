package day7

import (
	"advent-of-code/help"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const day = 7

var Cards = map[string]int{"A": 1, "K": 2, "Q": 3, "J": 4, "T": 5, "9": 6, "8": 7, "7": 8, "6": 9, "5": 10, "4": 11, "3": 12, "2": 13}

// var Cards2 = map[string]int{"A": 1, "K": 2, "Q": 3, "T": 4, "9": 5, "8": 6, "7": 7, "6": 8, "5": 9, "4": 10, "3": 11, "2": 12, "J": 13}

type Hand struct {
	cards []string
	bid   int
}

func (hand Hand) TypeRank1() int {
	count := make(map[string]int)

	for _, card := range hand.cards {
		count[card] += 1
	}
	values := make([]int, 0)
	for _, n := range count {
		values = append(values, n)
	}

	sort.SliceStable(values, func(i, j int) bool { return values[i] > values[j] })

	if values[0] == 5 {
		return 1
	}
	if values[0] == 4 {
		return 2
	}
	if values[0] == 3 && values[1] == 2 {
		return 3
	}
	if values[0] == 3 {
		return 4
	}
	if values[0] == 2 && values[1] == 2 {
		return 5
	}
	if values[0] == 2 {
		return 6
	}
	return 7
}

func (hand Hand) TypeRank2() int {
	count := make(map[string]int)

	for _, card := range hand.cards {
		count[card] += 1
	}
	j := count["J"]
	count["J"] = 0
	values := make([]int, 0)
	for _, n := range count {
		values = append(values, n)
	}

	sort.SliceStable(values, func(i, j int) bool { return values[i] > values[j] })

	if values[0]+j == 5 {
		return 1
	}
	if values[0]+j == 4 {
		return 2
	}
	if values[0]+j == 3 && values[1] == 2 {
		return 3
	}
	if values[0]+j == 3 {
		return 4
	}
	if values[0]+j == 2 && values[1] == 2 {
		return 5
	}
	if values[0]+j == 2 {
		return 6
	}
	return 7
}

type Input = []Hand

func Run() {
	input := Parse(help.GetInput(day))
	part1Result := Part1(input)
	part2Result := Part2(input)

	fmt.Println("Part1", part1Result)
	fmt.Println("Part2", part2Result)
}

func Part1(input Input) string {
	result := 0
	sort.SliceStable(input, func(i, j int) bool {
		hand1, hand2 := input[i], input[j]
		hand1Rank := hand1.TypeRank1()
		hand2Rank := hand2.TypeRank1()
		if hand1Rank != hand2Rank {
			return hand1Rank > hand2Rank
		}
		for i := 0; i < len(hand1.cards); i++ {
			if hand1.cards[i] != hand2.cards[i] {
				return Cards[hand1.cards[i]] > Cards[hand2.cards[i]]
			}
		}
		return false
	})
	for i, v := range input {
		result += (i + 1) * v.bid
	}
	return strconv.Itoa(result)
}

func Part2(input Input) string {
	result := 0
	Cards["J"] = 14
	sort.SliceStable(input, func(i, j int) bool {
		hand1, hand2 := input[i], input[j]
		hand1Rank := hand1.TypeRank2()
		hand2Rank := hand2.TypeRank2()
		if hand1Rank != hand2Rank {
			return hand1Rank > hand2Rank
		}
		for i := 0; i < len(hand1.cards); i++ {
			if hand1.cards[i] != hand2.cards[i] {
				return Cards[hand1.cards[i]] > Cards[hand2.cards[i]]
			}
		}
		return false
	})
	for i, v := range input {
		result += (i + 1) * v.bid
	}
	return strconv.Itoa(result)
}

func Parse(input string) Input {
	result := make([]Hand, 0)
	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, " ")

		cards := strings.Split(split[0], "")
		bid, _ := strconv.Atoi(split[1])

		result = append(result, Hand{cards, bid})
	}

	return result
}
