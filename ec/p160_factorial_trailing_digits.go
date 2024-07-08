package eulerchallenge

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func slow160(n int) int {
	fives := 0
	dd := 1
	for i := 1; i < n; i++ {
		k := i

		// Remove 0s
		for ; k%10 == 0; k /= 10 {
		}

		// Increment 5s
		for ; k%5 == 0; fives, k = fives+1, k/5 {
		}

		// Cancel out as many 5s with as many 2s as we can/as necessary.
		for ; fives > 0 && k%2 == 0; fives, k = fives-1, k/2 {
		}

		dd = (dd * k) % 100_000
	}
	return dd
}

func fast160(tenPow int) int {
	// With some input from the problem thread, I noticed the
	// following relationship:
	// f(n * 5^k) == f(n) * 2^k, up to some max values for k
	// For 1_000_000_000, it works for up to k == 5.
	// Given that, I assumed that k could be (log_10(n) - 4) = (tenPow - 4)
	n := maths.Pow(10, tenPow)
	k := tenPow - 4
	r := slow160(n/maths.Pow(5, k)) * maths.Pow(2, k)
	return r % 100_000
}

func P160() *problem {
	return intInputNode(160, func(o command.Output, n int) {
		o.Stdoutln(fast160(n))
	}, []*execution{
		{
			args: []string{"9"},
			want: "38144",
		},
		{
			args: []string{"12"},
			want: "16576",
		},
	})
}
