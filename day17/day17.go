package day17

import (
	"advent-of-code/help"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

const day = 17

type Dir int

const (
	UP Dir = iota
	DOWN
	LEFT
	RIGHT
)

type Step struct {
	pos   Pos
	dir   Dir
	steps int
	heat  int
}

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
type Input struct {
	m    [][]int
	size Pos
}

type Dirs struct {
	x   int
	y   int
	dir Dir
	op  Dir
}

type Visit struct {
	pos   Pos
	dir   Dir
	steps int
}

func (s Step) ToVisit() Visit {
	return Visit{
		s.pos,
		s.dir,
		s.steps,
	}
}

func Part(input Input, minSteps, maxSteps int) string {
	result := 0
	visited := make(map[Visit]bool)

	steps := []Step{{dir: RIGHT}, {dir: DOWN}}
	visited[steps[0].ToVisit()] = true
	visited[steps[1].ToVisit()] = true

	for len(steps) > 0 {
		slices.SortFunc(steps, func(a, b Step) int {
			return a.heat - b.heat
		})
		step := steps[0]
		if step.pos.x == input.size.x-1 && step.pos.y == input.size.y-1 && step.steps >= minSteps {
			return strconv.Itoa(step.heat)
		}
		steps = steps[1:]

		directions := [4]Dirs{
			{1, 0, RIGHT, LEFT},
			{-1, 0, LEFT, RIGHT},
			{0, 1, DOWN, UP},
			{0, -1, UP, DOWN},
		}
		for _, dirs := range directions {
			dir := dirs.dir
			nextPos := Pos{x: step.pos.x + dirs.x, y: step.pos.y + dirs.y}
			if nextPos.x < input.size.x &&
				nextPos.y < input.size.y &&
				nextPos.x >= 0 &&
				nextPos.y >= 0 &&
				(step.dir != dir || step.steps < maxSteps) &&
				(step.dir == dir || step.steps >= minSteps) &&
				step.dir != dirs.op {
				
				st := 1
				if step.dir == dir {
					st = step.steps + 1
				}
				nextStep := Step{
					pos:   nextPos,
					dir:   dir,
					steps: st,
					heat:  step.heat + input.m[nextPos.y][nextPos.x],
				}
				nextVisit := nextStep.ToVisit()
				if visited[nextVisit] {
					continue
				}
				visited[nextVisit] = true
				steps = append(steps, nextStep)
			}
		}
	}

	return strconv.Itoa(result)
}

func Part1(input Input) string {
	return Part(input, 0, 3)
}

func Part2(input Input) string {
	return Part(input, 4, 10)
}

func Parse(input string) Input {
	lines := strings.Split(input, "\n")
	size := Pos{
		len(lines[0]),
		len(lines),
	}
	result := make([][]int, size.y)
	for y, line := range lines {
		result[y] = make([]int, size.x)
		for x, c := range strings.Split(line, "") {
			n, _ := strconv.Atoi(c)
			result[y][x] = n
		}
	}
	return Input{
		result,
		size,
	}
}
