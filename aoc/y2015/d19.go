package y2015

import (
	"fmt"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/maths"
)

func Day19() aoc.Day {
	return &day19{}
}

type day19 struct{}

type molecule struct {
	m     string
	depth int
}

func (m *molecule) Code(mc *moleculeContext) string {
	return fmt.Sprintf("%s %d", m.m, m.depth)
}

/*func (m *molecule) Distance(mc *moleculeContext) bfs.Int {
	return bfs.Int(len(m.m))
}*/

var (
	dpth = 0
)

func (m *molecule) Done(mc *moleculeContext) bool {
	if m.depth > dpth {
		dpth = m.depth
		fmt.Println("DEPTH", dpth)
	}
	return m.m == "e"
}

func (m *molecule) AdjacentStates(mc *moleculeContext) []*molecule {
	var r []*molecule
	for t := range mc.transformations(m.m) {
		r = append(r, &molecule{t, m.depth + 1})
	}
	return r
	/*if mc.best.Set() && m.depth >= mc.best.Best() {
		return nil
	}

	var r []*molecule
	trans := mc.transformations(m.m)
	slices.SortFunc(maps.Keys(trans), func(a, b string) bool {
		return len(a) < len(b)
	})
	for k := range trans {
		r = append(r, &molecule{k, m.depth + 1})
	}
	return r*/
}

func (d *day19) rec2(oneof map[string]bool, mc *moleculeContext, mol string, depth int, cache map[string]int, best *maths.Bester[int, int]) int {
	// Solve in sections of `Rn[^rRn]*r`
	return 0
}

func (d *day19) rec(mc *moleculeContext, mol string, depth int, cache map[string]int, best *maths.Bester[int, int]) int {
	// if best.Set() && depth >= best.Best() {
	// 	return
	// }

	if v, ok := cache[mol]; ok {
		return v
	}

	if mol == "e" {
		best.Check(depth)
		// fmt.Println("HURRAY", depth)
		cache[mol] = 0
		return 0
	}

	if depth > 5 {
		return 1_000_000_000
	}

	// if v, ok := cache[mol]; ok {
	// 	betterOne := maths.Min(v[0], depth)
	// 	best.Check(betterOne + v[1])
	// 	cache[mol] = []int{betterOne, v[1]}
	// 	return
	// }

	// bestAfter := -1
	shortest := maths.Smallest[int, int]()
	for newMol := range mc.transformations(mol) {
		shortest.Check(d.rec(mc, newMol, depth+1, cache, best))
	}
	cache[mol] = shortest.Best() + 1
	return shortest.Best() + 1
	// cache[mol] = []int{}
}

func (d *day19) solve2(depth int, mol string, mc *moleculeContext, leftOfRn map[string]bool, rightOfRn map[string]bool, best *maths.Bester[int, int]) {
	parts := strings.SplitN(mol, "Rn", 2)

	if len(parts) == 1 {
		_, dist := bfs.ContextSearch[*moleculeContext, string](mc, []*molecule{&molecule{mol, 0}})
		best.Check(dist)
		return
	}

	leftOps
	for opt, dist := range mc.Explore(mol, leftOfRn) {

	}
}

func (d *day19) Solve(lines []string, o command.Output) {
	fmt.Println("START")
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
	mc := &moleculeContext{ops, maxLen, mol, nil}
	revMC := &moleculeContext{revOps, revMaxLen, "e", maths.Smallest[int, int]()}

	part1 := len(mc.transformations(mol))
	fmt.Println("UNO", part1)

	leftOfRnArr := map[string][]string{}
	leftOfRnBool := map[string]bool{}
	for conv := range mc.ops {
		if parts := strings.Split(conv, "Rn"); len(parts) > 1 {
			leftOfRnArr[parts[0]] = append(leftOfRnArr[parts[0]], parts[1])
			leftOfRnBool[parts[0]] = true
		}
	}

	best := maths.Smallest[int, int]()
	d.solve2(0, mol, revMC, leftOfRnBool, best)
	fmt.Println(best)

	/*
		path, dist := bfs.ContextSearch[*moleculeContext, string](revMC, []*molecule{{mol, 0}})
		fmt.Println(dist, path)*/
	// o.Stdoutln(part1, revMC.best.Best())

	// b := maths.Smallest[int, int]()
	// d.rec(revMC, mol, 0, map[string]int{}, b)
	// fmt.Println(b)
}

type moleculeContext struct {
	ops    map[string][]string
	maxLen int
	target string
	best   *maths.Bester[int, int]
}

func (mc *moleculeContext) Explore(molecule string, valid map[string]bool) map[string]int {
	r := map[string]int{}
	mc.explore(molecule, valid, r, 0)
	return r
}

func (mc *moleculeContext) explore(molecule string, valid map[string]bool, values map[string]int, depth int) {
	if valid[molecule] {
		if v, ok := values[molecule]; ok {
			values[molecule] = maths.Min(v, depth)
		} else {
			values[molecule] = depth
		}
	}

	for t := range mc.transformations(molecule) {
		mc.explore(t, valid, values, depth+1)
	}
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
