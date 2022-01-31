package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P92() *problem {
	return noInputNode(92, func(o command.Output) {
		cache := map[int]bool{}
		var count int
		for i := 1; i < 10000000; i++ {
			cur := i
			for cur != 1 && cur != 89 {
				if v, ok := cache[cur]; ok {
					if v {
						cur = 89
					} else {
						cur = 1
					}
					break
				}
				var next int
				for _, d := range maths.Digits(cur) {
					next += d * d
				}
				cur = next
			}
			if cur == 89 {
				count++
			}
			cache[cur] = cur == 89
		}
		o.Stdoutln(count)
	})
}
