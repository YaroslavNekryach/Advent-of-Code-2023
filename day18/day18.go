package day18

import (
	"advent-of-code/help"
	"fmt"
	"strconv"
	"strings"
)

const day = 18

func Run() {
	input := Parse(help.GetInput(day))
	part1Result := Part1(input)
	part2Result := Part2(input)

	fmt.Println("Part1", part1Result)
	fmt.Println("Part2", part2Result)
}

type Pos struct {
	x int
	y int
}
type Input = []Inst

type Dir int

const (
	R Dir = iota
	D
	L
	U
)

func (d Dir) String() string {
	if d == U {
		return "UP"
	}
	if d == D {
		return "DOWN"
	}
	if d == L {
		return "LEFT"
	}
	if d == R {
		return "RIGHT"
	}
	return ""
}

type Inst struct {
	dir   Dir
	steps int
	corol string
}

func MyPart1(input Input) string {
	pos := Pos{}
	shiftMap := map[Dir]Pos{
		U: {-1, 0},
		D: {1, 0},
		L: {0, -1},
		R: {0, 1},
	}

	m := make(map[Pos]bool)
	m[pos] = true
	for _, inst := range input {
		for i := 0; i < inst.steps; i++ {
			pos = Pos{
				pos.x + shiftMap[inst.dir].x,
				pos.y + shiftMap[inst.dir].y,
			}
			m[pos] = true
		}
	}
	li := make([]Pos, 0)
	li = append(li, Pos{1, 1})
	for len(li) > 0 {
		p := li[0]
		li = li[1:]
		for _, sh := range shiftMap {
			newPos := Pos{
				p.x + sh.x,
				p.y + sh.y,
			}
			if m[newPos] {
				continue
			}
			m[newPos] = true
			li = append(li, newPos)
		}
	}
	return strconv.Itoa(len(m))
}

func Part1(input Input) string {
	return Part(input)
}

func Part2(input Input) string {
	fixedInput := make(Input, 0, len(input))
	for _, inst := range input {
		dirNum, _ := strconv.Atoi(inst.corol[5:])
		steps, _ := strconv.ParseInt(inst.corol[:5], 16, 0)
		fixedInput = append(fixedInput, Inst{
			Dir(dirNum),
			int(steps),
			"",
		})
	}
	
	return Part(fixedInput)
}

func Part(input Input) string {
	result := 0
	pos := Pos{}
	shiftMap := map[Dir]Pos{
		U: {-1, 0},
		D: {1, 0},
		L: {0, -1},
		R: {0, 1},
	}

	pathSum := 2
	for _, inst := range input {
		nextPos := Pos{
			pos.x + inst.steps*shiftMap[inst.dir].x,
			pos.y + inst.steps*shiftMap[inst.dir].y,
		}
		pathSum += inst.steps
		result += pos.x*nextPos.y - nextPos.x*pos.y
		pos = nextPos
	}

	if result < 0 {
		result = result * -1
	}
	return strconv.Itoa((result + pathSum)/2)
}

func Parse(input string) Input {
	result := make(Input, 0)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		dirStr, stepsStr, colorStr := parts[0], parts[1], parts[2]
		var dir Dir
		if dirStr == "U" {
			dir = U
		}
		if dirStr == "D" {
			dir = D
		}
		if dirStr == "L" {
			dir = L
		}
		if dirStr == "R" {
			dir = R
		}
		steps, _ := strconv.Atoi(stepsStr)
		color := colorStr[2 : len(colorStr)-1]
		result = append(result, Inst{dir, steps, color})
	}
	return result
}
