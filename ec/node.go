package eulerchallenge

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/leep-frog/command"
)

func getProblems() []*problem {
	return []*problem{
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
		// 68 was solved in python (TODO: in go?)
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
		P140(),
		P141(),
		P142(),
		P143(),
		P144(),
		P145(),
		P146(),
		P147(),
		P148(),
		P149(),
		P150(),
		P151(),
		P152(),
		P153(),
		P155(),
		P184(),
		P222(),
		P234(),
		P235(),
		P236(),
		P252(),
		P333(),
		P456(),
		P154(),
		P156(),
		P157(),
		P158(),
		P159(),
		P160(),
		P161(),
		P162(),
		P165(),
		P164(),
		P163(),
		P169(),
		P243(),
		P233(),
		P166(),
		P167(),
		P168(),
		P170(),
		P171(),
		P173(),
		P172(),
		// END_LIST (needed for file_generator.go)
	}
}

func Branches() map[string]*command.Node {
	m := map[string]*command.Node{}
	for i, p := range getProblems() {
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

func intInputNode(num int, f func(command.Output, int), executions []*execution) *problem {
	return &problem{
		num: num,
		n: command.SerialNodes(
			descNode(num),
			command.Arg[int](N, "", command.Positive[int]()),
			&command.ExecutorProcessor{F: func(o command.Output, d *command.Data) error {
				f(o, d.Int(N))
				return nil
			}},
		),
		executions: executions,
	}
}

func intsInputNode(num, size int, f func(command.Output, []int), executions []*execution) *problem {
	return &problem{
		num: num,
		n: command.SerialNodes(
			descNode(num),
			command.ListArg[int](N, "", size, 0),
			&command.ExecutorProcessor{F: func(o command.Output, d *command.Data) error {
				f(o, d.IntList(N))
				return nil
			}},
		),
		executions: executions,
	}
}

type problem struct {
	num        int
	n          *command.Node
	executions []*execution
}

type execution struct {
	args     []string
	want     string
	estimate float64
	skip     string
}

func fileInputNode(num int, f func([]string, command.Output), executions []*execution) *problem {
	_, dir, _, ok := runtime.Caller(3)
	if !ok {
		panic("failed to fetch file caller")
	}
	dir = filepath.Dir(dir)

	return &problem{
		num: num,
		n: command.SerialNodes(
			descNode(num),
			// TODO: RelativeFileNode
			command.Arg[string]("FILE", "", &command.FileCompleter[string]{
				Directory: filepath.Join(dir, "input"),
			}),
			&command.ExecutorProcessor{F: func(o command.Output, d *command.Data) error {
				b, err := os.ReadFile(filepath.Join(dir, "input", d.String("FILE")))
				if err != nil {
					return o.Annotatef(err, "failed to read fileee")
				}
				f(strings.Split(strings.TrimSpace(string(b)), "\n"), o)
				return nil
			}},
		),
		executions: executions,
	}
}

func noInputNode(num int, f func(command.Output), ex *execution) *problem {
	return &problem{
		num: num,
		n: command.SerialNodes(
			descNode(num),
			&command.ExecutorProcessor{F: func(o command.Output, d *command.Data) error {
				f(o)
				return nil
			}},
		),
		executions: []*execution{ex},
	}
}
