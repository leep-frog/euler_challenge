package ecmodels

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/command/commander"
)

const (
	N       = "N"
	Letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func DescNode(problem int) command.Processor {
	return commander.Descriptionf("https://projecteuler.net/problem=%d", problem)
}

func IntInputNode(num int, f func(command.Output, int), executions []*Execution) *Problem {
	return &Problem{
		Num: num,
		N: commander.SerialNodes(
			DescNode(num),
			commander.Arg[int](N, "", commander.Positive[int]()),
			&commander.ExecutorProcessor{F: func(o command.Output, d *command.Data) error {
				f(o, d.Int(N))
				return nil
			}},
		),
		Executions: executions,
	}
}

func IntsInputNode(num, numInputs, numOptionalInputs int, f func(command.Output, []int), executions []*Execution) *Problem {
	return &Problem{
		Num: num,
		N: commander.SerialNodes(
			DescNode(num),
			commander.ListArg[int](N, "", numInputs, numOptionalInputs),
			&commander.ExecutorProcessor{F: func(o command.Output, d *command.Data) error {
				f(o, d.IntList(N))
				return nil
			}},
		),
		Executions: executions,
	}
}

type Problem struct {
	Num        int
	N          command.Node
	Executions []*Execution
}

func (p *Problem) Node() command.Node {
	return p.N
}

type Execution struct {
	Args     []string
	Want     string
	Estimate float64
	Skip     string
}

var (
	exampleFlag = commander.BoolFlag("example", 'x', "Whether to use the example input or not")
)

func FileInputNode(num int, f func([]string, command.Output), executions []*Execution) *Problem {
	_, ecModelsGo, _, ok := runtime.Caller(0)
	if !ok {
		panic("failed to fetch file caller")
	}
	inputDir := filepath.Join(filepath.Dir(filepath.Dir(ecModelsGo)), "input")

	inputFileArg := commander.OptionalArg[string]("INPUT_FILE", "The input file for the problem", &commander.FileCompleter[string]{
		Directory:         inputDir,
		IgnoreDirectories: true,
	})

	return &Problem{
		Num: num,
		N: commander.SerialNodes(
			DescNode(num),
			commander.FlagProcessor(
				exampleFlag,
			),
			inputFileArg,
			&commander.ExecutorProcessor{F: func(o command.Output, d *command.Data) error {

				// Determine which input file to use
				var inputFile string
				if inputFileArg.Provided(d) {
					inputFile = filepath.Join(inputDir, inputFileArg.Get(d))
				} else if exampleFlag.Get(d) {
					inputFile = filepath.Join(inputDir, fmt.Sprintf("p%d_example.txt", num))
				} else {
					inputFile = filepath.Join(inputDir, fmt.Sprintf("p%d.txt", num))
				}

				// Parse the file
				b, err := os.ReadFile(inputFile)
				if err != nil {
					return o.Annotatef(err, "failed to read file")
				}

				// Run the problem
				f(strings.Split(strings.TrimSpace(string(b)), "\n"), o)
				return nil
			}},
		),
		Executions: executions,
	}
}

func NoInputNode(num int, f func(command.Output), ex *Execution) *Problem {
	return &Problem{
		Num: num,
		N: commander.SerialNodes(
			DescNode(num),
			&commander.ExecutorProcessor{F: func(o command.Output, d *command.Data) error {
				f(o)
				return nil
			}},
		),
		Executions: []*Execution{ex},
	}
}

func NoInputWithExampleNode(num int, f func(output command.Output, example bool), executions []*Execution) *Problem {
	return &Problem{
		Num: num,
		N: commander.SerialNodes(
			DescNode(num),
			commander.FlagProcessor(
				exampleFlag,
			),
			&commander.ExecutorProcessor{F: func(o command.Output, d *command.Data) error {
				f(o, exampleFlag.Get(d))
				return nil
			}},
		),
		Executions: executions,
	}
}
