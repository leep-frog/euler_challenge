package p86

import (
	"fmt"
	"sort"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P86() *ecmodels.Problem {
	return ecmodels.IntInputNode(86, func(o command.Output, n int) {
		unique := map[string]bool{}
		for a := 1; ; a++ {
			if len(unique) >= n {
				o.Stdoutln(a - 1)
				return
			}

			for bL := 1; bL <= a; bL++ {
				for bR := bL; bR <= a; bR++ {
					b := bL + bR
					c2 := a*a + b*b
					if !maths.IsSquare(c2) {
						continue
					}

					cL2 := (a+bL)*(a+bL) + bR*bR
					cR2 := (a+bR)*(a+bR) + bL*bL
					// Verify c2 is the shortest route
					if c2 > cL2 || c2 > cR2 {
						continue
					}

					paths := []int{a, bL, bR}
					sort.Ints(paths)
					pStr := fmt.Sprintf("%d_%d_%d", paths[0], paths[1], paths[2])
					if unique[pStr] {
						continue
					}
					unique[pStr] = true
				}
			}
		}
	}, []*ecmodels.Execution{
		{
			Args:     []string{"1000000"},
			Want:     "1818",
			Estimate: 2.5,
		},
		{
			Args: []string{"2000"},
			Want: "100",
		},
	})
}
