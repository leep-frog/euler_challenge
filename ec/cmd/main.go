package main

import (
	"github.com/leep-frog/command"
	eulerchallenge "github.com/leep-frog/euler_challenge/ec"
)

func main() {
	command.RunNodes(command.BranchNode(
		eulerchallenge.Branches(), nil, false,
	))
}
