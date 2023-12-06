package day6

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	expedted := "288"
	result := Part1(Parse(input))

	if result != expedted {
		t.Fatalf(`Expected = %s, get %s`, expedted, result)
	}
}

func TestPart2(t *testing.T) {
	expedted := "71503"
	result := Part2(Parse(input))

	if result != expedted {
		t.Fatalf(`Expected = %s, get %s`, expedted, result)
	}
}

func TestParse(t *testing.T) {
	result := Parse(input)
	fmt.Printf("%v", result)
}

const input = `Time:      7  15   30
Distance:  9  40  200`
