package eulerchallenge

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P16() *problem {
	return intInputNode(16, func(o command.Output, ni int) {
		n := maths.NewInt(ni)

		two := maths.NewInt(2)
		pow := maths.NewInt(1)
		for n.GT(maths.Zero()) {
			pow = pow.Times(two)
			n.MM()
		}

		o.Stdoutln(pow.DigitSum())
	}, []*execution{
		{
			args: []string{"10"},
			want: "7",
		},
		{
			args: []string{"1000"},
			want: "1366",
		},
	})
}
