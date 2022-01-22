package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P12() *problem {
	return intInputNode(12, func(o command.Output, n int) {
		ts := generator.Triangulars()
		for i := ts.Next(); ; i = ts.Next() {
			var count int
			max := i / 2
			for j := 1; j < max; j++ {
				if i%j == 0 {
					max = (i / j) - 1
					if j*j == i {
						count++
					} else {
						count += 2
					}
					if count > n {
						o.Stdoutln(i)
						return
					}
				}
			}
		}
	})
}
