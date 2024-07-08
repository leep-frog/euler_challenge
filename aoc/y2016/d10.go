package y2016

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/euler_challenge/rgx"
	"github.com/leep-frog/euler_challenge/topology"
	"github.com/leep-frog/functional"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func Day10() aoc.Day {
	return &day10{}
}

type day10 struct{}

type robotCtx struct {
	robots map[int]*robot
	output map[int]int
}

func (rc *robotCtx) getRobot(id int) *robot {
	if rc.robots[id] == nil {
		rc.robots[id] = &robot{id: id}
	}
	return rc.robots[id]
}

type robot struct {
	id           int
	chips        []int
	dependencies []int
	// If false, then goes to output bin
	lowRobot  bool
	lowId     int
	highRobot bool
	highId    int
}

func (r *robot) String() string {
	return fmt.Sprintf("%d %v", r.id, r.chips)
}

func (r *robot) Code(*robotCtx) int {
	return r.id
}

func (r *robot) Dependencies(*robotCtx) []int {
	return r.dependencies
}

func (r *robot) Process(ctx *robotCtx) {
	if len(r.chips) != 2 {
		panic(fmt.Sprintf("Expected 2 chips got %v", r.chips))
	}

	slices.Sort(r.chips)
	if r.lowRobot {
		ctx.robots[r.lowId].chips = append(ctx.robots[r.lowId].chips, r.chips[0])
	} else {
		ctx.output[r.lowId] = r.chips[0]
	}

	if r.highRobot {
		ctx.robots[r.highId].chips = append(ctx.robots[r.highId].chips, r.chips[1])
	} else {
		ctx.output[r.highId] = r.chips[1]
	}
}

func (d *day10) Solve(lines []string, o command.Output) {
	ctx := &robotCtx{
		map[int]*robot{},
		map[int]int{},
	}

	r1 := rgx.New("^value ([0-9]+) goes to bot ([0-9]+)$")
	r2 := rgx.New("^bot ([0-9]+) gives low to ([a-z]+) ([0-9]+) and high to ([a-z]+) ([0-9]+)$")
	for _, line := range lines {
		if m, ok := r1.Match(line); ok {
			value, bot := parse.Atoi(m[0]), parse.Atoi(m[1])
			rb := ctx.getRobot(bot)
			rb.chips = append(rb.chips, value)
			continue
		}

		m := r2.MustMatch(line)
		rb := ctx.getRobot(parse.Atoi(m[0]))
		rb.lowRobot = m[1] == "bot"
		rb.lowId = parse.Atoi(m[2])
		rb.highRobot = m[3] == "bot"
		rb.highId = parse.Atoi(m[4])

		if rb.lowRobot {
			lowRb := ctx.getRobot(rb.lowId)
			lowRb.dependencies = append(lowRb.dependencies, rb.id)
		}
		if rb.highRobot {
			highRb := ctx.getRobot(rb.highId)
			highRb.dependencies = append(highRb.dependencies, rb.id)
		}
	}

	topology.Process[*robotCtx, *robot, int](ctx, maps.Values(ctx.robots))

	var part1 int
	want := functional.If(len(lines) < 10, []int{2, 5}, []int{17, 61})
	for _, r := range ctx.robots {
		if slices.Equal(r.chips, want) {
			part1 = r.id
		}
	}
	o.Stdoutln(part1, ctx.output[0]*ctx.output[1]*ctx.output[2])
}

func (d *day10) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"2 30",
			},
		},
		{
			ExpectedOutput: []string{
				"73 3965",
			},
		},
	}
}
