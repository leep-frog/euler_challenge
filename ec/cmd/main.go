package main

import (
	"os"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/command/commander"
	"github.com/leep-frog/command/sourcerer"
	eulerchallenge "github.com/leep-frog/euler_challenge/ec"
)

func main() {
	os.Exit(sourcerer.Source([]sourcerer.CLI{&ecCLI{}}))
}

type ecCLI struct{}

func (ecCLI) Name() string    { return "ec" }
func (ecCLI) Changed() bool   { return false }
func (ecCLI) Setup() []string { return nil }

func (ecCLI) Node() command.Node {
	startFlag := commander.BoolFlag("start", 's', "If set, then START is outputted at the beginning of the command execution (so you can tell if delay is due to problem complexity or go compilation)")
	return commander.SerialNodes(
		commander.FlagProcessor(
			startFlag,
		),
		commander.SimpleProcessor(func(i *command.Input, o command.Output, d *command.Data, ed *command.ExecuteData) error {
			if startFlag.Get(d) {
				o.Stdoutln("START")
			}
			return nil
		}, nil),
		&commander.BranchNode{
			Branches:          eulerchallenge.Branches(),
			DefaultCompletion: true,
			Default:           eulerchallenge.FileGenerator(),
		},
	)
}
