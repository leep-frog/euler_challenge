package y2022

import (
	"fmt"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/functional"
	"github.com/leep-frog/euler_challenge/parse"
	"golang.org/x/exp/slices"
)

func Day11() aoc.Day {
	return &day11{}
}

type day11 struct{}

type monkey struct {
	items           []int
	operation       func(int) int
	testToMonkey    func(int) int
	inspectionCount int
	mod             int
}

func (m *monkey) String() string {
	return fmt.Sprintf("%v %d %d", m.items, m.mod, m.inspectionCount)
}

func (d *day11) solve(lines []string, o command.Output, part1 bool, rounds int) int {
	var monkeys []*monkey
	for i := 1; i < len(lines); i += 7 {
		// Items
		items := parse.AtoiArray(strings.Split(strings.Split(lines[i], ":")[1], ","))

		// Operation
		parts := strings.Split(lines[i+1], " ")
		parts = parts[len(parts)-2:]
		opType, v := parts[0], parts[1]
		op := func(k int) int {
			if v == "old" {
				return k + k
			}
			return k + parse.Atoi(v)
		}
		if opType == "*" {
			op = func(k int) int {
				if v == "old" {
					return k * k
				}
				return k * parse.Atoi(v)
			}
		}

		parts = strings.Split(lines[i+2], " ")
		mod := parse.Atoi(parts[len(parts)-1])

		parts = strings.Split(lines[i+3], " ")
		trueMonkey := parse.Atoi(parts[len(parts)-1])
		parts = strings.Split(lines[i+4], " ")
		falseMonkey := parse.Atoi(parts[len(parts)-1])
		test := func(k int) int {
			if k%mod == 0 {
				return trueMonkey
			}
			return falseMonkey
		}

		// fmt.Println(items, op(100), test(13*23))
		monkeys = append(monkeys, &monkey{
			items, op, test, 0, mod,
		})
	}

	superMod := functional.Reduce(1, monkeys, func(b int, m *monkey) int {
		return b * m.mod
	})

	for round := 0; round < rounds; round++ {
		for _, monkey := range monkeys {
			monkey.inspectionCount += len(monkey.items)
			for _, item := range monkey.items {
				if part1 {
					item = monkey.operation(item) / 3
				} else {
					// item = monkey.operation(item) / 3
					item = monkey.operation(item) % superMod
				}
				next := monkeys[monkey.testToMonkey(item)]
				next.items = append(next.items, item)
			}
			monkey.items = nil
		}
	}

	counts := functional.Map(monkeys, func(m *monkey) int { return m.inspectionCount })
	slices.Sort(counts)
	return counts[len(counts)-1] * counts[len(counts)-2]
}

func (d *day11) Solve(lines []string, o command.Output) {
	o.Stdoutln(
		d.solve(lines, o, true, 20),
		d.solve(lines, o, false, 10000),
	)
}

func (d *day11) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"10605 2713310158",
			},
		},
		{
			ExpectedOutput: []string{
				"64032 12729522272",
			},
		},
	}
}
