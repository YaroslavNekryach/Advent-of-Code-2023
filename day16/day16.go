package day16

import (
	"advent-of-code/help"
	"fmt"
	"strconv"
	"strings"
)

const day = 16

func Run() {
	input := Parse(help.GetInput(day))
	part1Result := Part1(input)
	part2Result := Part2(input)

	fmt.Println("Part1", part1Result)
	fmt.Println("Part2", part2Result)
}

type Dir int

type Pos struct {
	x int
	y int
}

type PosDir struct {
	pos Pos
	dir Dir
}

func (pd *PosDir) Next() {
	if pd.dir == RIGHT {
		pd.pos.x += 1
	}
	if pd.dir == LEFT {
		pd.pos.x -= 1
	}
	if pd.dir == UP {
		pd.pos.y -= 1
	}
	if pd.dir == DOWN {
		pd.pos.y += 1
	}
}

const (
	UP Dir = iota
	DOWN
	LEFT
	RIGHT
)



type Input struct {
	m []string
	size Pos
}

func Part(input Input, start PosDir) int {
	posMap := make(map[Pos]map[Dir]bool)
	beamList := []PosDir{ start }
	
	for len(beamList) > 0 {
		beam:= beamList[len(beamList) - 1]
		beamList = beamList[0: len(beamList)-1]
		for beam.pos.x >= 0 && beam.pos.x < input.size.x && beam.pos.y >= 0 && beam.pos.y < input.size.y {
			if posMap[beam.pos] != nil && posMap[beam.pos][beam.dir] {
				break 
			}
			if posMap[beam.pos] == nil {
				posMap[beam.pos] = make(map[Dir]bool)
			}
			posMap[beam.pos][beam.dir] = true
			cell := input.m[beam.pos.y][beam.pos.x]
			if cell == '/' {
				if beam.dir == RIGHT {
					beam.dir = UP
				}else if beam.dir == DOWN {
					beam.dir = LEFT
				} else if beam.dir == LEFT {
					beam.dir = DOWN
				} else if beam.dir == UP {
					beam.dir = RIGHT
				}
			} else if cell == '\\' {
				if beam.dir == RIGHT {
					beam.dir = DOWN
				}else if beam.dir == DOWN {
					beam.dir = RIGHT
				} else if beam.dir == LEFT {
					beam.dir = UP
				} else if beam.dir == UP {
					beam.dir = LEFT
				}
			} else if cell == '|' {
				if beam.dir == RIGHT || beam.dir == LEFT {
					beam.dir = DOWN
					cop := beam
					cop.dir = UP
					beamList = append(beamList, cop)
				}
			} else if cell == '-' {
				if beam.dir == UP || beam.dir == DOWN {
					beam.dir = RIGHT
					cop := beam
					cop.dir = LEFT
					beamList = append(beamList, cop)
				}
			}
			beam.Next()
		}
	}
	return len(posMap)
}

func Part1(input Input) string {
	return strconv.Itoa(Part(input, PosDir{Pos{0, 0}, RIGHT}))
}

func Part2(input Input) string {
	result := 0
	for i := 0; i < input.size.y; i++ {
		m := Part(input, PosDir{Pos{0, i}, RIGHT})
		result = max(result, m)
	}
	for i := 0; i < input.size.y; i++ {
		m := Part(input, PosDir{Pos{input.size.x - 1, i}, LEFT})
		result = max(result, m)
	}
	for i := 0; i < input.size.x; i++ {
		m := Part(input, PosDir{Pos{i, 0}, DOWN})
		result = max(result, m)
	}
	for i := 0; i < input.size.x; i++ {
		m := Part(input, PosDir{Pos{i, input.size.y - 1}, UP})
		result = max(result, m)
	}
	return strconv.Itoa(result)
}

func Parse(input string) Input {
	m := strings.Split(input, "\n")
	size := Pos{
		x: len(m[0]),
		y: len(m),
	}
	return Input{
		m, size,
	}
}
