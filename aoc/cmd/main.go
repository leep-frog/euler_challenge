package main

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoccmd"
)

func main() {
	command.RunNodes(aoccmd.CLI().Node())
}
