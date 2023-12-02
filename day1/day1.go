package day1

import (
	"advent-of-code/help"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const day = 1

type Input = []string

func Run() {
	input := Parse(help.GetInput(day))
	part1Result := Part1(input)
	part2Result := Part2(input)

	fmt.Println("Part1", part1Result)
	fmt.Println("Part2", part2Result)
}

func Part1(input Input) string {
	sum := 0
	for _, line := range input {

		var first int
		var last int
		gotFirst := false
		gotLast := false
		for _, sym := range strings.Split(line, "") {
			if sym >= "0" && sym <= "9" {
				if gotFirst {
					last, _ = strconv.Atoi(sym)
					gotLast = true
				} else {
					first, _ = strconv.Atoi(sym)
					gotFirst = true
				}
			}
		}
		if !gotLast {
			last = first
		}

		sum += first*10 + last
	}
	return strconv.Itoa(sum)
}

func Part2(input Input) string {
	words := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	sum := 0
	for _, line := range input {

		var firstIndex int = math.MaxInt
		var lastIndex int = -1
		var first int
		var last int
		for num, word := range words {
			var i int
			i = strings.Index(line, word)

			if i >= 0 && i < firstIndex {
				firstIndex = i
				first = num
			}

			i = strings.LastIndex(line, word)

			if i >= 0 && i > lastIndex {
				lastIndex = i
				last = num
			}

			i = strings.Index(line, strconv.Itoa(num))

			if i >= 0 && i < firstIndex {
				firstIndex = i
				first = num
			}

			i = strings.LastIndex(line, strconv.Itoa(num))

			if i >= 0 && i > lastIndex {
				lastIndex = i
				last = num
			}

		}
		var val = first*10 + last
		// fmt.Printf("%v\n", val)
		sum += val
	}
	return strconv.Itoa(sum)
}

func Parse(input string) Input {
	return strings.Split(input, "\n")
	// var result Input
	// for _, elve := range strings.Split(input, "\n\n") {
	// 	var elves []int
	// 	items := strings.Split(elve, "\n")
	// 	for _, item := range items {
	// 		v, _ := strconv.Atoi(item)
	// 		elves = append(elves, v)
	// 	}
	// 	result = append(result, elves)
	// }
	// return result
}
