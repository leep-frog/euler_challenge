package eulerchallenge

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/leep-frog/command"
)

func Branches() map[string]*command.Node {
	problems := []*problem{
		P1(),
		P2(),
		P3(),
		P4(),
		P5(),
		P6(),
		P7(),
		P8(),
		P9(),
		P10(),
		P11(),
		P12(),
		P13(),
		P14(),
		P15(),
		P16(),
		P17(),
		P18(),
		P19(),
		P20(),
		P21(),
		P22(),
		P23(),
		P24(),
		P25(),
		P26(),
		P27(),
		P28(),
		P29(),
		P30(),
		P31(),
		P32(),
		P33(),
		P34(),
		P35(),
		P36(),
		P37(),
		P38(),
		P39(),
		P40(),
		P41(),
		P42(),
		P43(),
		P44(),
		P45(),
		P46(),
		P47(),
		P48(),
		P49(),
		P50(),
		P51(),
		P52(),
		P53(),
		P54(),
		P55(),
		P56(),
		P57(),
		P58(),
		P59(),
		P60(),
		P61(),
		P62(),
		P63(),
		P64(),
		P65(),
		P66(),
		// 67 is a bigger version of problem 18
		// 68 was solved in python
		P69(),
		P70(),
		P71(),
		P72(),
		P73(),
		P74(),
		P75(),
		P76(),
		P77(),
		P78(),
		P79(),
		P80(),
		P81(),
		P82(),
		P83(),
		P84(),
		P85(),
		P86(),
		P87(),
		P88(),
		P89(),
		P90(),
		P91(),
		P92(),
		P93(),
		P94(),
		P95(),
		P96(),
		P97(),
		P98(),
		P99(),
		P100(),
		P101(),
		P102(),
		P103(),
		P104(),
		P105(),
		P106(),
		P107(),
		P108(),
		P109(),
		// P110 is P108 with a different input
		P111(),
		P112(),
		P113(),
		P114(),
		P115(),
		P116(),
		P117(),
		P118(),
		P119(),
		P120(),
		P121(),
		P122(),
		P123(),
		P124(),
		P125(),
		P126(),
		P127(),
		P128(),
		P129(),
		P130(),
		P131(),
		P132(),
		P133(),
		P134(),
		P135(),
		P136(),
		P137(),
		P138(),
		P139(),
		// END_LIST (needed for file_generator.go)
	}

	m := map[string]*command.Node{
		"fg": FileGenerator(),
	}
	for i, p := range problems {
		pStr := fmt.Sprintf("%d", p.num)
		if _, ok := m[pStr]; ok {
			log.Fatalf("Duplicate problem entry: %d, %d", i, p.num)
		}
		m[pStr] = p.n
	}
	return m
}

func descNode(problem int) command.Processor {
	return command.Descriptionf("https://projecteuler.net/problem=%d", problem)
}

func intInputNode(num int, f func(command.Output, int)) *problem {
	return &problem{
		num: num,
		n: command.SerialNodes(
			descNode(num),
			command.Arg[int](N, "", command.Positive[int]()),
			command.ExecutorNode(func(o command.Output, d *command.Data) {
				f(o, d.Int(N))
			}),
		),
	}
}

type problem struct {
	num int
	n   *command.Node
}

func fileInputNode(num int, f func([]string, command.Output)) *problem {
	return &problem{
		num: num,
		n: command.SerialNodes(
			descNode(num),
			command.FileContents("FILE", "", command.NewTransformer[string](func(s string) (string, error) {
				return filepath.Join("input", s), nil
			}, false)),
			command.ExecutorNode(func(o command.Output, d *command.Data) {
				f(d.StringList("FILE"), o)
			}),
		),
	}
}

func noInputNode(num int, f func(command.Output)) *problem {
	return &problem{
		num: num,
		n: command.SerialNodes(
			descNode(num),
			command.ExecutorNode(func(o command.Output, d *command.Data) {
				f(o)
			}),
		),
	}
}
