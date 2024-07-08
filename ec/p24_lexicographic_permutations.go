package eulerchallenge

import (
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P24() *problem {
	return intInputNode(24, func(o command.Output, n int) {

		vs := []string{
			"0",
			"1",
			"2",
			"3",
			"4",
			"5",
			"6",
			"7",
			"8",
			"9",
		}

		// Since we are sorting, we know that the first 9! values will start with 0,
		// the factorials from (3 * 9! + 2 * 8!) will start with 32, etc.
		digits := []string{}
		index := 0
		f := maths.FactorialI(len(vs))
		for len(vs) > 0 {
			f /= len(vs)

			i := 0
			for ; index < n; index += f {
				i++
			}
			index -= f
			digits = append(digits, vs[i-1])
			vs = append(vs[:i-1], vs[i:]...)
		}

		o.Stdoutln(strings.Join(digits, ""))

		/* Brute force approach
		ps := maths.Permutations(vs)
		sort.Strings(ps)
		o.Stdoutln(ps[n-1])*/
	}, []*execution{
		{
			args: []string{"1000000"},
			want: "2783915460",
		},
		{
			args: []string{maths.Factorial(9).Plus(maths.One()).String()},
			want: "1023456789",
		},
		{
			args: []string{maths.Factorial(9).String()},
			want: "0987654321",
		},
	})
}
