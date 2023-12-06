package day3

import (
	"advent-of-code/help"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const day = 3

type Pos struct {
	x int
	y int
}

type Num struct {
	value int
	pos   Pos
}

func (num Num) Len() int {
	return len(strconv.Itoa(num.value))
}

func (num Num) AroundPos() []Pos {
	numLen := num.Len()
	result := make([]Pos, numLen*2+2)
	for i := -1; i < numLen+1; i++ {
		result = append(result, Pos{x: i + num.pos.x, y: num.pos.y - 1})
		result = append(result, Pos{x: i + num.pos.x, y: num.pos.y + 1})
	}
	result = append(result, Pos{x: num.pos.x - 1, y: num.pos.y})
	result = append(result, Pos{x: num.pos.x + numLen, y: num.pos.y})
	return result
}

type Sym struct {
	sym  string
	pos  Pos
	nums []int
}

type Input struct {
	nums []Num
	syms []Sym
}

func Run() {
	input := Parse(help.GetInput(day))
	part1Result := Part1(input)
	part2Result := Part2(input)

	fmt.Println("Part1", part1Result)
	fmt.Println("Part2", part2Result)
}

func Part1(input Input) string {
	result := 0
	for _, num := range input.nums {
		aroundPos := num.AroundPos()
	posLoop:
		for _, pos := range aroundPos {
			for _, sym := range input.syms {
				if pos.x == sym.pos.x && pos.y == sym.pos.y {
					result += num.value
					break posLoop
				}
			}
		}
	}
	return strconv.Itoa(result)
}

func Part2(input Input) string {
	result := 0
	for _, num := range input.nums {
		aroundPos := num.AroundPos()
		for _, pos := range aroundPos {
			for i := 0; i < len(input.syms); i++ {
				sym := &input.syms[i]
				if pos.x == sym.pos.x && pos.y == sym.pos.y {
					sym.nums = append(sym.nums, num.value)
				}
			}
		}
	}
	for _, sym := range input.syms {
		if sym.sym == "*" && len(sym.nums) == 2 {
			result += sym.nums[0] * sym.nums[1]
		}
	}
	return strconv.Itoa(result)
}

func Parse(input string) Input {
	nums := make([]Num, 0)
	syms := make([]Sym, 0)
	for y, line := range strings.Split(input, "\n") {
		numReg := regexp.MustCompile(`\d+`)
		numRegRes := numReg.FindAllStringIndex(line, -1)
		for _, match := range numRegRes {
			value, _ := strconv.Atoi(line[match[0]:match[1]])
			pos := Pos{
				x: match[0],
				y: y,
			}
			nums = append(nums, Num{value, pos})
		}
		symReg := regexp.MustCompile(`[^\d\.]`)
		symRegRes := symReg.FindAllStringIndex(line, -1)
		for _, match := range symRegRes {
			sym := line[match[0]:match[1]]
			pos := Pos{
				x: match[0],
				y: y,
			}
			syms = append(syms, Sym{sym: sym, pos: pos, nums: make([]int, 0)})
		}

	}

	return Input{
		nums,
		syms,
	}
}
