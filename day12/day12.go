package day12

import (
	"advent-of-code/help"
	"fmt"
	"strconv"
	"strings"
)

const day = 12

type Line struct {
	 need []int
	 line string
}

func (line Line) Key() string {
	needsString := make([]string, 0)
	for _, n := range line.need {
		needsString = append(needsString, strconv.Itoa(n))
	}
	return fmt.Sprintf("%s_%s", line.line, strings.Join(needsString, "_"))
}

func (line Line) Count(cache *map[string]int) int {
	key := line.Key()
	
	value, found := (*cache)[key]
	if found {
		return value
	}
	if found {
		return value
	}
	if len(line.need) == 0 {
		for i := 0; i < len(line.line); i++ {
			if line.line[i:i+1] == "#" {
				return 0
			}
		}
		return 1
	}
	
	item := line.need[0]
	count := 0
	next: for i := 0; i <= len(line.line) - item; i++ {
		if i > 0 && line.line[i-1:i] == "#" {
			return count
		}
		for j := 0; j < item; j++ {
			sub := line.line[i+j:i+j+1]
			if (sub == ".") {
				continue next
			}
		}
		if i+item < len(line.line) && line.line[i+item:i+item + 1] == "#" {
			continue next
		}
		nextNeed := make([]int, 0)
		if len(line.need) > 1 {
			nextNeed = line.need[1:]
		}
		nextL := ""
		if len(line.line) > i + item + 1 {
			nextL = line.line[i + item + 1:]
		}
		nextLine := Line{
			need: nextNeed,
			line: nextL,
		}
		nextCount := nextLine.Count(cache)
		if (len(line.need) > 1) {
			(*cache)[nextLine.Key()] = nextCount
		}
		count += nextCount
	}
	return count
}

type Input  = []Line

func Run() {
	input := Parse(help.GetInput(day))
	part1Result := Part1(input)
	part2Result := Part2(input)

	fmt.Println("Part1", part1Result)
	fmt.Println("Part2", part2Result)
}

func Part1(input Input) string {
	result := 0
	cache := make(map[string]int)
	
	for _, line := range input {
		result += line.Count(&cache)
	}
	return strconv.Itoa(result)
}

func Part2(input Input) string {
	result := 0
	cache := make(map[string]int)
	for _, line := range input {
		l := line.line
		need := line.need
		for i := 0; i < 4; i++ {
			line.line += "?" + l
			line.need = append(line.need, need...)
		}
		count := line.Count(&cache)
		result += count
	}
	return strconv.Itoa(result)
}

func Parse(input string) Input {
	result := make(Input, 0)
	for _, row := range strings.Split(input, "\n") {
		lineParts := strings.Split(row, " ")
		line, needString := lineParts[0], lineParts[1]
		
		needsString := strings.Split(needString, ",")
		need := make([]int, 0)
		for _, s := range needsString {
			n, _ := strconv.Atoi(s)
			need = append(need, n)
		}
		result = append(result, Line{
			need,
			line,
		})
	}
	
	return result
}
