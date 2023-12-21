package day20

import (
	"advent-of-code/help"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

const day = 20

func Run() {
	input := Parse(help.GetInput(day))
//	part1Result := Part1(input)
	part2Result := Part2(input)

//	fmt.Println("Part1", part1Result)
	fmt.Println("Part2", part2Result)
}

type Input struct {
	modules map[string]*Module
	broadcaster []*Module
}

type ModuleType int
type State int

const (
	End ModuleType = iota
	FlipFlop
	Conjunction
	
)

const (
	Lo State = iota
	Hi
)


type Module struct {
	moduleType ModuleType
	name string
	inputs []*Module
	outputs []*Module
	state State
}

func (m *Module) In(signal State) bool {
	if m.moduleType == FlipFlop {
		if signal == Lo {
			if m.state == Lo {
				m.state = Hi
			} else {
				m.state = Lo
			}
			return true
		} else {
			return false
		}
	} else if m.moduleType == Conjunction {
		for _, in := range m.inputs {
			if in.state == Lo {
				m.state = Hi
				return true
			}
		}
		m.state = Lo
		return true
	} else {
		return false
	}
}

func Part1(input Input) string {
	lo, hi := 0, 0
	for i := 0; i < 1000; i++ {
		lo++
		signals := make([]*Module, 0)
		for _, b := range input.broadcaster {
			b.In(Lo)
			lo++
			signals = append(signals, b)
		}

		for len(signals) > 0 {
			signal := signals[0]
			signals = signals[1:]

			for _, out := range signal.outputs {
				if signal.state == Lo {
					lo++
				} else {
					hi++
				}
				if out.In(signal.state) {
					signals = append(signals, out)
				}
			}
		}
	}
	return strconv.Itoa(lo * hi)
}

func Part2(input Input) string {
	watch := input.modules["rx"].inputs[0].inputs
	watchResult := make(map[string]int)
	for i := 1; i < 4096 ; i++ {
		signals := make([]*Module, 0)
		for _, b := range input.broadcaster {
			b.In(Lo)
			signals = append(signals, b)
		}

		for len(signals) > 0 {
			signal := signals[0]
			signals = signals[1:]

			for _, out := range signal.outputs {
				if slices.Contains(watch, out) && signal.state == Lo {
					watchResult[out.name] = i
				}
				if out.In(signal.state) {
					signals = append(signals, out)
				}
			}
		}
	}
	result := 1
	for _, v := range watchResult {
		result *= v
	}
	return strconv.Itoa(result)
}

func Parse(input string) Input {
	modules := make(map[string]*Module)
	outMap := make(map[string][]string)
	for _, line := range strings.Split(input, "\n") {
		lineSp := strings.Split(line, " -> ")
		if lineSp[0] != "broadcaster" {
			moduleType := Conjunction
			if lineSp[0][0] == '%' {
				moduleType = FlipFlop
			}
			name := lineSp[0][1:]
			modules[name] = &Module{
				moduleType: moduleType,
				name:       name,
				inputs: 	make([]*Module, 0),
				outputs: 	make([]*Module, 0),
				state:      Lo,
			}
			outMap[name] = strings.Split(lineSp[1], ", ") 
		} else {
			outMap["broadcaster"] = strings.Split(lineSp[1], ", ") 
		}
	}
	broadcaster := make([]*Module, 0)
	for name, to := range outMap {
		for _, v := range to {
			if _, ok := modules[v]; !ok {
				modules[v] = &Module{name: v}
			}
			modules[v].inputs = append(modules[v].inputs, modules[name])
		}
		if name != "broadcaster" {
			for _, v := range to {
				modules[name].outputs = append(modules[name].outputs, modules[v])
			}
		} else {
			for _, v := range to {
				broadcaster = append(broadcaster, modules[v])
			}
		}
	}
	
	return Input{
		modules ,
		broadcaster,
	}
}
