package y2017

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day23() aoc.Day {
	return &day23{}
}

type day23 struct{}

type cmd struct {
}

func (d *day23) Solve(lines []string, o command.Output) {
	registers := map[string]int{}

	numOrReg := func(k string) int {
		if v, ok := parse.AtoiOK(k); ok {
			return v
		}
		return registers[k]
	}

	mc := 0
	partsArr := parse.SplitWhitespace(lines)
	for i := 0; i < len(partsArr); i++ {
		parts := partsArr[i]
		a, b := numOrReg(parts[1]), numOrReg(parts[2])
		switch parts[0] {
		case "set":
			registers[parts[1]] = b
		case "mul":
			registers[parts[1]] = a * b
			mc++
		case "sub":
			registers[parts[1]] = a - b
		case "jnz":
			if a != 0 {
				i = i + b - 1
			}
		}
	}

	b := 93*100 + 100_000
	c := b + 17_000
	g := generator.Primes()
	var cnt int
	for k := b; k <= c; k += 17 {
		if !g.Contains(k) {
			cnt++
		}
	}
	o.Stdoutln(mc, cnt)
}

func (d *day23) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"0 911",
			},
		},
		{
			ExpectedOutput: []string{
				"8281 911",
			},
		},
	}
}

/*

TLDR: The program counts the number of composite (non-prime) numbers
from b to (b +17,000)

// Note: Ignore comment lines in jnz distances
set b 93
set c b
jnz a 2
jnz 1 5
mul b 100
sub b -100000
set c b
sub c -17000
// b = 109300
// c = 126300
// Start main loop

// Check all integers up to b and see if
// two numbers multiply to b
set f 1
set d 2
set e 2
set g d
mul g e
sub g b
jnz g 2
set f 0
sub e -1
set g e
sub g b
jnz g -8
sub d -1
set g d
sub g b
jnz g -13
jnz f 2
// Increment h if we found a pair
sub h -1
set g b
sub g c
jnz g 2
jnz 1 3
// Subtract 17 and try again
sub b -17
jnz 1 -23
*/
