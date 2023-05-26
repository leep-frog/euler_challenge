package y2017

import (
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/linkedlist"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day17() aoc.Day {
	return &day17{}
}

type day17 struct{}

func (d *day17) Solve(lines []string, o command.Output) {
	cur := linkedlist.NewCircularList(0)

	num := parse.Atoi(lines[0])

	// for i := 0; i < 2017; i++ {
	for i := 0; i < 50_000_000; i++ {
		if i%100_000 == 0 {
			fmt.Println(i)
		}
		cur = cur.Nth(num % (i + 1))
		cur.PushAt(1, &linkedlist.Node[int]{Value: i + 1})
		cur = cur.Next
	}
	// o.Stdoutln(cur.Next.Value)
	n, _ := cur.Index(0)
	o.Stdoutln(n.Next.Value)
}

func (d *day17) Cases() []*aoc.Case {
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
