package day13

import (
	"advent-of-code/help"
	"fmt"
	"strconv"
	"strings"
)

const day = 13

type Map struct {
	hor []string
	ver []string
}

type Input  = []Map

func Run() {
	input := Parse(help.GetInput(day))
	part1Result := Part1(input)
	part2Result := Part2(input)

	fmt.Println("Part1", part1Result)
	fmt.Println("Part2", part2Result)
}

func cmp(a, b string) int  {
	result := 0 
	for i := 0; i < len(a); i++ {
		if a[i:i+1] != b[i:i+1] {
			result++
		}
	}
	return result
}

func getMirror(input []string, diff int) int  {
	out: for i := 0; i < len(input); i++ {
		if i + 1 >= len(input) {
			return 0
		}
		c := cmp(input[i], input[i + 1])
		if c <= diff {
			for u, d := i-1, i + 2; u >= 0 && d < len(input); u, d = u - 1, d + 1 {
				c += cmp(input[u], input[d])
				if c > diff {
					continue out				
				}
			}
			if c == diff {
				return i + 1
			}
		}
	}
	return 0
}

func Part(input Input, diff int) string {
	result := 0
	for _, m := range input {
		hor := getMirror(m.hor, diff)
		if (hor > 0) {
			result += 100 * hor
			continue
		}
		ver := getMirror(m.ver, diff)
		result += ver
	}
	return strconv.Itoa(result)
}

func Part1(input Input) string {
	return Part(input, 0)
}

func Part2(input Input) string {
	return Part(input, 1)
}

func Parse(input string) Input {
	result := make(Input, 0)
	for _, m := range strings.Split(input, "\n\n") {
		hor := strings.Split(m, "\n")
		ver := make([]string, 0)
		for i := 0; i < len(hor[0]); i++ {
			l := ""
			for _, h := range hor {
				l += h[i:i+1]
			}
			ver = append(ver, l)
		}
		result = append(result, Map{
			hor,
			ver,
		})
	}
	
	return result
}
