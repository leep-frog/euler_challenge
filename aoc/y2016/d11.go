package y2016

import (
	"fmt"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/rgx"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func Day11() aoc.Day {
	return &day11{}
}

type day11 struct{}

type floor struct {
	chips      map[string]bool
	generators map[string]bool
}

func (f *floor) Valid() bool {
	if len(f.generators) == 0 {
		return true
	}

	for c := range f.chips {
		if !f.generators[c] {
			return false
		}
	}
	return true
}

func (f *floor) String(seen map[string]int) string {
	cs := maps.Keys(f.chips)
	gs := maps.Keys(f.generators)
	slices.Sort(cs)
	slices.Sort(gs)

	var ss []string
	for _, c := range cs {
		if _, ok := seen[c]; !ok {
			seen[c] = len(seen)
		}
		ss = append(ss, fmt.Sprintf("CHIP-%d", seen[c]))
	}
	for _, g := range gs {
		if _, ok := seen[g]; !ok {
			seen[g] = len(seen)
		}
		ss = append(ss, fmt.Sprintf("GEN-%d", seen[g]))
	}
	slices.Sort(ss)
	return strings.Join(ss, " ")
}

func (f *floor) Copy() *floor {
	return &floor{
		maps.Clone(f.chips),
		maps.Clone(f.generators),
	}
}

func (b *building) String() string {
	var fs []string
	seen := map[string]int{}
	for i, f := range b.floors {
		prefix := fmt.Sprintf("%d", i+1)
		if i == b.elevator {
			prefix = "E"
		}
		fs = append(fs, fmt.Sprintf("%s %s", prefix, f.String(seen)))
	}
	return strings.Join(bread.Reverse(fs), "\n")
}

type building struct {
	floors   []*floor
	elevator int
}

func (b *building) Done() bool {
	for _, f := range b.floors[:len(b.floors)-1] {
		if len(f.chips) > 0 || len(f.generators) > 0 {
			return false
		}
	}
	return true
}

func (b *building) AStarEstimate() int {
	var est int
	for i, f := range b.floors {
		floorDist := len(b.floors) - 1 - i
		trips := (len(f.chips) + len(f.generators) + 1) / 2
		est += trips * floorDist
	}
	return est
}

func (b *building) Code() string {
	return b.String()
}

func (b *building) Copy(elevator int) *building {
	var fs []*floor
	for _, f := range b.floors {
		fs = append(fs, f.Copy())
	}
	return &building{fs, elevator}
}

func (b *building) Valid() bool {
	for _, f := range b.floors {
		if !f.Valid() {
			return false
		}
	}
	return true
}

func (b *building) AdjacentStates() []*building {
	var ns []*building
	fi := b.elevator

	f := b.floors[b.elevator]
	cs := maps.Keys(f.chips)
	slices.Sort(cs)
	gs := maps.Keys(f.generators)
	slices.Sort(gs)

	// Move all chips
	for i, c := range cs {
		// Move chip and generator
		if f.generators[c] {
			// Floor below
			if fi > 0 {
				newB := b.Copy(fi - 1)
				delete(newB.floors[fi].chips, c)
				delete(newB.floors[fi].generators, c)
				newB.floors[fi-1].chips[c] = true
				newB.floors[fi-1].generators[c] = true
				ns = append(ns, newB)
			}

			// Floor above
			if fi < len(b.floors)-1 {
				newB := b.Copy(fi + 1)
				delete(newB.floors[fi].chips, c)
				delete(newB.floors[fi].generators, c)
				newB.floors[fi+1].chips[c] = true
				newB.floors[fi+1].generators[c] = true
				ns = append(ns, newB)
			}
		}

		// Floor below
		if fi > 0 {
			newB := b.Copy(fi - 1)
			delete(newB.floors[fi].chips, c)
			newB.floors[fi-1].chips[c] = true
			ns = append(ns, newB)
		}

		// Floor above
		if fi < len(b.floors)-1 {
			newB := b.Copy(fi + 1)
			delete(newB.floors[fi].chips, c)
			newB.floors[fi+1].chips[c] = true
			ns = append(ns, newB)
		}

		for _, c2 := range cs[i+1:] {
			// Floor below
			if fi > 0 {
				newB := b.Copy(fi - 1)
				delete(newB.floors[fi].chips, c)
				delete(newB.floors[fi].chips, c2)
				newB.floors[fi-1].chips[c] = true
				newB.floors[fi-1].chips[c2] = true
				ns = append(ns, newB)
			}

			// Floor above
			if fi < len(b.floors)-1 {
				newB := b.Copy(fi + 1)
				delete(newB.floors[fi].chips, c)
				delete(newB.floors[fi].chips, c2)
				newB.floors[fi+1].chips[c] = true
				newB.floors[fi+1].chips[c2] = true
				ns = append(ns, newB)
			}
		}
	}

	// Move all generators
	for i, g := range gs {
		// Floor below
		if fi > 0 {
			newB := b.Copy(fi - 1)
			delete(newB.floors[fi].generators, g)
			newB.floors[fi-1].generators[g] = true
			ns = append(ns, newB)
		}

		// Floor above
		if fi < len(b.floors)-1 {
			newB := b.Copy(fi + 1)
			delete(newB.floors[fi].generators, g)
			newB.floors[fi+1].generators[g] = true
			ns = append(ns, newB)
		}

		for _, g2 := range gs[i+1:] {
			// Floor below
			if fi > 0 {
				newB := b.Copy(fi - 1)
				delete(newB.floors[fi].generators, g)
				delete(newB.floors[fi].generators, g2)
				newB.floors[fi-1].generators[g] = true
				newB.floors[fi-1].generators[g2] = true
				ns = append(ns, newB)
			}

			// Floor above
			if fi < len(b.floors)-1 {
				newB := b.Copy(fi + 1)
				delete(newB.floors[fi].generators, g)
				delete(newB.floors[fi].generators, g2)
				newB.floors[fi+1].generators[g] = true
				newB.floors[fi+1].generators[g2] = true
				ns = append(ns, newB)
			}
		}
	}

	var rs []*building
	for _, n := range ns {
		if n.Valid() {
			rs = append(rs, n)
		}
	}

	return rs
}

func (b *building) AdjacentStatesOld() []*building {
	var ns []*building

	fi := b.elevator
	f := b.floors[b.elevator]

	cs := maps.Keys(f.chips)
	slices.Sort(cs)
	gs := maps.Keys(f.generators)
	slices.Sort(gs)

	// Move all individual chips or chip/generator pairs
	for i, c := range cs {
		// If the generator is on the same floor, then they
		// can move together
		if f.generators[c] {

			// Floor below
			if fi > 0 && len(b.floors[fi-1].chips) == len(b.floors[fi-1].generators) {
				newB := b.Copy(fi - 1)
				delete(newB.floors[fi].chips, c)
				delete(newB.floors[fi].generators, c)
				newB.floors[fi-1].chips[c] = true
				newB.floors[fi-1].generators[c] = true
				ns = append(ns, newB)
			}

			// Floor above
			if fi < len(b.floors)-1 && len(b.floors[fi+1].chips) == len(b.floors[fi+1].generators) {
				newB := b.Copy(fi + 1)
				delete(newB.floors[fi].chips, c)
				delete(newB.floors[fi].generators, c)
				newB.floors[fi+1].chips[c] = true
				newB.floors[fi+1].generators[c] = true
				ns = append(ns, newB)
			}
		}

		// Floor below
		if fi > 0 && (len(b.floors[fi-1].generators) == 0 || b.floors[fi-1].generators[c]) {
			newB := b.Copy(fi - 1)
			delete(newB.floors[fi].chips, c)
			newB.floors[fi-1].chips[c] = true
			ns = append(ns, newB)
		}

		// Floor above
		if fi < len(b.floors)-1 && (len(b.floors[fi+1].generators) == 0 || b.floors[fi+1].generators[c]) {
			newB := b.Copy(fi + 1)
			delete(newB.floors[fi].chips, c)
			newB.floors[fi+1].chips[c] = true
			ns = append(ns, newB)
		}

		for _, c2 := range cs[i+1:] {
			// Floor below
			if fi > 0 && (len(b.floors[fi-1].generators) == 0 || (b.floors[fi-1].generators[c] && b.floors[fi-1].generators[c2])) {
				newB := b.Copy(fi - 1)
				delete(newB.floors[fi].chips, c)
				delete(newB.floors[fi].chips, c2)
				newB.floors[fi-1].chips[c] = true
				newB.floors[fi-1].chips[c2] = true
				ns = append(ns, newB)
			}

			// Floor above
			if fi < len(b.floors)-1 && (len(b.floors[fi+1].generators) == 0 || (b.floors[fi+1].generators[c] && b.floors[fi+1].generators[c2])) {
				newB := b.Copy(fi + 1)
				delete(newB.floors[fi].chips, c)
				delete(newB.floors[fi].chips, c2)
				newB.floors[fi+1].chips[c] = true
				newB.floors[fi+1].chips[c2] = true
				ns = append(ns, newB)
			}
		}
	}

	// Move generators
	for i, g := range gs {
		// Floor below if no chips on floor below or only the matching chip
		if fi > 0 && (len(b.floors[fi-1].chips) == 0 || (len(b.floors[fi-1].chips) == 1 && b.floors[fi-1].chips[g])) {
			newB := b.Copy(fi - 1)
			delete(newB.floors[fi].generators, g)
			newB.floors[fi-1].generators[g] = true
			ns = append(ns, newB)
		}

		// Floor above
		if fi < len(b.floors)-1 && (len(b.floors[fi+1].chips) == 0 || (len(b.floors[fi+1].chips) == 1 && b.floors[fi+1].chips[g])) {
			newB := b.Copy(fi + 1)
			delete(newB.floors[fi].generators, g)
			newB.floors[fi+1].generators[g] = true
			ns = append(ns, newB)
		}

		// Pair of generators
		for _, g2 := range gs[i+1:] {
			// Floor below
			if fi > 0 && (len(b.floors[fi-1].chips) == 0 || (len(b.floors[fi-1].chips) == 0 && b.floors[fi-1].chips[g] && b.floors[fi-1].chips[g2])) {
				newB := b.Copy(fi - 1)
				delete(newB.floors[fi].generators, g)
				delete(newB.floors[fi].generators, g2)
				newB.floors[fi-1].generators[g] = true
				newB.floors[fi-1].generators[g2] = true
				ns = append(ns, newB)
			}

			// Floor above
			if fi < len(b.floors)-1 && (len(b.floors[fi+1].chips) == 0 || (len(b.floors[fi+1].chips) == 0 && b.floors[fi+1].chips[g] && b.floors[fi+1].chips[g2])) {
				newB := b.Copy(fi + 1)
				delete(newB.floors[fi].generators, g)
				delete(newB.floors[fi].generators, g2)
				newB.floors[fi+1].generators[g] = true
				newB.floors[fi+1].generators[g2] = true
				ns = append(ns, newB)
			}
		}
	}

	return ns
}

func (d *day11) Solve(lines []string, o command.Output) {

	var floors []*floor
	for _, line := range lines {
		f := &floor{map[string]bool{}, map[string]bool{}}
		if strings.Contains(line, "nothing relevant") {
			floors = append(floors, f)
			continue
		}

		line = strings.ReplaceAll(line, ", and a ", ",")
		line = strings.ReplaceAll(line, " and a ", ",")
		line = strings.ReplaceAll(line, ", a ", ",")
		line = strings.ReplaceAll(line, " a ", ",")
		line = strings.TrimSuffix(line, ".")
		m := rgx.New("The [a-z]* floor contains,(.*)").MustMatch(line)[0]
		for _, comp := range strings.Split(m, ",") {
			words := strings.Split(comp, " ")
			if words[1] == "microchip" {
				f.chips[strings.TrimSuffix(words[0], "-compatible")] = true
			} else {
				f.generators[words[0]] = true
			}
		}

		floors = append(floors, f)
	}

	bOne := &building{floors, 0}
	bTwo := bOne.Copy(0)
	e, di := "elerium", "dilithium"
	bTwo.floors[0].chips[e] = true
	bTwo.floors[0].generators[e] = true
	bTwo.floors[0].chips[di] = true
	bTwo.floors[0].generators[di] = true

	_, dOne := bfs.AStarSearch[string]([]*building{bOne})
	_, dTwo := bfs.AStarSearch[string]([]*building{bTwo})

	o.Stdoutln(dOne, dTwo)
}

func (d *day11) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"11 0",
			},
		},
		{
			ExpectedOutput: []string{
				"37 61",
			},
		},
	}
}
