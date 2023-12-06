package day5

import (
	"advent-of-code/help"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const day = 5

type ConvertRange struct {
	dest   int
	source int
	rng    int
}

func (rng ConvertRange) Throuthg(value int) (int, bool) {
	if value >= rng.source && value < rng.source+rng.rng {
		return rng.dest + value - rng.source, true
	}
	return value, false
}

func (conRng ConvertRange) RangeThrouthg(rng Range) ([]Range, []Range) {
	ranges := make([]Range, 0)
	conRngTo := conRng.source + conRng.rng - 1
	if rng.from < conRng.source && rng.to >= conRng.source {
		ranges = append(ranges, Range{
			from: rng.from,
			to:   conRng.source - 1,
		})
		rng.from = conRng.source
	}
	if rng.from <= conRngTo && rng.to > conRngTo {
		ranges = append(ranges, Range{
			from: rng.from,
			to:   conRngTo,
		})
		rng.from = conRngTo + 1
	}

	diff := conRng.dest - conRng.source

	ranges = append(ranges, rng)
	modified := make([]Range, 0)
	unmodified := make([]Range, 0)

	for _, r := range ranges {
		if r.from >= conRng.source && r.to <= conRngTo {
			modified = append(modified, Range{
				from: r.from + diff,
				to:   r.to + diff,
			})
		} else {
			unmodified = append(unmodified, r)
		}
	}
	return modified, unmodified
}

type Range struct {
	from int
	to   int
}

type Section struct {
	name   string
	ranges []ConvertRange
}

type Input struct {
	seeds    []int
	sections []Section
}

func Run() {
	input := Parse(help.GetInput(day))
	part1Result := Part1(input)
	part2Result := Part2(input)

	fmt.Println("Part1", part1Result)
	fmt.Println("Part2", part2Result)
}

func Part1(input Input) string {
	result := math.MaxInt
	for _, seed := range input.seeds {
		for _, section := range input.sections {
			for _, rng := range section.ranges {
				var changed bool
				seed, changed = rng.Throuthg(seed)
				if changed {
					break
				}
			}
		}
		if seed < result {
			result = seed
		}
	}
	return strconv.Itoa(result)
}

func Part2(input Input) string {
	result := math.MaxInt
	seedRanges := make([]Range, 0)
	for i := 0; i < len(input.seeds); i += 2 {
		seedRanges = append(seedRanges, Range{
			from: input.seeds[i],
			to:   input.seeds[i] + input.seeds[i+1],
		})
	}

	for _, section := range input.sections {
		modifiedRanges := make([]Range, 0)

		for _, conRng := range section.ranges {
			unmodifiedRanges := make([]Range, 0)
			for _, rng := range seedRanges {
				modified, unmodified := conRng.RangeThrouthg(rng)
				unmodifiedRanges = append(unmodifiedRanges, unmodified...)
				modifiedRanges = append(modifiedRanges, modified...)
			}
			seedRanges = unmodifiedRanges
		}
		seedRanges = append(seedRanges, modifiedRanges...)
	}

	for _, r := range seedRanges {
		if r.from < result {
			result = r.from
		}
	}

	return strconv.Itoa(result)
}

func Parse(input string) Input {
	sectionsString := strings.Split(input, "\n\n")
	seedsString := sectionsString[0]
	seedsSplit := strings.Split(seedsString, ": ")
	seeds := make([]int, 0)
	for _, v := range strings.Split(seedsSplit[1], " ") {
		seed, _ := strconv.Atoi(v)
		seeds = append(seeds, seed)
	}
	sectionsString = sectionsString[1:]

	sections := make([]Section, 0)
	for _, sectionString := range sectionsString {
		sectionLines := strings.Split(sectionString, "\n")
		firstLineSplit := strings.Split(sectionLines[0], " ")
		name := firstLineSplit[0]

		ranges := make([]ConvertRange, 0)
		for _, line := range sectionLines[1:] {
			vals := strings.Split(line, " ")
			dest, _ := strconv.Atoi(vals[0])
			source, _ := strconv.Atoi(vals[1])
			rng, _ := strconv.Atoi(vals[2])
			rngOb := ConvertRange{
				dest, source, rng,
			}
			ranges = append(ranges, rngOb)
		}
		sections = append(sections, Section{
			name, ranges,
		})

	}

	return Input{
		seeds, sections,
	}
}
