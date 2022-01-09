package eulerchallenge

import (
	"sort"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P62() *command.Node {
	return command.SerialNodes(
		command.Description("https://projecteuler.net/problem=62"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			n := d.Int(N)
			cubes := generator.PowerGenerator(3)
			permutationCount := map[string]int{}
			lowestPerm := map[string][]*maths.Int{}
			for i := 0; ; i++ {
				cn := cubes.Nth(i)
				parts := strings.Split(cn.String(), "")
				sort.Strings(parts)
				perm := strings.Join(parts, "")
				permutationCount[perm]++
				lowestPerm[perm] = append(lowestPerm[perm], cn)
				if permutationCount[perm] >= n {
					o.Stdoutln(maths.BigMin(lowestPerm[perm]))
					return
				}
			}
		}),
	)
}
