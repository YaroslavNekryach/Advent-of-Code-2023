package day10

import (
	"advent-of-code/help"
	"fmt"
	"strconv"
	"strings"
)

const day = 10

type Pos struct {
	 x int
	 y int
}

func (p Pos) String() string {
	return fmt.Sprintf("%v_%v", p.x, p.y)
}

type Input = [][]string

func Run() {
	input := Parse(help.GetInput(day))
	part1Result := Part1(input)
	part2Result := Part2(input)

	fmt.Println("Part1", part1Result)
	fmt.Println("Part2", part2Result)
}

func Part1(input Input) string {
	start := findStart(input)
	p := start
	n, _ := getStartPositions(p, input)
	for i:= 0;;i++ {
		n1, n2 := getPosChange(n, input[n.y][n.x])
		if n2 == p {
			p, n = n, n1
		} else {
			p, n = n, n2
		}
		if n == start {
			return  strconv.Itoa((i / 2) + 1) 
		}
	}
}

func Part2(input Input) string {
	start := findStart(input)
	p := start
	c, _ := getStartPositions(p, input)
	l, r, circleMap := lrCount(start, c, input)
	innerMap := make(map[Pos]bool)
	circleMap[p] = true

	isLeft := l > r
	for i:= 0;;i++ {
		n1, n2 := getPosChange(c, input[c.y][c.x])
		var n Pos
		if n2 == p {
			n = n1
		} else {
			n = n2
		}

		var pIn, cIn Pos
		if !isLeft {
			if p.x - c.x == 1 {
				pIn = Pos{p.x, p.y - 1}
				cIn = Pos{c.x, c.y - 1}
			}
			if p.x - c.x == -1 {
				pIn = Pos{p.x, p.y + 1}
				cIn = Pos{c.x, c.y + 1}
			}
			if p.y - c.y == 1 {
				pIn = Pos{p.x + 1, p.y}
				cIn = Pos{c.x + 1, c.y}
			}
			if p.y - c.y == -1 {
				pIn = Pos{p.x - 1, p.y}
				cIn = Pos{c.x - 1, c.y}
			}
		} else {
			if p.x - c.x == 1 {
				pIn = Pos{p.x, p.y + 1}
				cIn = Pos{c.x, c.y + 1}
			}
			if p.x - c.x == -1 {
				pIn = Pos{p.x, p.y - 1}
				cIn = Pos{c.x, c.y - 1}
			}
			if p.y - c.y == 1 {
				pIn = Pos{p.x - 1, p.y}
				cIn = Pos{c.x - 1, c.y}
			}
			if p.y - c.y == -1 {
				pIn = Pos{p.x + 1, p.y}
				cIn = Pos{c.x + 1, c.y}
			}
		}
		if !circleMap[pIn] {
			innerMap[pIn] = true
		}
		if !circleMap[cIn] {
			innerMap[cIn] = true
		}
		p, c = c, n
		if c == start {
			break
		}
	}

	lis := make([]Pos, 0)
	for key, _ := range innerMap {
		lis = append(lis, key)
	}
	for len(lis) > 0 {
		it := lis[0]
		lis = lis[1:]
		up := Pos{it.x - 1, it.y}
		down := Pos{it.x + 1, it.y}
		left := Pos{it.x, it.y - 1}
		right := Pos{it.x, it.y + 1}

		for _, val := range []Pos{up, down, left, right} {
			if !circleMap[val] && !innerMap[val] {
				innerMap[val] = true
				lis = append(lis, val)
			}
		}
	}
	return strconv.Itoa(len(innerMap))
}

func lrCount(start, next Pos, input Input) (int, int, map[Pos]bool)  {
	l, r := 0, 0
	p, c := start, next
	circleMap := make(map[Pos]bool)
	circleMap[p] = true


	for i:= 0;;i++ {
		n1, n2 := getPosChange(c, input[c.y][c.x])
		var n Pos
		if n2 == p {
			n = n1
		} else {
			n = n2
		}

		rl := rorl(p, c, n)
		if rl == "l" {
			l++
		} else if rl == "r" {
			r++
		}
		circleMap[c] = true
		p, c = c, n
		if c == start {
			return l, r, circleMap
		}
	}
}

func rorl(p, c, n Pos) string  {
	d1 := Pos{c.x - p.x, c.y - p.y}
	d2 := Pos{n.x - p.x, n.y - p.y}
	if p.x != n.x && p.y != n.y {
		if d1.x == 0 {
			if d1.y - d2.x == 0 {
				return "l"
			} else {
				return "r"
			}
		} else {
			if d1.x - d2.y == 0 {
				return "r"
			} else {
				return "l"
			}
		}
	}
	return ""
}

func findStart(input Input) Pos {
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if input[y][x] == "S" {
				return Pos{x ,y}
			}
		}
	}
	return Pos{}
}

func getStartPositions(startPos Pos, input Input) (Pos, Pos) {
	shifts := []Pos{Pos{x: 0, y: 1}, Pos{x: 0, y: -1}, Pos{x: 1, y: 0}, Pos{x: -1, y: 0}}
	res := make([]Pos, 0)
	for _, shift := range shifts {
		next := Pos{x: startPos.x + shift.x, y: startPos.y + shift.y}
		if next.x < 0 || next.y < 0 || next.x >= len(input[0]) || next.y >= len(input) {
			continue
		}
		a,b := getPosChange(next, input[next.y][next.x])
		if a == startPos || b == startPos{
			res = append(res, next)
		}
	}
	return res[0], res[1]
}

func getPosChange(pos Pos, s string) (Pos, Pos)  {
	a := Pos{}
	b := Pos{}
	
	if s == "|" {
		a, b = Pos{x: pos.x, y: pos.y - 1}, Pos{x: pos.x, y: pos.y + 1}
	} else if s == "-" {
		a, b = Pos{x: pos.x - 1, y: pos.y}, Pos{x: pos.x + 1, y: pos.y}
	} else if s == "L" {
		a, b = Pos{x: pos.x + 1, y: pos.y}, Pos{x: pos.x, y: pos.y - 1}
	} else if s == "J" {
		a, b = Pos{x: pos.x - 1, y: pos.y}, Pos{x: pos.x, y: pos.y - 1}
	} else if s == "7" {
		a, b = Pos{x: pos.x - 1, y: pos.y}, Pos{x: pos.x, y: pos.y + 1}
	} else if s == "F" {
		a, b = Pos{x: pos.x + 1, y: pos.y}, Pos{x: pos.x, y: pos.y + 1}
	}
	
//	fmt.Printf("Pos change: from %v to %v, %v on %v\n", pos, a, b, s)
	
	return a, b
}

func Parse(input string) Input {
	result := make(Input, 0)
	for i, line := range strings.Split(input, "\n") {
		result = append(result, make([]string, 0))
		for _, str := range strings.Split(line, "") {
			result[i] = append(result[i], str)
		}
	}

	return result
}
