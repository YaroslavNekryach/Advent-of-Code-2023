package day22

import (
	"advent-of-code/help"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

const day = 22

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
	z int
}

type Brick = []*Pos

type Input = []Brick

func Down(brick Brick) {
	for i := 0; i < len(brick); i++ {
		brick[i].z--
	}
}

func getUnderAbove(input Input) (map[int]map[int]bool, map[int]map[int]bool) {
	slices.SortFunc(input, func(a, b Brick) int {
		za := math.MaxInt
		zb := math.MaxInt
		for _, aa := range a {
			za = min(za, aa.z)
		}
		for _, bb := range b {
			zb = min(zb, bb.z)
		}
		return za - zb
	})
	m := make(map[Pos]int)
	for i, brick := range input {
		for {
			stop := false
			for _, c := range brick {
				if _, ok := m[Pos{c.x, c.y, c.z - 1}]; c.z == 0 || ok {
					stop = true
				}
			}
			if stop {
				for _, c := range brick {
					m[*c] = i
				}
				break
			} else {
				Down(brick)
			}
		}
	}
	under := make(map[int]map[int]bool)
	above := make(map[int]map[int]bool)

	for pos, brick := range m {
		if _, ok := above[brick]; !ok {
			above[brick] = make(map[int]bool)
		}
		if v, ok := m[Pos{pos.x, pos.y, pos.z + 1}]; ok && v != brick {
			above[brick][v] = true

			if _, ok := under[v]; !ok {
				under[v] = make(map[int]bool)
			}
			under[v][brick] = true
		}
	}

	return under, above
}

func Part1(input Input) string {
	result := 0
	under, above := getUnderAbove(input)
out:
	for _, ab := range above {
		for u := range ab {
			if len(under[u]) < 2 {
				continue out
			}
		}
		result++

	}

	return strconv.Itoa(result)
}

func getFallCount(under, above map[int]map[int]bool, brick int) int {
	fallMap := make(map[int]bool)
	nextList := make([]int, 0)
	for k := range above[brick] {
		nextList = append(nextList, k)
	}
	fallMap[brick] = true
out:
	for len(nextList) > 0 {
		slices.Sort(nextList)
		b := nextList[0]
		nextList = nextList[1:]
		for u := range under[b] {
			if !fallMap[u] {
				continue out
			}
		}
		for k := range above[b] {
			nextList = append(nextList, k)
		}
		fallMap[b] = true
	}
	return len(fallMap) - 1
}

func Part2(input Input) string {
	result := 0
	under, above := getUnderAbove(input)
	for i := 0; i < len(input); i++ {
		result += getFallCount(under, above, i)
	}

	return strconv.Itoa(result)
}

func Parse(input string) Input {
	lines := strings.Split(input, "\n")
	result := make(Input, 0)
	for _, line := range lines {
		brick := make(Brick, 0)
		sp := strings.Split(line, "~")
		sp1 := strings.Split(sp[0], ",")
		sp2 := strings.Split(sp[1], ",")
		x1, _ := strconv.Atoi(sp1[0])
		y1, _ := strconv.Atoi(sp1[1])
		z1, _ := strconv.Atoi(sp1[2])
		x2, _ := strconv.Atoi(sp2[0])
		y2, _ := strconv.Atoi(sp2[1])
		z2, _ := strconv.Atoi(sp2[2])

		if x1 != x2 {
			f, t := min(x1, x2), max(x1, x2)
			for i := f; i <= t; i++ {
				brick = append(brick, &Pos{i, y1, z1})
			}
		} else if y1 != y2 {
			f, t := min(y1, y2), max(y1, y2)
			for i := f; i <= t; i++ {
				brick = append(brick, &Pos{x1, i, z1})
			}
		} else if z1 != z2 {
			f, t := min(z1, z2), max(z1, z2)
			for i := f; i <= t; i++ {
				brick = append(brick, &Pos{x1, y1, i})
			}
		} else {
			brick = append(brick, &Pos{x1, y1, z1})
		}
		result = append(result, brick)
	}
	return result
}
