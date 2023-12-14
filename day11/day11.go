package day11

import (
	"advent-of-code/help"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const day = 11

type Pos struct {
	 x int
	 y int
}

func (p Pos) String() string {
	return fmt.Sprintf("%v_%v", p.x, p.y)
}

type Input struct {
	size Pos
	galaxies []Pos
}

func Run() {
	input := Parse(help.GetInput(day))
	galaxies :=  make([]Pos, len(input.galaxies))
	copy(galaxies, input.galaxies)
	part1Result := Part1(Input{input.size, galaxies})
	part2Result := Part2(input, 1000000)

	fmt.Println("Part1", part1Result)
	fmt.Println("Part2", part2Result)
}

func Part1(input Input) string {
	result := 0
	emptyX, emptyY := getEmptyXY(input)
	
	for i:= 0; i < len(input.galaxies); i++ {
		gal := &input.galaxies[i]
		xCount := findLastCount(emptyX, gal.x)
		yCount := findLastCount(emptyY, gal.y)
		gal.x += xCount
		gal.y += yCount
	}
	for a := 0; a < len(input.galaxies) - 1; a++ {
		for b := a + 1; b < len(input.galaxies); b++ {
			result += int(math.Abs(float64(input.galaxies[a].x - input.galaxies[b].x)) + math.Abs(float64(input.galaxies[a].y - input.galaxies[b].y)))
		}	
	}
	return strconv.Itoa(result)
}

func Part2(input Input, times int) string {
	result := 0
	emptyX, emptyY := getEmptyXY(input)

	for i:= 0; i < len(input.galaxies); i++ {
		gal := &input.galaxies[i]
		xCount := findLastCount(emptyX, gal.x)
		yCount := findLastCount(emptyY, gal.y)
		gal.x += xCount * (times - 1)
		gal.y += yCount * (times - 1)
	}
	for a := 0; a < len(input.galaxies) - 1; a++ {
		for b := a + 1; b < len(input.galaxies); b++ {
			result += int(math.Abs(float64(input.galaxies[a].x - input.galaxies[b].x)) + math.Abs(float64(input.galaxies[a].y - input.galaxies[b].y)))
		}	
	}
	return strconv.Itoa(result)
}

func findLastCount(empty []int, index int) int {
	for i, v := range empty {
		if v > index {
			return i
		}
	}
	return len(empty)
}

func getEmptyXY(input Input) ([]int, []int)  {
	my := make(map[int]bool)
	mx := make(map[int]bool)
	
	for _, gal := range input.galaxies {
		my[gal.y] = true
		mx[gal.x] = true
	}
	resultY := make([]int, 0)
	for i := 0; i < input.size.y; i++ {
		if !my[i] {
			resultY = append(resultY, i)
		}
	}
	resultX := make([]int, 0)
	for i := 0; i < input.size.x; i++ {
		if !mx[i] {
			resultX = append(resultX, i)
		}
	}
	return resultX, resultY 
}
func Parse(input string) Input {
//	fmt.Print(input)
	galaxies := make([]Pos, 0)
	var sizeX, sizeY int
	for y, line := range strings.Split(input, "\n") {
		sizeX = len(line)
		for x, str := range strings.Split(line, "") {
			if str == "#" {
				galaxies = append(galaxies, Pos{x, y})
			}
		}
		sizeY = y + 1
	}

	return Input{
		size: Pos{x: sizeX, y: sizeY},
		galaxies: galaxies,
	}
}
