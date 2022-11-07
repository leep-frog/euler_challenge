package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P20() *problem {
	return intInputNode(20, func(o command.Output, n int) {
		o.Stdoutln(maths.Factorial(n).DigitSum())
	}, []*execution{
		{
			args: []string{"100"},
			want: "648",
		},
		{
			args: []string{"10"},
			want: "27",
		},
	})
}
