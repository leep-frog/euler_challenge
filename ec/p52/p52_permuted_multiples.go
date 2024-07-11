package p52

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/ec/p49"
	"github.com/leep-frog/euler_challenge/parse"
)

func P52() *ecmodels.Problem {
	return ecmodels.IntInputNode(52, func(o command.Output, n int) {
		start := "1"
		end := "1"
		for {
			start += "0"
			end += "6"
			sn := parse.Atoi(start)
			en := parse.Atoi(end)
			for i := sn + 1; i <= en; i++ {
				allSame := true
				for j := 2; j <= n; j++ {
					if !p49.SameDigits(i, i*j) {
						allSame = false
						break
					}
				}
				if allSame {
					o.Stdoutln(i)
					return
				}
			}
		}
	}, []*ecmodels.Execution{
		{
			Args: []string{"6"},
			Want: "142857",
		},
		{
			Args: []string{"2"},
			Want: "125874",
		},
	})
}
