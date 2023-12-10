package day9

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	expedted := "114"
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

const input = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

