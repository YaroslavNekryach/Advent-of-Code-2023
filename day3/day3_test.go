package day3

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	expedted := "4361"

	result := Part1(Parse(input))

	if result != expedted {
		t.Fatalf(`Expected = %s, get %s`, expedted, result)
	}
}

func TestPart2(t *testing.T) {
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	expedted := "467835"

	result := Part2(Parse(input))

	if result != expedted {
		t.Fatalf(`Expected = %s, get %s`, expedted, result)
	}
}

func TestParse(t *testing.T) {
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	result := Parse(input)
	fmt.Printf("%v", result)
}
