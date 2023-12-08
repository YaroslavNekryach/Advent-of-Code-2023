package day8

import (
	"advent-of-code/help"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const day = 8

type Node struct {
	name  string
	left  string
	right string
}

type Input struct {
	nodes map[string]Node
	path  []string
}

func Run() {
	input := Parse(help.GetInput(day))
	part1Result := Part1(input)
	part2Result := Part2(input)

	fmt.Println("Part1", part1Result)
	fmt.Println("Part2", part2Result)
}

func Part1(input Input) string {
	i := 0
	node := "AAA"
	for {
		step := input.path[i%len(input.path)]
		if step == "L" {
			node = input.nodes[node].left
		} else {
			node = input.nodes[node].right
		}
		i++
		if node == "ZZZ" {
			break
		}
	}
	return strconv.Itoa(i)
}

// function to calculate gcd (Greatest Common Divisor)
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// function to calculate lcm (Least Common Multiple)
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func Part2(input Input) string {
	result := 1
	nodes := make([]string, 0)
	for k, _ := range input.nodes {
		if k[2:] == "A" {
			nodes = append(nodes, k)
		}
	}
	resultList := make([]int, 0)
	
	for _, node := range nodes {
		i := 0
		for {
			step := input.path[i%len(input.path)]
			if step == "L" {
				node = input.nodes[node].left
			} else {
				node = input.nodes[node].right
			}
			i++
			if node[2:] == "Z" {
				break
			}
		}
		resultList = append(resultList, i)
		
	}
	for _, r := range resultList {
		result = lcm(r, result)
	}
	
	return strconv.Itoa(result)
}

func Parse(input string) Input {
	split := strings.Split(input, "\n\n")
	pathStr, nodesStr := split[0], split[1]
	reg := regexp.MustCompile(`(\S+) = \((\S+), (\S+)\)`)
	nodes := make(map[string]Node)
	for _, line := range strings.Split(nodesStr, "\n") {
		match := reg.FindAllStringSubmatch(line, 1)
		nodes[match[0][1]] = Node{
			name:  match[0][1],
			left:  match[0][2],
			right: match[0][3],
		}
	}

	return Input{
		nodes: nodes,
		path:  strings.Split(pathStr, ""),
	}
}
