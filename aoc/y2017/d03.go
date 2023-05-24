package y2017

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/euler_challenge/point"
	"github.com/leep-frog/euler_challenge/walker"
)

func Day03() aoc.Day {
	return &day03{}
}

type day03 struct{}

func (d *day03) Solve(lines []string, o command.Output) {
	k := parse.Atoi(lines[0])

	/*
		This was a cool way to solve for part1, but we still need the second part for part two
		for p := 3; ; p += 2 {
			p2 := p * p
			if k > p2 {
				continue
			}
			shells := p / 2
			toMiddle := maths.Abs(((p2 - k) % (p - 1)) - p/2)
			fmt.Println(shells + toMiddle)
			break
		}*/

	positions := map[int]map[int]int{
		0: {0: 1},
	}
	positions2 := map[int]map[int]int{
		0: {0: 1},
	}

	idx := 1
	var part1, part2 int
	pos := point.New(0, 0)

	insert := func() {
		idx++
		maths.Insert(positions, pos.X, pos.Y, idx)
		if idx == k && part1 == 0 {
			part1 = pos.ManhattanDistance(point.Origin[int]())
		}

		var sum int
		for _, nb := range walker.NeighborsWithDiagonals() {
			sum += positions2[pos.X+nb.X][pos.Y+nb.Y]
		}
		maths.Insert(positions2, pos.X, pos.Y, sum)
		if sum > k && part2 == 0 {
			part2 = sum
		}
	}

	for square := 3; ; square += 2 {

		// To the right
		pos.X++
		insert()

		// Up
		for i := 0; i < square-2; i++ {
			pos.Y++
			insert()
		}

		// Left
		for i := 0; i < square-1; i++ {
			pos.X--
			insert()
		}

		// Down
		for i := 0; i < square-1; i++ {
			pos.Y--
			insert()
		}

		// Right
		for i := 0; i < square-1; i++ {
			pos.X++
			insert()
		}

		if part1 != 0 && part2 != 0 {
			break
		}
	}
	o.Stdoutln(part1, part2)
}

func (d *day03) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"31 1968",
			},
		},
		{
			ExpectedOutput: []string{
				"480 349975",
			},
		},
	}
}
