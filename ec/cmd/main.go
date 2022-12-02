package main

import (
	"github.com/leep-frog/command"
	eulerchallenge "github.com/leep-frog/euler_challenge/ec"
)

func main() {
	command.RunNodes(command.AsNode(&command.BranchNode{
		Branches:          eulerchallenge.Branches(),
		DefaultCompletion: true,
		Default:           eulerchallenge.FileGenerator(),
	}))
}
