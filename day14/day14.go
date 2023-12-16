package day14

import (
	"advent-of-code/help"
	"fmt"
	"strconv"
	"strings"
)

const day = 14

func Run() {
	input := Parse(help.GetInput(day))
	part1Result := Part1(input)
	part2Result := Part2(input)

	fmt.Println("Part1", part1Result)
	fmt.Println("Part2", part2Result)
}

type RockType int

const (
	EMPTY RockType  = iota
	ROUND
	CUBE
)
type Pos struct {
	x int
	y int
}

func (p Pos) String() string {
	return fmt.Sprintf("%v_%v", p.x, p.y)
}

type Input  struct {
	size Pos
	rocks map[Pos]RockType
}
func Part1(input Input) string {
	result := 0
	for x := 0; x < input.size.x; x++ {
		count := 0
		for y := input.size.y - 1; y >=0 ; y-- {
			pos := Pos{x, y}
			if input.rocks[pos] == ROUND {
				count++
				delete(input.rocks, pos)
			} else if input.rocks[pos] == CUBE {
				for i := 0; i < count; i++ {
					input.rocks[Pos{x, y + 1 + i}] = ROUND
				}
				count = 0
			}
		}
		for i := 0; i < count; i++ {
			input.rocks[Pos{x, i}] = ROUND
		}
	}
	
	for pos, rockType := range input.rocks {
		if rockType == ROUND {
			result += input.size.y - pos.y
		}
	}
	return strconv.Itoa(result)
}

func Part2(input Input) string {
	result := 0
	steps := 1000000000
	cache := make(map[string]int)
	for j := 0; j < steps; j++ {
		str := fmt.Sprintf("%v", input.rocks)
		cache[str] = j
		for x := 0; x < input.size.x; x++ {
			count := 0
			for y := input.size.y - 1; y >=0 ; y-- {
				pos := Pos{x, y}
				if input.rocks[pos] == ROUND {
					count++
					delete(input.rocks, pos)
				} else if input.rocks[pos] == CUBE {
					for i := 0; i < count; i++ {
						input.rocks[Pos{x, y + 1 + i}] = ROUND
					}
					count = 0
				}
			}
			for i := 0; i < count; i++ {
				input.rocks[Pos{x, i}] = ROUND
			}
		}

		for y := 0; y < input.size.y; y++ {
			count := 0
			for x := input.size.x - 1; x >=0 ; x-- {
				pos := Pos{x, y}
				if input.rocks[pos] == ROUND {
					count++
					delete(input.rocks, pos)
				} else if input.rocks[pos] == CUBE {
					for i := 0; i < count; i++ {
						input.rocks[Pos{x + 1 + i, y}] = ROUND
					}
					count = 0
				}
			}
			for i := 0; i < count; i++ {
				input.rocks[Pos{i, y}] = ROUND
			}
		}

		for x := 0; x < input.size.x; x++ {
			count := 0
			for y := 0; y < input.size.y ; y++ {
				pos := Pos{x, y}
				if input.rocks[pos] == ROUND {
					count++
					delete(input.rocks, pos)
				} else if input.rocks[pos] == CUBE {
					for i := 0; i < count; i++ {
						input.rocks[Pos{x, y - 1 - i}] = ROUND
					}
					count = 0
				}
			}
			for i := 0; i < count; i++ {
				input.rocks[Pos{x, input.size.y - 1 - i}] = ROUND
			}
		}

		for y := 0; y < input.size.y; y++ {
			count := 0
			for x := 0; x < input.size.x ; x++ {
				pos := Pos{x, y}
				if input.rocks[pos] == ROUND {
					count++
					delete(input.rocks, pos)
				} else if input.rocks[pos] == CUBE {
					for i := 0; i < count; i++ {
						input.rocks[Pos{x - 1 - i, y}] = ROUND
					}
					count = 0
				}
			}
			for i := 0; i < count; i++ {
				input.rocks[Pos{input.size.x - 1 - i, y}] = ROUND
			}
		}
		str = fmt.Sprintf("%v", input.rocks)
		if cache[str] > 0 {
			diff := j + 1 - cache[str]
			j += diff * ((steps - j) / diff)
		}
		
	}
	
	for pos, rockType := range input.rocks {
		if rockType == ROUND {
			result += input.size.y - pos.y
		}
	}
	return strconv.Itoa(result)
}

func Parse(input string) Input {
	rocks := make(map[Pos]RockType)
	size := Pos{}
	for y, line := range strings.Split(input, "\n") {
		size.y = y + 1
		size.x = len(line)
		for x, sym := range strings.Split(line, "") {
			
			if sym == "O" {
				rocks[Pos{x, y}] = ROUND
			} else if (sym == "#") {
				rocks[Pos{x, y}] = CUBE
			}
		}
	}
	
	return Input{
		size,
		rocks,
	}
}
