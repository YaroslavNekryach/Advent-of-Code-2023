package day19

import (
	"advent-of-code/help"
	"fmt"
	"strconv"
	"strings"
)

const day = 19

func Run() {
	input := Parse(help.GetInput(day))
	part1Result := Part1(input)
	part2Result := Part2(input)

	fmt.Println("Part1", part1Result)
	fmt.Println("Part2", part2Result)
}

type Part = map[string]int

type Condition struct {
	category string
	gt       bool
	value    int
	result   string
}
type Worckflow struct {
	name       string
	conditions []Condition
	othrwise   string
}
type Input struct {
	worckflows map[string]Worckflow
	parts      []Part
}

type Range struct {
	from int
	to int
}

type Ranges = map[string]Range

func SplitRanges(rs *Ranges, c Condition) (*Ranges, *Ranges) {
	r := (*rs)[c.category]
	t, f := r.Split(c.value, c.gt)
	var tr, fr *Ranges
	if t != nil {
		cp := make(Ranges)
		for k, v := range *rs {
			if k == c.category {
				cp[k] = *t
			} else {
				cp[k] = v
			}
		}
		tr = &cp
	}
	if f != nil {
		cp := make(Ranges)
		for k, v := range *rs {
			if k == c.category {
				cp[k] = *f
			} else {
				cp[k] = v
			}
		}
		fr = &cp
	}
	return tr, fr
}

func (r *Range) Split(value int, gt bool) (*Range, *Range)  {
	if gt {
		if value < r.from {
			return r, nil
		}
		if r.from <= value && value < r.to {
			return &Range{value + 1, r.to}, &Range{r.from, value} 
		}
		if value <= r.to {
			return nil, r
		}
	} else {
		if value <= r.from {
			return  nil, r
		}
		if r.from < value && value <= r.to {
			return &Range{r.from, value - 1}, &Range{value, r.to}
		}
		if value < r.to {
			return r, nil
		}
	}
	return nil, nil
}

func Part1(input Input) string {
	result := 0
	acepted := make([]Part, 0)
	for _, part := range input.parts {
		wf := input.worckflows["in"]
	out:
		for {
			for _, cond := range wf.conditions {
				if (cond.gt && part[cond.category] > cond.value) || (!cond.gt && part[cond.category] < cond.value) {
					if cond.result == "A" {
						acepted = append(acepted, part)
						break out
					} else if cond.result == "R" {
						break out
					} else {
						wf = input.worckflows[cond.result]
						continue out
					}
				}
			}
			if wf.othrwise == "A" {
				acepted = append(acepted, part)
				break out
			} else if wf.othrwise == "R" {
				break out
			} else {
				wf = input.worckflows[wf.othrwise]
			}
		}
	}
	
	for _, a := range acepted {
		for _, v := range a {
			result += v
		}
	}
	return strconv.Itoa(result)
}

func Part2(input Input) string {
	result := 0
	r := Ranges{
		"x": {1, 4000},
		"m": {1, 4000},
		"a": {1, 4000},
		"s": {1, 4000},
	}
	acc := getAccepted(&r, "in", &input.worckflows)
	
	for _, rr := range acc {
		sub := 1
		for _, v := range rr {
			sub *= v.to - v.from + 1
		}
		result += sub
	}
	return strconv.Itoa(result)
}

func getAccepted(rs *Ranges, start string, wf *map[string]Worckflow) []Ranges {
	if start == "A" {
		return []Ranges{*rs}
	}
	if start == "R" {
		return []Ranges{}
	}
	w := (*wf)[start]
	accepted := make([]Ranges, 0)
	for _, cond := range w.conditions {
		t, f := SplitRanges(rs, cond)
		accepted = append(accepted, getAccepted(t, cond.result, wf)...)
		rs = f
	}
	accepted = append(accepted, getAccepted(rs, w.othrwise, wf)...)
	return accepted
}

func Parse(input string) Input {
	inputSplit := strings.Split(input, "\n\n")
	worckflowsStr, partsStr := inputSplit[0], inputSplit[1]

	worckflows := make(map[string]Worckflow)
	for _, wStr := range strings.Split(worckflowsStr, "\n") {
		s1 := strings.Split(wStr, "{")
		name, other := s1[0], s1[1][:len(s1[1])-1]
		condsString := strings.Split(other, ",")
		othrwise := condsString[len(condsString)-1]
		condsString = condsString[:len(condsString)-1]
		conditions := make([]Condition, 0)
		for _, condS := range condsString {
			part := condS[0:1]
			cond := condS[1:2]
			last := condS[2:]
			sp := strings.Split(last, ":")
			valS, result := sp[0], sp[1]
			val, _ := strconv.Atoi(valS)
			gt := cond == ">"
			conditions = append(conditions, Condition{
				category: part,
				gt:       gt,
				value:    val,
				result:   result,
			})
		}
		worckflows[name] = Worckflow{
			name:       name,
			conditions: conditions,
			othrwise:   othrwise,
		}
	}
	parts := make([]Part, 0)
	for _, partS := range strings.Split(partsStr, "\n") {
		partS = partS[1 : len(partS)-1]
		part := make(Part, 0)
		for _, p := range strings.Split(partS, ",") {
			ps := strings.Split(p, "=")
			name, vs := ps[0], ps[1]
			value, _ := strconv.Atoi(vs)
			part[name] = value
		}
		parts = append(parts, part)
	}

	return Input{
		worckflows: worckflows,
		parts:      parts,
	}
}
