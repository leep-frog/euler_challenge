package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/command/commander"
	"github.com/leep-frog/command/sourcerer"
	eulerchallenge "github.com/leep-frog/euler_challenge/ec"
	"github.com/leep-frog/euler_challenge/maths/commandths"
)

func main() {
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalf("Failed to get file path from runtime.Caller()")
	}

	ecCmdDir := filepath.Dir(thisFile)
	ecDir := filepath.Dir(ecCmdDir)
	aliasers := sourcerer.Aliasers(map[string][]string{
		"e":  {"goleep", "-d", ecCmdDir, "ec"},
		"es": {"goleep", "-d", ecCmdDir, "ec", "-s"},
		"et": {"gt", "-v", ecDir, "-t", "6000"},
	})
	os.Exit(sourcerer.Source(
		"ecCLIs",
		[]sourcerer.CLI{
			&ecCLI{},
			commandths.CLI(),
		},
		aliasers,
		commandths.Aliasers(),
	))
}

type ecCLI struct{}

func (ecCLI) Name() string    { return "ec" }
func (ecCLI) Changed() bool   { return false }
func (ecCLI) Setup() []string { return nil }

func (ecCLI) Node() command.Node {
	startFlag := commander.BoolFlag("start", 's', "If not set, then START is outputted at the beginning of the command execution (so you can tell if delay is due to problem complexity or go compilation)")
	return commander.SerialNodes(
		commander.FlagProcessor(
			startFlag,
		),
		// commander.IfData(startFlag.Name(), commander.PrintlnProcessor("START", time.Now())),
		commander.SimpleProcessor(func(i *command.Input, o command.Output, d *command.Data, ed *command.ExecuteData) error {
			if !startFlag.Get(d) {
				o.Stdoutln("START", time.Now())
			}
			return nil
		}, nil),
		&commander.BranchNode{
			Branches:          eulerchallenge.Branches(),
			DefaultCompletion: true,
			Default:           eulerchallenge.FileGenerator(),
		},
		&commander.ExecutorProcessor{func(o command.Output, d *command.Data) error {
			if !startFlag.Get(d) {
				o.Stdoutln("END", time.Now())
			}
			return nil
		}},
	)
}
