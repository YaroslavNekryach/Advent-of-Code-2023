package day10

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	expedted := "4"
	result := Part1(Parse(input))

	if result != expedted {
		t.Fatalf(`Expected = %s, get %s`, expedted, result)
	}
}

func TestPart2(t *testing.T) {
	
	expedted := "2"
	result := Part2(Parse(input))

	if result != expedted {
		t.Fatalf(`Expected = %s, get %s`, expedted, result)
	}
}

func TestParse(t *testing.T) {
	result := Parse(input)
	fmt.Printf("%v", result)
}

const input = `-L|F7
7S-7|
L|7||
-L-J|
L|-JF`

