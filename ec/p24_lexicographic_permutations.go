package eulerchallenge

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P24() *command.Node {
	return command.SerialNodes(
		command.Description("Find the nth permutation of 0 1 2 3 4 5 6 7 8 9"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			n := d.Int(N)

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
			f := maths.FacotiralI(len(vs))
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
		}),
	)
}
