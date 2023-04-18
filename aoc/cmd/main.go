package main

import (
	"os"

	"github.com/leep-frog/command/sourcerer"
	"github.com/leep-frog/euler_challenge/aoc/aoccmd"
)

func main() {
	os.Exit(sourcerer.Source([]sourcerer.CLI{aoccmd.CLI()}))
}
