package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P46() *problem {
	return noInputNode(46, func(o command.Output) {
		primes := generator.Primes()
		for i := 3; ; i += 2 {
			if generator.IsPrime(i, primes) {
				continue
			}

			primes.Reset()
			for a := primes.Next(); a < i; a = primes.Next() {
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
	})
}
