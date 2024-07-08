package eulerchallenge

import (
	"github.com/leep-frog/command/command"
)

func P6() *problem {
	return intInputNode(6, func(o command.Output, n int) {
		var squareSums, sumSquares int
		for i := 1; i <= n; i++ {
			sumSquares += i * i
			squareSums += i
		}
		squareSums *= squareSums
		o.Stdoutln(squareSums - sumSquares)
	}, []*execution{
		{
			args: []string{"10"},
			want: "2640",
		},
		{
			args: []string{"100"},
			want: "25164150",
		},
	})
}
