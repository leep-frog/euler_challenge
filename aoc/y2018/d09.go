package y2018

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/linkedlist"
	"github.com/leep-frog/euler_challenge/maths"
	"golang.org/x/exp/maps"
)

func Day09() aoc.Day {
	return &day09{}
}

type day09 struct{}

func (d *day09) Solve(lines []string, o command.Output) {
	p, m := 424, 71144
	o.Stdoutln(d.solve(p, m), d.solve(p, m*100))
}

func (d *day09) solve(numPlayers, marbles int) int {
	scores := map[int]int{}
	cur := linkedlist.NewCircularList(0)
	for next, player := 1, 1; next <= marbles; next, player = next+1, (player+1)%numPlayers {
		if next%23 != 0 {
			left, right := cur.Next, cur.Next.Next
			new := &linkedlist.Node[int]{
				Value: next,
				Prev:  left,
				Next:  right,
			}
			left.Next = new
			right.Prev = new
			cur = new
		} else {
			scores[player] += next
			for i := 0; i < 8; i++ {
				cur = cur.Prev
			}
			scores[player] += cur.PopNext().Value
			cur = cur.Next
		}
	}
	return maths.Max(maps.Values(scores)...)
}

func (d *day09) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"405143 3411514667",
			},
		},
		{
			ExpectedOutput: []string{
				"405143 3411514667",
			},
		},
	}
}
