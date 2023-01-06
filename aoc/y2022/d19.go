package y2022

import (
	"regexp"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/functional"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day19() aoc.Day {
	return &day19{}
}

type day19 struct{}

type resource int

const (
	ore resource = iota
	clay
	obsidian
	geode
	resourceCount
)

func (r resource) String() string {
	switch r {
	case ore:
		return "ore"
	case clay:
		return "clay"
	case obsidian:
		return "obsidian"
	case geode:
		return "geode"
	}
	panic("Unknown resource")
}

type blueprint struct {
	id           int
	robots       []*robot
	maxOreRobots int
}

type robot struct {
	requirements []int
}

func (r *robot) construct(resources []int) {
	for res, cnt := range r.requirements {
		resources[res] -= cnt
	}
}

func (r *robot) deconstruct(resources []int) {
	for res, cnt := range r.requirements {
		resources[res] += cnt
	}
}

func (r *robot) until(resources []int, robots []int) (int, bool) {
	var until int
	for res, required := range r.requirements {
		moreNeeded := required - resources[res]
		if moreNeeded <= 0 {
			continue
		}
		if robots[res] == 0 {
			return 0, false
		}
		until = maths.Max(until, (moreNeeded+robots[res]-1)/robots[res])
	}

	return until, true
}

func (d *day19) elapse(minutes int, resources, robots []int) {
	for res, cnt := range robots {
		resources[res] += cnt * minutes
	}
}

func (d *day19) unelapse(minutes int, resources, robots []int) {
	for res, cnt := range robots {
		resources[res] -= cnt * minutes
	}
}

func (d *day19) evalBlueprint(minutes int, bp *blueprint, resources, robots []int, best *maths.Bester[int, int]) {
	if minutes <= 0 {
		return
	}
	best.Check(resources[geode] + robots[geode]*minutes)
	if minutes == 0 {
		return
	}

	// The best we can possibly do is construct one per minute. See if current path can no longer beat that.
	maxBest := resources[geode] + robots[geode]*minutes
	maxBest += (minutes * (minutes - 1)) / 2
	if best.Best() >= maxBest {
		return
	}

	// Otherwise, iterate through the rest of the tree
	for res, rbt := range bp.robots {
		if res == int(ore) && robots[ore] >= bp.maxOreRobots {
			continue
		}
		if mins, ok := rbt.until(resources, robots); ok {
			d.elapse(mins+1, resources, robots)
			robots[res]++
			rbt.construct(resources)
			d.evalBlueprint(minutes-mins-1, bp, resources, robots, best)
			rbt.deconstruct(resources)
			robots[res]--
			d.unelapse(mins+1, resources, robots)
		}
	}
}

func (d *day19) solveBlueprint(bp *blueprint, minutes int) int {
	resources := make([]int, resourceCount, resourceCount)
	robots := make([]int, resourceCount, resourceCount)
	robots[ore] = 1
	best := maths.Largest[int, int]()
	d.evalBlueprint(minutes, bp, resources, robots, best)
	return best.Best()
}

func (d *day19) Solve(lines []string, o command.Output) {
	r := regexp.MustCompile("^Blueprint ([0-9]+): Each ore robot costs ([0-9]+) ore. Each clay robot costs ([0-9]+) ore. Each obsidian robot costs ([0-9]+) ore and ([0-9]+) clay. Each geode robot costs ([0-9]+) ore and ([0-9]+) obsidian.")

	// Construct the blueprints
	bps := functional.Map(lines, func(line string) *blueprint {
		m := r.FindStringSubmatch(line)
		robots := make([]*robot, resourceCount, resourceCount)

		var reqs [][]int
		for i := 0; i < int(resourceCount); i++ {
			reqs = append(reqs, make([]int, resourceCount, resourceCount))
		}

		reqs[ore][ore] = parse.Atoi(m[2])
		reqs[clay][ore] = parse.Atoi(m[3])
		reqs[obsidian][ore] = parse.Atoi(m[4])
		reqs[obsidian][clay] = parse.Atoi(m[5])
		reqs[geode][ore] = parse.Atoi(m[6])
		reqs[geode][obsidian] = parse.Atoi(m[7])

		var maxOre int // Maximum number of ore required for any robot
		for res := 0; res < int(resourceCount); res++ {
			robots[res] = &robot{reqs[res]}
			// Don't need to check the max against ore robot requirements
			if res != int(ore) {
				maxOre = maths.Max(maxOre, reqs[res][ore])
			}
		}
		bp := &blueprint{
			parse.Atoi(m[1]),
			robots,
			maxOre,
		}
		return bp
	})

	// Solve blueprints
	var part1 int
	part2 := 1
	for i, bp := range bps {
		// Part 1: Get quality (index times best)
		part1 += bp.id * d.solveBlueprint(bp, 24)

		// Part 2: Multiplicative product of first three blueprints
		if i < 3 {
			part2 *= d.solveBlueprint(bp, 32)
		}
	}

	o.Stdoutln(part1, part2)
}

func (d *day19) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"33 3472",
			},
		},
		{
			ExpectedOutput: []string{
				"1427 4400",
			},
		},
	}
}
