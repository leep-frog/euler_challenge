package y2018

import (
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/point"
	"github.com/leep-frog/euler_challenge/rgx"
)

func Day10() aoc.Day {
	return &day10{}
}

type day10 struct{}

type particle struct {
	position, velocity *point.Point[int]
}

func (d *day10) Solve(lines []string, o command.Output) {
	r := rgx.New(`^position=< *(-?[0-9]*), *(-?[0-9]*)> velocity=< *(-?[0-9]*), *(-?[0-9]*)>$`)
	var particles []*particle
	for _, line := range lines {
		m := r.MatchInts(line)
		particles = append(particles, &particle{
			point.New(m[0], m[1]),
			point.New(m[2], m[3]),
		})
	}

	var prev, prevPrev int
	var minX, maxX, minY, maxY *maths.Bester[int, int]
	var iter int
	for ; iter <= 1 || prev < prevPrev; iter++ {
		for _, p := range particles {
			p.position.X += p.velocity.X
			p.position.Y += p.velocity.Y
		}
		minX, maxX = maths.Smallest[int, int](), maths.Largest[int, int]()
		minY, maxY = maths.Smallest[int, int](), maths.Largest[int, int]()
		for _, p := range particles {
			minX.Check(p.position.X)
			maxX.Check(p.position.X)
			minY.Check(p.position.Y)
			maxY.Check(p.position.Y)
		}
		prevPrev, prev = prev, (maxX.Best()-minX.Best())*(maxY.Best()-minY.Best())
	}

	for _, p := range particles {
		p.position.X -= p.velocity.X
		p.position.Y -= p.velocity.Y
	}
	minX, maxX = maths.Smallest[int, int](), maths.Largest[int, int]()
	minY, maxY = maths.Smallest[int, int](), maths.Largest[int, int]()
	for _, p := range particles {
		minX.Check(p.position.X)
		maxX.Check(p.position.X)
		minY.Check(p.position.Y)
		maxY.Check(p.position.Y)
	}

	o.Stdoutln(iter - 1)
	d.printParticles(o, particles, minX, maxX, minY, maxY)

}

func (d *day10) printParticles(o command.Output, particles []*particle, minX, maxX, minY, maxY *maths.Bester[int, int]) {
	pMap := map[int]map[int]bool{}
	for _, p := range particles {
		maths.Insert(pMap, p.position.Y, p.position.X, true)
	}
	var pic []string
	for y := minY.Best(); y <= maxY.Best(); y++ {
		var row []string
		for x := minX.Best(); x <= maxX.Best(); x++ {
			if pMap[y][x] {
				row = append(row, "#")
			} else {
				row = append(row, ".")
			}
		}
		pic = append(pic, strings.Join(row, ""))
	}

	o.Stdoutln(strings.Join(pic, "\n"))
}

func (d *day10) Cases() []*aoc.Case {
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
