package main

import (
	"fmt"

	"github.com/leep-frog/command/command"
	twentyone "github.com/leep-frog/euler_challenge/aoc/2021"
	eulerchallenge "github.com/leep-frog/euler_challenge/ec"
)

func main() {
	fmt.Println("Needs fixing")
	// command.RunNodes(node())
}

// Separate method for testing purposes.
func node() command.Node {
	return &command.BranchNode{
		DefaultCompletion: true,
		Branches: map[string]command.Node{
			"1":  eulerchallenge.P1().Node(),
			"2":  eulerchallenge.P2().Node(),
			"3":  eulerchallenge.P3().Node(),
			"4":  eulerchallenge.P4().Node(),
			"5":  eulerchallenge.P5().Node(),
			"6":  eulerchallenge.P6().Node(),
			"7":  eulerchallenge.P7().Node(),
			"8":  eulerchallenge.P8().Node(),
			"9":  eulerchallenge.P9().Node(),
			"10": eulerchallenge.P10().Node(),
			"11": eulerchallenge.P11().Node(),
			"12": eulerchallenge.P12().Node(),
			"13": eulerchallenge.P13().Node(),
			"14": eulerchallenge.P14().Node(),
			"15": eulerchallenge.P15().Node(),
			"aoc": &command.BranchNode{
				DefaultCompletion: true,
				Branches: map[string]command.Node{
					"2021": &command.BranchNode{
						DefaultCompletion: true,
						Branches: map[string]command.Node{
							"19":   twentyone.D19(),
							"21":   twentyone.D21(),
							"21_2": twentyone.D21_2(),
							"22":   twentyone.D22(),
							"23":   twentyone.D23(),
						}},
				}},
		}}
}
