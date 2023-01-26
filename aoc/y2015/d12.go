package y2015

import (
	"encoding/json"
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
)

func Day12() aoc.Day {
	return &day12{}
}

type day12 struct{}

func (d *day12) search(k interface{}, sum *int, part2 bool) {
	switch v := k.(type) {
	case int:
		*sum += v
	case []interface{}:
		for _, inf := range v {
			d.search(inf, sum, part2)
		}
	case map[string]interface{}:
		if part2 {
			for _, val := range v {
				if val == "red" {
					return
				}
			}
		}
		for _, inf := range v {
			d.search(inf, sum, part2)
		}
	case string:
	case float64:
		*sum += int(v)
	case bool:
	default:
		fmt.Printf("UNKONWN TYPE: %t", v)
	}
}

func (d *day12) Solve(lines []string, o command.Output) {
	var result []interface{}
	var sum1, sum2 int
	json.Unmarshal([]byte(lines[0]), &result)
	d.search(result, &sum1, false)
	d.search(result, &sum2, true)

	o.Stdoutln(sum1, sum2)
}

func (d *day12) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"6 4",
			},
		},
		{
			ExpectedOutput: []string{
				"191164 87842",
			},
		},
	}
}
