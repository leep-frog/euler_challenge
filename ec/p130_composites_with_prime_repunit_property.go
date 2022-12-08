package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P130() *problem {
	return intInputNode(130, func(o command.Output, n int) {
		diffs := []int{2, 4, 2, 2}
		g := generator.Primes()

		var count, sum int
		for i, j := 3, 1; ; i, j = i+diffs[j], (j+1)%len(diffs) {
			if g.Contains(i) {
				continue
			}
			k := repunitSmallest(i)

			if (i-1)%k == 0 {
				count++
				sum += i
				if count >= n {
					o.Stdoutln(sum)
					return
				}
			}
		}
	}, []*execution{
		{
			args: []string{"25"},
			want: "149253",
		},
		{
			args: []string{"5"},
			want: "1985",
		},
	})
}
