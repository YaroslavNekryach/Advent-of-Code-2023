package day15

import (
	"advent-of-code/help"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

const day = 15

func Run() {
	input := Parse(help.GetInput(day))
	part1Result := Part1(input)
	part2Result := Part2(input)

	fmt.Println("Part1", part1Result)
	fmt.Println("Part2", part2Result)
}

type Lense struct {
	lable string
	lense int
}

func (l Lense) Hash() int {
	return Hash(l.lable)
}

type Input = []string

func Hash(s string) int  {
	result := 0
	chars := []rune(s)
	for _, c := range chars {
		result += int(c)
		result *= 17
		result %= 256
	}
	return result
}
func Part1(input Input) string {
	result := 0
	for _, s := range input {
		result += Hash(s)
	}
	return strconv.Itoa(result)
}

func Part2(input Input) string {
	result := 0
	shelves := make(map[int][]Lense)
	for _, s := range input {
		ss := strings.Split(s, "=")
		if len(ss) == 2 {
			n, _ := strconv.Atoi(ss[1])
			lense := Lense{
				lable: ss[0],
				lense: n,
			}
			hash := lense.Hash()
			if shelves[hash] == nil {
				shelves[hash] = make([]Lense, 0)
			}
			i := slices.IndexFunc(shelves[hash], func(l Lense) bool {
				return l.lable == lense.lable
			})
			if (i >= 0) {
				slices.Replace(shelves[hash], i, i+ 1, lense)
			} else {
				shelves[hash] = append(shelves[hash], lense)
			}
		} else {
			ss = strings.Split(s, "-")
			lable := ss[0]
			hash := Hash(lable)
			
			i := slices.IndexFunc(shelves[hash], func(l Lense) bool {
				return l.lable == lable
			})
			if (i >= 0) {
				shelves[hash] = slices.Delete(shelves[hash], i, i + 1)
			}
		}
	}
	
	for n, shelf := range shelves {
		for i, lense := range shelf {
			result += (n + 1) * (i + 1) * lense.lense
		}
	}
	return strconv.Itoa(result)
}

func Parse(input string) Input {
	return strings.Split(input, ",")
}
