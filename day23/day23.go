package day23

import (
	"advent-of-code/help"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

const day = 23

type Pos struct {
	x, y int
}

type Graph = map[Pos]map[Pos]int

func Run() {
	input := Parse(help.GetInput(day))
	part1Result := Part1(input)
	part2Result := Part2(input)

	fmt.Println("Part1", part1Result)
	fmt.Println("Part2", part2Result)
}

type Input struct {
	m    []string
	size Pos
}

type PathResult struct {
	pos   Pos
	value int
}

var shiftList = [4]Pos{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
var rr = []uint8{'>', '<', 'v', '^'}

func findNear(input Input, pos Pos, part2 bool) []Pos {
	nextList := make([]Pos, 0)
	for i, sh := range shiftList {
		next := Pos{pos.x + sh.x, pos.y + sh.y}
		r := rr[i]
		if next.x >= 0 && next.y >= 0 && next.x < input.size.x && next.y < input.size.y {
			v := input.m[next.y][next.x]
			if (!part2 && (v == '.' || v == r)) || (part2 && v != '#') {
				nextList = append(nextList, next)
			}
		}
	}
	return nextList
}

func findConnect(input Input, start Pos, part2 bool) []PathResult {
	nextList := findNear(input, start, part2)
	result := make([]PathResult, 0)
out:
	for _, next := range nextList {
		s := start
		for i := 1; ; i++ {

			nn := findNear(input, next, part2)
			si := slices.Index(nn, s)
			if si >= 0 {
				nn = append(nn[:si], nn[si+1:]...)
			}
			if len(nn) == 1 {
				for _, nPos := range nn {
					if nPos != s {
						s = next
						next = nPos
						break
					}
				}
			} else {
				result = append(result, PathResult{
					pos:   next,
					value: i,
				})
				continue out
			}
		}
	}
	return result
}

func Part1(input Input) string {
	start := Pos{1, 0}
	end := Pos{input.size.x - 2, input.size.y - 1}
	posList := []Pos{start}
	graph := make(Graph)
	for len(posList) > 0 {
		pos := posList[0]
		posList = posList[1:]
		if _, ok := graph[pos]; ok {
			continue
		}
		graph[pos] = make(map[Pos]int)
		connections := findConnect(input, pos, false)
		for _, connection := range connections {
			graph[pos][connection.pos] = max(graph[pos][connection.pos], connection.value)
			if connection.pos != end {
				posList = append(posList, connection.pos)
			}
		}
	}

	results := make([]int, 0)
	paths := []PathResult{{
		pos:   start,
		value: 0,
	}}

	for len(paths) > 0 {
		path := paths[0]
		paths = paths[1:]
		g := graph[path.pos]
		for p, steps := range g {
			if p == end {
				results = append(results, path.value+steps)
			} else {
				paths = append(paths, PathResult{
					pos:   p,
					value: path.value + steps,
				})
			}
		}
	}
	slices.Sort(results)

	return strconv.Itoa(results[len(results)-1])
}

func allPaths(prev []Pos, graph Graph, end Pos) [][]Pos {
	last := prev[len(prev)-1]
	result := make([][]Pos, 0)
	for pos, _ := range graph[last] {
		if !slices.Contains(prev, pos) {
			nextSlice := append(prev, pos)
			if pos == end {
				result = append(result, nextSlice)
			} else {
				nextSlices := allPaths(nextSlice, graph, end)
				for _, slice := range nextSlices {
					if slice[len(slice)-1] == end {
						result = append(result, slice)
					}
				}
			}
		} else {
			if prev[len(prev)-1] == end {
				result = append(result, prev)
			}
		}
	}
	return result
}

func Part2(input Input) string {
	start := Pos{1, 0}
	end := Pos{input.size.x - 2, input.size.y - 1}
	posList := []Pos{start}
	graph := make(Graph)
	for len(posList) > 0 {
		pos := posList[0]
		posList = posList[1:]
		if _, ok := graph[pos]; ok {
			continue
		}
		graph[pos] = make(map[Pos]int)
		connections := findConnect(input, pos, true)
		for _, connection := range connections {
			graph[pos][connection.pos] = max(graph[pos][connection.pos], connection.value)
			if connection.pos != end {
				posList = append(posList, connection.pos)
			}
		}
	}

	results := make([]int, 0)
	paths := allPaths([]Pos{start}, graph, end)
	slices.Sort(results)

	result := 0
	for _, p := range paths {

		s := 0
		for i := 0; i < len(p)-1; i++ {
			a, b := p[i], p[i+1]
			s += graph[a][b]
		}

		result = max(result, s)
	}

	return strconv.Itoa(result)
}

func Parse(input string) Input {
	lines := strings.Split(input, "\n")
	size := Pos{len(lines[0]), len(lines)}

	return Input{lines, size}
}
