package eulerchallenge

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P174() *problem {
	return intInputNode(174, func(o command.Output, n int) {
		cnts := make([]int, n+1, n+1)
		squares := generator.SmallPowerGenerator(2)
		for i := 3; squares.Nth(i)-squares.Nth(i-2) < n; i++ {
			sqi := squares.Nth(i)
			for j := i - 2; j > 0 && sqi-squares.Nth(j) < n; j -= 2 {
				cnts[sqi-squares.Nth(j)]++
			}
		}

		var sum int
		for _, cnt := range cnts {
			if 0 < cnt && cnt <= 10 {
				sum++
			}
		}
		o.Stdoutln(sum)
	}, []*execution{
		{
			args: []string{"33"},
			// 8, 12, 16, 20, 24, 28, 32
			want: "7",
		},
		{
			args: []string{"1_000_000"},
			want: "209566",
		},
	})
}
