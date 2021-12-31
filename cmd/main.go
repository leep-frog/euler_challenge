package main

import (
	"github.com/leep-frog/command"
	twentyone "github.com/leep-frog/euler_challenge/aoc/2021"
	eulerchallenge "github.com/leep-frog/euler_challenge/ec"
)

func main() {
	command.RunNodes(node())
}

// Separate method for testing purposes.
func node() *command.Node {
	return command.BranchNode(map[string]*command.Node{
		"1":  eulerchallenge.P1(),
		"2":  eulerchallenge.P2(),
		"3":  eulerchallenge.P3(),
		"4":  eulerchallenge.P4(),
		"5":  eulerchallenge.P5(),
		"6":  eulerchallenge.P6(),
		"7":  eulerchallenge.P7(),
		"8":  eulerchallenge.P8(),
		"9":  eulerchallenge.P9(),
		"10": eulerchallenge.P10(),
		"11": eulerchallenge.P11(),
		"12": eulerchallenge.P12(),
		"13": eulerchallenge.P13(),
		"14": eulerchallenge.P14(),
		"15": eulerchallenge.P15(),
		"aoc": command.BranchNode(map[string]*command.Node{
			"2021": command.BranchNode(map[string]*command.Node{
				"19":   twentyone.D19(),
				"21":   twentyone.D21(),
				"21_2": twentyone.D21_2(),
				"22":   twentyone.D22(),
				"23":   twentyone.D23(),
			}, nil, true),
		}, nil, true),
	}, nil, true)
}
