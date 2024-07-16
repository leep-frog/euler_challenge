package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/command/commander"
	"github.com/leep-frog/command/sourcerer"
	eulerchallenge "github.com/leep-frog/euler_challenge/ec"
)

func main() {
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalf("Failed to get file path from runtime.Caller()")
	}

	ecDir := filepath.Dir(thisFile)
	aliasers := sourcerer.Aliasers(map[string][]string{
		"e":  {"goleep", "-d", ecDir, "ec"},
		"es": {"goleep", "-d", ecDir, "ec", "-s"},
	})
	os.Exit(sourcerer.Source(
		[]sourcerer.CLI{&ecCLI{}},
		aliasers,
	))
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
