package y2022

import (
	"fmt"
	"regexp"
	"time"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
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
	panic("IDK")
}

type blueprint struct {
	id           int
	robots       []*robot
	maxOreRobots int
}

type robot struct {
	requirements []int
}

func (r *robot) constructable(resources []int) bool {
	for res, required := range r.requirements {
		if required > resources[res] {
			return false
		}
	}
	return true
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

var (
	cc = map[string]bool{}
	// maxMinutes = 24
	maxMinutes = 32
)

func (d *day19) reconstruct(bp *blueprint, path [][]int) {
	idx := 0
	robots := make([]int, resourceCount, resourceCount)
	robots[ore] = 1
	resources := make([]int, resourceCount, resourceCount)

	for min := maxMinutes; min >= 0; min-- {
		fmt.Printf("== Minute %d ==\n", maxMinutes-min+1)

		built := make([]bool, resourceCount, resourceCount)
		if idx < len(path) {
			minute, robot := path[idx][0], path[idx][1]
			if min == minute {
				fmt.Printf("Spend resources to start building a %v-collecting robot\n", resource(robot))
				idx++
				bp.robots[robot].construct(resources)
				robots[robot]++
				built[robot] = true
			}
		}

		for r, cnt := range robots {
			rbt := resource(r)
			if built[r] {
				cnt--
			}
			if cnt > 0 {
				fmt.Printf("%d %v-collecting robot collects %d %v; you now have %d %v\n", cnt, rbt, cnt, rbt, resources[rbt]+cnt, rbt)
				resources[rbt] += cnt
			}
		}
	}
}

func (d *day19) evalBlueprint2(minutes int, bp *blueprint, resources, robots []int, cur [][]int, best *maths.Bester[[][]int, int]) {
	if minutes <= 0 {
		return
	}
	// best.IndexCheck(slices.Clone(cur), resources[geode]+robots[geode]*minutes)
	best.IndexCheck(nil, resources[geode]+robots[geode]*minutes)
	if minutes == 0 {
		return
	}
	// best we can possibly do is construct one per minute
	maxBest := resources[geode] + robots[geode]*minutes
	maxBest += (minutes * (minutes - 1)) / 2
	if best.Best() >= maxBest {
		return
	}
	for res, rbt := range bp.robots {
		if res == int(ore) && robots[ore] >= bp.maxOreRobots {
			continue
		}
		if mins, ok := rbt.until(resources, robots); ok {
			d.elapse(mins+1, resources, robots)
			robots[res]++
			rbt.construct(resources)
			d.evalBlueprint2(minutes-mins-1, bp, resources, robots, append(cur, []int{minutes - mins, res}), best)
			rbt.deconstruct(resources)
			robots[res]--
			d.unelapse(mins+1, resources, robots)
		}
	}
}

// TODO: Change this to branch at each robot we will build next.
// For example, determine the earliest turn at which we can build
// a clay robot and then jump to that turn.
// TODO: Also check best (aka no more buliding) at each step
func (d *day19) evalBlueprint(minutes int, bp *blueprint, resources, robots []int, cur []int, dont []bool) int {
	if minutes <= 0 {
		return resources[geode]
	}

	for rbt, cnt := range robots {
		resources[rbt] += cnt
	}

	var max int
	for res, r := range bp.robots {
		if !dont[res] && r.constructable(resources) {
			r.construct(resources)
			robots[res]++
			max = maths.Max(max, d.evalBlueprint(minutes-1, bp, resources, robots, append(cur, res), make([]bool, resourceCount)))
			r.deconstruct(resources)
			robots[res]--
		} else {
			dont[res] = true
		}
	}

	max = maths.Max(max, d.evalBlueprint(minutes-1, bp, resources, robots, append(cur, -1), dont))

	for rbt, cnt := range robots {
		resources[rbt] -= cnt
	}

	return max
}

func (d *day19) Solve(lines []string, o command.Output) {
	fmt.Println("START")
	r := regexp.MustCompile("^Blueprint ([0-9]+): Each ore robot costs ([0-9]+) ore. Each clay robot costs ([0-9]+) ore. Each obsidian robot costs ([0-9]+) ore and ([0-9]+) clay. Each geode robot costs ([0-9]+) ore and ([0-9]+) obsidian.")

	bps := parse.Map(lines, func(line string) *blueprint {
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

		var maxOre int
		for res := 0; res < int(resourceCount); res++ {
			robots[res] = &robot{reqs[res]}
			// Don't need to check against ore robot
			if res != int(ore) {
				maxOre = maths.Max(maxOre, reqs[res][ore])
			}
		}
		bp := &blueprint{
			parse.Atoi(m[1]),
			robots,
			maxOre,
		}
		fmt.Println(line, bp)
		return bp
	})

	fmt.Println("MID")

	var sum int
	bestBest := 1
	for i, bp := range bps {
		fmt.Println("BP", time.Now())
		resources := make([]int, resourceCount, resourceCount)
		robots := make([]int, resourceCount, resourceCount)
		robots[ore] = 1

		best := maths.Largest[[][]int, int]()
		d.evalBlueprint2(maxMinutes, bp, resources, robots, nil, best)
		fmt.Println("HEYO", bp.id, best)
		sum += bp.id * best.Best()
		fmt.Println(time.Now(), best.Best())
		if i < 3 {
			bestBest *= best.Best()
		}

		// d.reconstruct(bp, best.BestIndex())

		// break
	}

	fmt.Println(sum, bestBest)
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
