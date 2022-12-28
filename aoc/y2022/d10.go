package y2022

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day10() aoc.Day {
	return &day10{}
}

type day10 struct{}

func (d *day10) Solve(lines []string, o command.Output) {
	var cycles []int
	var crt []bool
	x := 1
	sprite := 1
	for _, line := range lines {

		if line == "noop" {
			cycles = append(cycles, x)
			crt = append(crt, maths.Abs(sprite-(len(crt)%40)) <= 1)
		} else {
			crt = append(crt, maths.Abs(sprite-(len(crt)%40)) <= 1)
			crt = append(crt, maths.Abs(sprite-(len(crt)%40)) <= 1)
			delta := parse.Atoi(strings.Split(line, " ")[1])
			cycles = append(cycles, x, x)
			x += delta
			sprite = (sprite + delta) % 40
		}
	}

	var sum int
	for i := 19; i < len(cycles); i += 40 {
		sum += cycles[i] * (i + 1)
	}
	o.Stdoutln(sum)

	var r []string
	for i, c := range crt {
		if c {
			r = append(r, "#")
		} else {
			r = append(r, ".")
		}
		if (i+1)%40 == 0 {
			r = append(r, "\n")
		}
	}
	o.Stdoutln(strings.Join(r, ""))
}

func (d *day10) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"13140",
				"##..##..##..##..##..##..##..##..##..##..",
				"###...###...###...###...###...###...###.",
				"####....####....####....####....####....",
				"#####.....#####.....#####.....#####.....",
				"######......######......######......####",
				"#######.......#######.......#######.....",
				"",
			},
		},
		{
			ExpectedOutput: []string{
				"11780",
				"###..####.#..#.#....###...##..#..#..##..",
				"#..#....#.#..#.#....#..#.#..#.#..#.#..#.",
				"#..#...#..#..#.#....###..#..#.#..#.#..#.",
				"###...#...#..#.#....#..#.####.#..#.####.",
				"#....#....#..#.#....#..#.#..#.#..#.#..#.",
				"#....####..##..####.###..#..#..##..#..#.",
				"",
			},
		},
	}
}
