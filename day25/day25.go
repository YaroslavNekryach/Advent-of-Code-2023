package day25

import (
	"advent-of-code/help"
	"fmt"
	"strconv"
	"strings"
)

const day = 25

func Run() {
	input := Parse(help.GetInput(day))
	part1Result := Part1(input)

	fmt.Println("Part1", part1Result)
}

type Input = Graph

type Graph = map[string]map[string]bool

func getPath(graph Graph, exclude [][]string, start string, end string) ([]string, int) {
	visited := map[string]bool{start: true}
	list := [][]string{{start}}
	for len(list) > 0 {
		nextList := make([][]string, 0)
		for _, path := range list {
			from := path[len(path)-1]
		out:
			for to := range graph[from] {
				if visited[to] {
					continue
				}
				for _, ex := range exclude {
					if (from == ex[0] && to == ex[1]) || (from == ex[1] && to == ex[0]) {
						continue out
					}
				}
				if to == end {
					return append(path, to), len(visited)
				}
				visited[to] = true
				newPath := make([]string, 0, len(path)+1)
				newPath = append(newPath, path...)
				newPath = append(newPath, to)
				nextList = append(nextList, newPath)
			}
		}
		list = nextList
	}

	return make([]string, 0), len(visited)
}

func Part1(input Input) string {
	list := make([]string, 0)
	for k := range input {
		list = append(list, k)
	}
	for a := 0; a < len(list); a++ {
		af := list[a]
		for at := range input[af] {
			aPath, _ := getPath(input, [][]string{{af, at}}, af, at)
			for b := 0; b < len(aPath)-1; b++ {
				bf, bt := aPath[b], aPath[b+1]
				bPath, _ := getPath(input, [][]string{{af, at}, {bf, bt}}, bf, bt)
				for c := 0; c < len(bPath)-1; c++ {
					cf, ct := bPath[c], bPath[c+1]
					path, visited := getPath(input, [][]string{{af, at}, {bf, bt}, {cf, ct}}, cf, ct)
					if len(path) == 0 {
						return strconv.Itoa(visited * (len(list) - visited))
					}
				}

			}
		}

	}
	return strconv.Itoa(0)
}

func Parse(input string) Input {
	result := make(Input)
	for _, line := range strings.Split(input, "\n") {
		sp := strings.Split(line, ": ")
		from, toStr := sp[0], sp[1]
		to := strings.Split(toStr, " ")

		if result[from] == nil {
			result[from] = make(map[string]bool)
		}
		for _, t := range to {
			if result[t] == nil {
				result[t] = make(map[string]bool)
			}
			result[from][t] = true
			result[t][from] = true
		}
	}
	return result
}

//
//
//jqt -> rhn
//jqt -> xhk
//jqt -> nvd
//rsh -> frs
//rsh -> pzl
//rsh -> lsr
//xhk -> hfx
//cmg -> qnr
//cmg -> nvd
//cmg -> lhk
//cmg -> bvb
//rhn -> xhk
//rhn -> bvb
//rhn -> hfx
//bvb -> xhk
//bvb -> hfx
//pzl -> lsr
//pzl -> hfx
//pzl -> nvd
//qnr -> nvd
//ntq -> jqt
//ntq -> hfx
//ntq -> bvb
//ntq -> xhk
//nvd -> lhk
//lsr -> lhk
//rzs -> qnr
//rzs -> cmg
//rzs -> lsr
//rzs -> rsh
//frs -> qnr
//frs -> lhk
//frs -> lsr
