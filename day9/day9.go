package day9

import (
	"advent-of-code/help"
	"fmt"
	"strconv"
	"strings"
)

const day = 9

type Input = [][]int

func Run() {
	input := Parse(help.GetInput(day))
	part1Result := Part1(input)
	part2Result := Part2(input)

	fmt.Println("Part1", part1Result)
	fmt.Println("Part2", part2Result)
}

func Part1(input Input) string {
	result := 0
	
	for _, line := range input {
		data := [][]int{}
		lastLine := line
		next: for {
			nextLine := make([]int, len(lastLine) - 1)
			for i := 0; i < len(lastLine) - 1; i++ {
				nextLine[i] = lastLine[i + 1] - lastLine[i]
			}
			data = append(data, lastLine)
			lastLine = nextLine
			for _, v := range lastLine {
				if v != 0 {
					continue next
				}
			}
			break
		}
		for _, d := range data {
			result += d[len(d) - 1]
		}
	}
	return strconv.Itoa(result)
}

func Part2(input Input) string {
	newInput := make(Input, 0)
	for _, line := range input {
		for i, j := 0, len(line)-1; i < j; i, j = i+1, j-1 {
			line[i], line[j] = line[j], line[i]
		}
		newInput = append(newInput, line)
	}
	return Part1(newInput)
}

func Parse(input string) Input {
	result := make([][]int, 0)
	for i, line := range strings.Split(input, "\n") {
		result = append(result, make([]int, 0))
		for _, str := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(str)
			result[i] = append(result[i], num)
		}
	}

	return result
}
