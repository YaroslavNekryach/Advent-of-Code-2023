package day20

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	expedted := "32000000"
	result := Part1(Parse(input))

	if result != expedted {
		t.Fatalf(`Expected = %s, get %s`, expedted, result)
	}
}

func TestPart12(t *testing.T) {
	expedted := "11687500"
	result := Part1(Parse(input2))

	if result != expedted {
		t.Fatalf(`Expected = %s, get %s`, expedted, result)
	}
}

func TestParse(t *testing.T) {
	result := Parse(input)
	fmt.Printf("%v", result)
}


const input = `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`

const input2 = `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output`