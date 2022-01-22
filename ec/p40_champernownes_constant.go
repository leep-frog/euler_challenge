package eulerchallenge

import (
	"strconv"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/parse"
)

func P40() *problem {
	return noInputNode(40, func(o command.Output) {
		want := map[int]bool{
			1:       true,
			10:      true,
			100:     true,
			1000:    true,
			10000:   true,
			100000:  true,
			1000000: true,
		}
		index := 0
		prod := 1
		for i := 1; len(want) > 0; i++ {
			for s, j := strconv.Itoa(i), 0; j < len(s); j++ {
				index++
				if want[index] {
					delete(want, index)
					prod *= parse.Atoi(s[j : j+1])
				}
			}
		}
		o.Stdoutln(prod)
	})
}
