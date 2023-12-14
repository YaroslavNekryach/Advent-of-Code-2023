package day11

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	expedted := "374"
	result := Part1(Parse(input))

	if result != expedted {
		t.Fatalf(`Expected = %s, get %s`, expedted, result)
	}
}

func TestPart2(t *testing.T) {
	
	expedted := "1030"
	result := Part2(Parse(input), 10)

	if result != expedted {
		t.Fatalf(`Expected = %s, get %s`, expedted, result)
	}
}

func TestParse(t *testing.T) {
	result := Parse(input)
	fmt.Printf("%v", result)
}

const input = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

