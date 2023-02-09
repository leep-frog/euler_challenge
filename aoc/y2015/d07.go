package y2015

import (
	"regexp"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/euler_challenge/topology"
)

func Day07() aoc.Day {
	return &day07{}
}

type expValue struct {
	key string
	v   string
}

func (ev *expValue) Code() string {
	return ev.key
}

func (ev *expValue) Process(g topology.Graphical[int]) int {
	if parse.IsNumberFormat(ev.v) {
		return parse.Atoi(ev.v)
	}
	return g.Get(ev.v)
}

type expNot struct {
	key string
	v   string
}

func (en *expNot) Code() string {
	return en.key
}

func (en *expNot) Process(g topology.Graphical[int]) int {
	var v int
	if parse.IsNumberFormat(en.v) {
		v = parse.Atoi(en.v)
	} else {
		v = g.Get(en.v)
	}
	return 65535 - v
}

type expression struct {
	key string
	a   string
	op  string
	b   string
}

func (e *expression) Code() string {
	return e.key
}

func (e *expression) Process(g topology.Graphical[int]) int {
	var a, b int
	if parse.IsNumberFormat(e.a) {
		a = parse.Atoi(e.a)
	} else {
		a = g.Get(e.a)
	}

	if parse.IsNumberFormat(e.b) {
		b = parse.Atoi(e.b)
	} else {
		b = g.Get(e.b)
	}

	switch e.op {
	case "LSHIFT":
		return (a << b) % 65535
	case "RSHIFT":
		return (a >> b) % 65535
	case "OR":
		return a | b
	case "AND":
		return a & b
	}
	panic("Unknown operation")
}

type day07 struct{}

func (d *day07) Solve(lines []string, o command.Output) {
	part1 := d.solve(lines, 0)
	o.Stdoutln(part1, d.solve(lines, part1))
}

func (d *day07) solve(lines []string, part2 int) int {
	r := regexp.MustCompile(`^(.*) -> (.*)$`)
	var items []topology.Node[int]
	for _, line := range lines {
		m := r.FindStringSubmatch(line)
		dest := m[2]
		op := strings.Split(m[1], " ")

		// Setting variable
		if len(op) == 1 {
			if dest == "b" && part2 != 0 {
				items = append(items, &expValue{dest, parse.Itos(part2)})
			} else {
				items = append(items, &expValue{dest, op[0]})
			}
			continue
		}

		// NOT operation
		if len(op) == 2 {
			items = append(items, &expNot{dest, op[1]})
			continue
		}

		items = append(items, &expression{dest, op[0], op[1], op[2]})
	}
	return topology.NewGraph(items).Get("a")
}

func (d *day07) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"96 96",
			},
		},
		{
			ExpectedOutput: []string{
				"46065 14134",
			},
		},
	}
}
