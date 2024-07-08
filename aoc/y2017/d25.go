package y2017

import (
	"regexp"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/euler_challenge/rgx"
)

func Day25() aoc.Day {
	return &day25{}
}

type stateMachine struct {
	writeValueOne bool
	moveLeft      bool
	nextState     string
}

type day25 struct{}

func (d *day25) Solve(lines []string, o command.Output) {
	ignoreChars := regexp.MustCompile(`[-:\.]`)
	for i, line := range lines {
		lines[i] = ignoreChars.ReplaceAllString(line, "")
	}

	r0 := rgx.New(`Begin in state ([A-Z]+)$`)
	r1 := rgx.New(` ([0-9]+) `)

	curState := r0.MustMatch(lines[0])[0]
	steps := parse.Atoi(r1.MustMatch(lines[1])[0])

	// Turing maching (map from state tp 0/1 to state machine node).
	tm := map[string]map[int]*stateMachine{}

	partsArr := parse.SplitWhitespace(lines)
	for i := 3; i < len(partsArr); i += 10 {
		state := partsArr[i][2]

		zeroState := &stateMachine{
			writeValueOne: parse.Atoi(partsArr[i+2][3]) == 1,
			moveLeft:      partsArr[i+3][5] == "left",
			nextState:     partsArr[i+4][3],
		}
		oneState := &stateMachine{
			writeValueOne: parse.Atoi(partsArr[i+6][3]) == 1,
			moveLeft:      partsArr[i+7][5] == "left",
			nextState:     partsArr[i+8][3],
		}

		maths.Insert(tm, state, 0, zeroState)
		maths.Insert(tm, state, 1, oneState)
	}

	ones := map[int]int{}
	var index int
	for step := 0; step < steps; step++ {
		curValue := ones[index]

		sm := tm[curState][curValue]
		if sm.writeValueOne {
			ones[index] = 1
		} else {
			delete(ones, index)
		}
		if sm.moveLeft {
			index--
		} else {
			index++
		}
		curState = sm.nextState
	}

	o.Stdoutln(len(ones))
}

func (d *day25) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"3",
			},
		},
		{
			ExpectedOutput: []string{
				"3554",
			},
		},
	}
}
