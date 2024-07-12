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
	return &commander.BranchNode{
		Branches:          eulerchallenge.Branches(),
		DefaultCompletion: true,
		Default:           eulerchallenge.FileGenerator(),
	}
}
