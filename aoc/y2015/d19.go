package y2015

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/pair"
	"golang.org/x/exp/slices"
)

func Day19() aoc.Day {
	return &day19{}
}

type day19 struct{}

func (d *day19) donzo(depth int, mol string, transformations []*pair.Pair[string, string]) (int, bool) {
	if mol == "e" {
		return depth, true
	}

	for _, t := range transformations {
		parts := strings.Split(mol, t.A)
		for i := 1; i < len(parts); i++ {
			if v, ok := d.donzo(depth+1, strings.Join(parts[:i], t.A)+t.B+strings.Join(parts[i:], t.A), transformations); ok {
				return v, ok
			}
		}
	}
	return 0, false
}

func (d *day19) Solve(lines []string, o command.Output) {
	// Generate objects
	ops := map[string][]string{}
	revOps := map[string][]string{}
	var maxLen, revMaxLen int
	for _, line := range lines[:len(lines)-2] {
		parts := strings.Split(line, " => ")
		ops[parts[0]] = append(ops[parts[0]], parts[1])
		revOps[parts[1]] = append(revOps[parts[1]], parts[0])
		maxLen = maths.Max(maxLen, len(parts[0]))
		revMaxLen = maths.Max(revMaxLen, len(parts[1]))
	}
	mol := lines[len(lines)-1]
	mc := &moleculeContext{ops, maxLen}
	revMC := &moleculeContext{revOps, revMaxLen}

	// Solve part 1
	part1 := len(mc.transformations(mol))

	// Generate array of pairs
	var ts []*pair.Pair[string, string]
	for from, tos := range revMC.ops {
		for _, to := range tos {
			ts = append(ts, pair.New(from, to))
		}
	}
	slices.SortFunc(ts, func(this, that *pair.Pair[string, string]) bool {
		thisDist := len(this.A) - len(this.B)
		thatDist := len(that.A) - len(that.B)
		return thisDist > thatDist
	})

	part2, _ := d.donzo(0, mol, ts)
	o.Stdoutln(part1, part2)
}

type moleculeContext struct {
	ops    map[string][]string
	maxLen int
}

func (mc *moleculeContext) transformations(molecule string) map[string]bool {
	transformations := map[string]bool{}
	for i := range molecule {
		for size := 1; size <= mc.maxLen; size++ {
			if i+size > len(molecule) {
				break
			}
			for _, v := range mc.ops[molecule[i:i+size]] {
				newMol := molecule[:i] + v
				if i+size < len(molecule) {
					newMol += molecule[i+size:]
				}
				transformations[newMol] = true
			}
		}
	}
	return transformations
}

func (d *day19) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"4 3",
			},
		},
		{
			ExpectedOutput: []string{
				"518 200",
			},
		},
	}
}
