package y2022

import (
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
)

func Day02() aoc.Day {
	return &day02{}
}

type RPCMove int

const (
	Rock RPCMove = 1 + iota
	Paper
	Scissors
)

type RPCResult int

const (
	Win RPCResult = 6 - 3*iota
	Draw
	Lose
)

func (m RPCMove) String() string {
	switch m {
	case Rock:
		return "Rock"
	case Paper:
		return "Paper"
	default:
		return "Scissors"
	}
}

func (m RPCMove) play(that RPCMove) int {
	// draw
	if m == that {
		return int(Draw) + int(m)
	}

	// win
	if (that%3)+1 == m {
		return int(Win) + int(m)
	}

	// lose
	return int(Lose) + int(m)
}

type day02 struct{}

func (d *day02) toMove(v string) RPCMove {
	switch v {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	default:
		return Scissors
	}
}

func (d *day02) toResult(v string) RPCResult {
	switch v {
	case "X":
		return Lose
	case "Y":
		return Draw
	default:
		return Win
	}
}

func (r RPCResult) produceFrom(that RPCMove) int {
	switch r {
	case Draw:
		return int(r) + int(that)
	case Win:
		return int(r) + (int(that%3) + 1)
	default:
		return int(r) + (int((that+1)%3) + 1)
	}
}

func (d *day02) Solve(lines []string, o command.Output) {
	var sum int
	for _, line := range lines {
		moves := strings.Split(line, " ")
		them, us := d.toMove(moves[0]), d.toMove(moves[1])
		v := us.play(them)
		sum += v
	}
	o.Stdoutln(sum)

	// Part 2
	sum = 0
	for _, line := range lines {
		moves := strings.Split(line, " ")
		them, result := d.toMove(moves[0]), d.toResult(moves[1])
		sum += result.produceFrom(them)
	}
	o.Stdoutln(sum)
}

func (d *day02) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"15",
				"12",
			},
		},
		{
			ExpectedOutput: []string{
				"11475",
				"16862",
			},
		},
	}
}
