package day7

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	expedted := "6440"
	result := Part1(Parse(input))

	if result != expedted {
		t.Fatalf(`Expected = %s, get %s`, expedted, result)
	}
}

func TestPart2(t *testing.T) {
	expedted := "5905"
	result := Part2(Parse(input))

	if result != expedted {
		t.Fatalf(`Expected = %s, get %s`, expedted, result)
	}
}

func TestParse(t *testing.T) {
	result := Parse(input)
	fmt.Printf("%v", result)
}

const input = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`
