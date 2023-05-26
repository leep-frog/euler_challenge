package y2017

import (
	"sync/atomic"
	"time"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day18() aoc.Day {
	return &day18{}
}

type day18 struct{}

type program struct {
	registers    map[string]int
	part2        bool
	myChannel    chan int
	otherChannel chan int
	waiting      *atomic.Int32
}

func (p *program) run(partsArr [][]string) int {
	var sent int
	numOrReg := func(s string) int {
		if v, ok := parse.AtoiOK(s); ok {
			return v
		}
		return p.registers[s]
	}

	var sendCount int

	for i := 0; i < len(partsArr); i++ {
		parts := partsArr[i]
		switch parts[0] {
		case "set":
			p.registers[parts[1]] = numOrReg(parts[2])
		case "snd":
			sent = numOrReg(parts[1])
			if p.part2 {
				p.otherChannel <- sent
				sendCount++
			}
		case "add":
			p.registers[parts[1]] += numOrReg(parts[2])
		case "mul":
			p.registers[parts[1]] *= numOrReg(parts[2])
		case "mod":
			p.registers[parts[1]] %= numOrReg(parts[2])
		case "rcv":
			if !p.part2 {
				return sent
			}

			p.waiting.Add(1)

			if v, ok := <-p.myChannel; ok {
				p.waiting.Add(-1)
				p.registers[parts[1]] = v
			} else {
				return sendCount
			}
		case "jgz":
			if numOrReg(parts[1]) > 0 {
				i = i + numOrReg(parts[2]) - 1
			}
		}
	}
	return -1
}

func (d *day18) Solve(lines []string, o command.Output) {
	partsArr := parse.SplitWhitespace(lines)
	part1 := (&program{map[string]int{}, false, nil, nil, nil}).run(partsArr)

	c1, c2 := make(chan int, 100), make(chan int, 100)
	count := &atomic.Int32{}
	prg0 := &program{map[string]int{"p": 0}, true, c1, c2, count}
	prg1 := &program{map[string]int{"p": 1}, true, c2, c1, count}

	go func() {
		for count.Load() != 2 {
			time.Sleep(100 * time.Millisecond)
		}
		close(c1)
		close(c2)
	}()
	go prg0.run(partsArr)
	o.Stdoutln(part1, prg1.run(partsArr))
}

func (d *day18) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"",
			},
		},
		{
			ExpectedOutput: []string{
				"",
			},
		},
	}
}
