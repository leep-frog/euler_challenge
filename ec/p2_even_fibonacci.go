package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P2() *problem {
	return intInputNode(2, func(o command.Output, n int) {
		fibs := generator.Fibonaccis()
		var sum int
		for i := fibs.Next(); i < n; i = fibs.Next() {
			if i%2 == 0 {
				sum += i
			}
		}
		o.Stdoutln(sum)
	})
}
