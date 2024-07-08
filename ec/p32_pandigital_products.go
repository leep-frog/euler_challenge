package eulerchallenge

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/combinatorics"
	"github.com/leep-frog/euler_challenge/parse"
)

func P32() *problem {
	return noInputNode(32, func(o command.Output) {
		unique := map[int]bool{}
		for _, perm := range combinatorics.StringPermutations([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}) {
			for i := 1; i < len(perm); i++ {
				for j := i + 1; j < len(perm); j++ {
					a, b, c := parse.Atoi(perm[0:i]), parse.Atoi(perm[i:j]), parse.Atoi(perm[j:])
					if a*b == c {
						unique[c] = true
					}
				}
			}
		}

		var r int
		for c := range unique {
			r += c
		}
		o.Stdoutln(r)
	}, &execution{
		want:     "45228",
		estimate: 0.75,
	})
}
