package eulerchallenge

import (
	"fmt"
	"sort"

	"github.com/leep-frog/command"
)

func P91() *problem {
	return intInputNode(91, func(o command.Output, n int) {
		unique := map[string]bool{}
		for x1 := 0; x1 <= n; x1++ {
			for y1 := 0; y1 <= n; y1++ {
				for x2 := x1; x2 <= n; x2++ {
					for y2 := 0; y2 <= n; y2++ {
						if x1 == x2 && y1 == y2 {
							continue
						}

						if x1 == 0 && y1 == 0 {
							continue
						}
						if x2 == 0 && y2 == 0 {
							continue
						}

						sides := []int{
							x1*x1 + y1*y1,
							x2*x2 + y2*y2,
							(x2-x1)*(x2-x1) + (y2-y1)*(y2-y1),
						}
						sort.Ints(sides)
						if sides[0]+sides[1] == sides[2] {
							if x1 < x2 || (x1 == x2 && y1 < y2) {
								unique[fmt.Sprintf("(%d, %d) (%d, %d)", x1, y1, x2, y2)] = true
							} else {
								unique[fmt.Sprintf("(%d, %d) (%d, %d)", x2, y2, x1, y1)] = true
							}
						}
					}
				}
			}
		}
		o.Stdoutln(len(unique))
	}, []*execution{
		{
			args:     []string{"50"},
			want:     "14234",
			estimate: 0.5,
		},
		{
			args: []string{"2"},
			want: "14",
		},
	})
}
