package y2020

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day25() aoc.Day {
	return &day25{}
}

type day25 struct{}

func (d *day25) Solve(lines []string, o command.Output) {
	var loopSizes, publicKeys []int
	for _, line := range lines {
		publicKey := parse.Atoi(line)
		subjectNumber := 7
		v := 1
		loopSize := 0
		for ; v != publicKey; loopSize++ {
			v = (v * subjectNumber) % 20201227
		}
		loopSizes = append(loopSizes, loopSize)
		publicKeys = append(publicKeys, publicKey)
	}

	var handshake int
	for i, loopSize := range loopSizes {
		subjectNumber := publicKeys[1-i]
		v := 1
		for j := 0; j < loopSize; j++ {
			v = (v * subjectNumber) % 20201227
		}
		handshake = v
	}
	o.Stdoutln(handshake)
}

func (d *day25) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"14897079",
			},
		},
		{
			ExpectedOutput: []string{
				"6011069",
			},
		},
	}
}
