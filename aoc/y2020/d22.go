package y2020

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day22() aoc.Day {
	return &day22{}
}

type day22 struct{}

func (d *day22) Solve(lines []string, o command.Output) {
	var player2 bool
	deck1, deck2 := maths.NewQueue[int](), maths.NewQueue[int]()
	for i := 1; i < len(lines); i++ {
		if lines[i] == "" {
			player2 = true
			i++
			continue
		}

		if player2 {
			deck2.Push(parse.Atoi(lines[i]))
		} else {
			deck1.Push(parse.Atoi(lines[i]))
		}
	}

	part2, _ := d.recursiveComabt(deck1.Copy(), deck2.Copy(), map[string]bool{})
	o.Stdoutln(
		d.score(d.combat(deck1.Copy(), deck2.Copy())),
		d.score(part2),
	)
}

// Returns whether player 1 won or not
func (d *day22) recursiveComabt(deck1, deck2 *maths.Queue[int], config map[string]bool) (*maths.Queue[int], bool) {
	for !deck1.IsEmpty() && !deck2.IsEmpty() {
		code := fmt.Sprintf("{%v, %v}", deck1, deck2)
		if config[code] {
			return deck1, true
		}
		config[code] = true

		c1, c2 := deck1.Pop(), deck2.Pop()
		var winner1 bool
		if deck1.Len() < c1 || deck2.Len() < c2 {
			winner1 = c1 > c2
		} else {
			_, winner1 = d.recursiveComabt(deck1.SliceCopy(0, c1), deck2.SliceCopy(0, c2), map[string]bool{})
		}

		if winner1 {
			deck1.Push(c1, c2)
		} else {
			deck2.Push(c2, c1)
		}
	}

	if deck2.IsEmpty() {
		return deck1, true
	}
	return deck2, false
}

func (d *day22) score(deck *maths.Queue[int]) int {
	var sum int
	for !deck.IsEmpty() {
		sum += deck.Len() * deck.Pop()
	}

	return sum
}

func (d *day22) combat(deck1, deck2 *maths.Queue[int]) *maths.Queue[int] {
	for !deck1.IsEmpty() && !deck2.IsEmpty() {
		c1, c2 := deck1.Pop(), deck2.Pop()
		if c1 > c2 {
			deck1.Push(c1, c2)
		} else {
			deck2.Push(c2, c1)
		}
	}

	if deck2.IsEmpty() {
		return deck1
	}
	return deck2
}

func (d *day22) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"306 291",
			},
		},
		{
			ExpectedOutput: []string{
				"32102 34173",
			},
		},
	}
}
