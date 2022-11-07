package eulerchallenge

import (
	"strconv"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/parse"
)

func P37() *problem {
	return noInputNode(37, func(o command.Output) {
		p := generator.Primes()

		var count, sum int
		for i := 0; count < 11; i++ {
			pn := p.Nth(i)
			if pn < 10 {
				continue
			}
			pnStr := strconv.Itoa(pn)
			valid := true
			for j := 1; j < len(pnStr); j++ {
				leftTrunc := pnStr[j:]
				rightTrunc := pnStr[:j]
				if !generator.IsPrime(parse.Atoi(leftTrunc), p) || !generator.IsPrime(parse.Atoi(rightTrunc), p) {
					valid = false
					break
				}
			}
			if valid {
				count++
				sum += pn
			}
		}
		o.Stdoutln(sum)
	}, &execution{
		want: "748317",
	})
}
