package eulerchallenge

import (
	"fmt"
	"sort"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P86() *problem {
	return intInputNode(86, func(o command.Output, n int) {
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
	}, []*execution{
		{
			args:     []string{"1000000"},
			want:     "1818",
			estimate: 2.5,
		},
		{
			args: []string{"2000"},
			want: "100",
		},
	})
}
