package y2018

import (
	"fmt"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/linkedlist"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/rgx"
)

func Day05() aoc.Day {
	return &day05{}
}

type day05 struct{}

func (d *day05) Solve(lines []string, o command.Output) {
	polymer := lines[0]
	part1 := d.polymerLength(polymer)

	m := map[string]bool{}
	for _, c := range strings.ToLower(polymer) {
		m[string(c)] = true
	}

	best := maths.Smallest[int, int]()
	for k := range m {
		r := rgx.New(fmt.Sprintf("[%s%s]+", strings.ToLower(k), strings.ToUpper(k)))
		best.Check(d.polymerLength(r.ReplaceAll(polymer, "")))
	}

	o.Stdoutln(part1, best.Best())
}

func (d *day05) polymerLength(polymerStr string) int {
	length := len(polymerStr)
	polymer := linkedlist.NewList(strings.Split(polymerStr, "")...)
	start := &linkedlist.Node[string]{
		Value: "9",
		Next:  polymer,
	}
	polymer.Prev = start
	polymer = start
	for node := polymer; node.Next != nil; {
		cur, next := node.Value, node.Next.Value
		if strings.ToUpper(cur) == strings.ToUpper(next) && cur != next {
			length -= 2
			if node.Prev == nil {
				node = node.Next.Next
				node.Prev = nil
			} else if node.Next.Next == nil {
				node = node.Prev
				node.Next = nil
			} else {
				prevN, nextN := node.Prev, node.Next.Next
				prevN.Next = nextN
				nextN.Prev = prevN
				node = prevN
			}
		} else {
			node = node.Next
		}
	}

	return length
}

func (d *day05) Cases() []*aoc.Case {
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
