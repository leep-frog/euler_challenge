package p62

import (
	"sort"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P62() *ecmodels.Problem {
	return ecmodels.IntInputNode(62, func(o command.Output, n int) {
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
	}, []*ecmodels.Execution{
		{
			Args: []string{"5"},
			Want: "127035954683",
		},
		{
			Args: []string{"3"},
			Want: "41063625",
		},
	})
}
