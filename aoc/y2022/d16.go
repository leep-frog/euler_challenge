package y2022

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/functional"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func Day16() aoc.Day {
	return &day16{}
}

type day16 struct{}

func valvesCode(m map[string]bool) string {
	keys := maps.Keys(m)
	slices.Sort(keys)
	return strings.Join(keys, " ")
}

// Visit each node once per path and always assume we will turn it on.
// Keep track of the best we can do for any set of turned-on valves (aka the key of the 'bests' map).
func rec16(path []string, onValves map[string]bool, rem, flow int, v *valve, bests map[string]*maths.Bester[[]string, int]) {
	if rem >= 0 {
		code := valvesCode(onValves)
		if bests[code] == nil {
			bests[code] = maths.Largest[[]string, int]()
		}
		// Only clone slice if necessary.
		if flow > bests[code].Best() || !bests[code].Set() {
			bests[code].IndexCheck(slices.Clone(path), flow)
		}
	}
	if rem <= 0 {
		return
	}

	// Turn the valve on
	v.on = true
	if v.flow > 0 {
		rem--
		flow += v.flow * rem
		onValves[v.id] = true
	}

	code := valvesCode(onValves)
	if bests[code] == nil {
		bests[code] = maths.Largest[[]string, int]()
	}
	// Only clone slice if necessary.
	if flow > bests[code].Best() {
		bests[code].IndexCheck(slices.Clone(path), flow)
	}

	// Go to neighbor
	for _, vt := range v.neighbors {
		// Only go to valves that are off
		if vt.valve.on {
			continue
		}
		rec16(append(path, vt.valve.id), onValves, rem-vt.distance, flow, vt.valve, bests)
	}
	v.on = false
	delete(onValves, v.id)
}

type valveTunnel struct {
	valve    *valve
	distance int
}

type valve struct {
	id        string
	on        bool
	flow      int
	neighbors []*valveTunnel
}

type valveContext struct {
	vMap map[string][]*valve
	root *valve
}

func (vt *valveTunnel) String() string {
	return fmt.Sprintf("[%s %d]", vt.valve.id, vt.distance)
}

func (v *valve) String() string {
	return fmt.Sprintf("%s %v %d %v", v.id, v.on, v.flow, v.neighbors)
}

func (v *valve) Code(ctx *valveContext, path bfs.Path[*valve]) string {
	return v.id
}

func (v *valve) Done(ctx *valveContext, path bfs.Path[*valve]) bool {
	return false
}

func (v *valve) AdjacentStates(ctx *valveContext, path bfs.Path[*valve]) []*valve {
	if v.flow != 0 && v.id != ctx.root.id {
		ctx.root.neighbors = append(ctx.root.neighbors, &valveTunnel{v, path.Len() - 1})
	}
	return ctx.vMap[v.id]
}

func (d *day16) Solve(lines []string, o command.Output) {
	r := regexp.MustCompile("Valve ([A-Z]+) has flow rate=([0-9]+); tunnels? leads? to valves? ((?:[A-Z]+,? ?)+)")
	// Map from valve id to valve
	valves := map[string]*valve{}
	// Map from valve id to neighbor valve IDs
	vMapStr := map[string][]string{}
	// Map from valve id to neighbor valves
	vMap := map[string][]*valve{}

	// Parse lines and populate maps
	for _, line := range lines {
		m := r.FindStringSubmatch(line)
		v := &valve{m[1], false, parse.Atoi(m[2]), nil}

		neighbors := strings.Split(m[3], ", ")
		valves[v.id] = v
		vMapStr[v.id] = neighbors
	}

	// Populate vMap from vMapStr
	for v, n := range vMapStr {
		vMap[v] = functional.Map(n, func(id string) *valve { return valves[id] })
	}

	// Now populate valve.neighbors (and distance to them) for all neighbors.
	// These neighbors are the time it takes to get to valves with non-zero flow.
	// This way, we can assume that we are always going to turn on a valve when we
	// get to it (and not have to worry about re-visiting any valves).
	for _, v := range valves {
		if v.flow != 0 || v.id == "AA" {
			bfs.ContextPathSearch[string](&valveContext{vMap, v}, []*valve{v})
		}
	}

	// Part 1: Get the best we can do in 30 minutes.
	bests := map[string]*maths.Bester[[]string, int]{}
	rec16(nil, map[string]bool{}, 30, 0, valves["AA"], bests)
	absoluteBest := maths.Largest[[]string, int]()
	for _, b := range bests {
		absoluteBest.IndexCheck(b.BestIndex(), b.Best())
	}

	// Part 2: Compare all pairs of paths that don't visit the same valve.
	secondBests := map[string]*maths.Bester[[]string, int]{}
	rec16(nil, map[string]bool{}, 26, 0, valves["AA"], secondBests)
	secondAbsoluteBest := maths.Largest[[][]string, int]()
	for code, best := range secondBests {
		set := maths.NewSimpleSet(strings.Split(code, " ")...)
		for code2, best2 := range secondBests {
			set2 := maths.NewSimpleSet(strings.Split(code2, " ")...)

			// Ensure no elements in common, otherwise we did redundant work.
			if len(maths.Intersection(set, set2)) != 0 {
				continue
			}

			if best.Best()+best2.Best() > secondAbsoluteBest.Best() || !secondAbsoluteBest.Set() {
				index := [][]string{best.BestIndex(), best2.BestIndex()}
				// Sort indices so order is deterministic
				slices.SortFunc(index, func(this, that []string) bool {
					if len(this) != len(that) {
						return len(this) < len(that)
					}
					if len(this) == 0 {
						return true
					}
					return this[0] < that[0]
				})
				secondAbsoluteBest.IndexCheck(index, best.Best()+best2.Best())
			}
		}
	}

	o.Stdoutln(absoluteBest.Best(), secondAbsoluteBest.Best())
}

func (d *day16) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"1651 1707",
			},
		},
		{
			ExpectedOutput: []string{
				"2250 3015",
			},
		},
	}
}
