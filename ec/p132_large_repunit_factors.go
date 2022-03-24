package eulerchallenge

import (
	"strconv"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P132() *problem {
	return intInputNode(132, func(o command.Output, n int) {
		rLen := maths.Pow(10, n)
		g := generator.Primes()
		var sum, count int
		prod := 1
		for i, pi := 0, g.Nth(0); count < 40 && len(strconv.Itoa(prod)) < rLen; i, pi = i+1, g.Nth(i+1) {
			if !repunitable(pi) {
				continue
			}
			k := repunitSmallest(pi)
			if rLen%k == 0 {
				sum += pi
				count++
				prod *= pi
			}
		}
		o.Stdoutln(sum)
	})
}
