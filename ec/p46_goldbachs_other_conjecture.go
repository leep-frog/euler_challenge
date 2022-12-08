package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P46() *problem {
	return noInputNode(46, func(o command.Output) {
		primes := generator.Primes()
		for i := 3; ; i += 2 {
			if primes.Contains(i) {
				continue
			}

			//for i, a := 0, primes.Nth(0); a < i; i, a = i+1, primes.Nth(i+1) {
			for iter, a := primes.Start(0); a < i; a = iter.Next() {
				for b := 0; a+2*b*b <= i; b++ {
					if a+2*b*b == i {
						goto NEXT
					}
				}
			}
			o.Stdoutln(i)
			return

		NEXT:
		}
	}, &execution{
		want: "5777",
	})
}
