package p13

import (
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/parse"
)

func P13() *ecmodels.Problem {
	return ecmodels.FileInputNode(13, func(lines []string, o command.Output) {
		sum := []int{}
		curSum := 0
		for j := len(lines[0]) - 1; j >= 0; j-- {
			for i := 0; i < len(lines); i++ {
				curSum += parse.Atoi(lines[i][j : j+1])
			}
			sum = append(sum, curSum%10)
			curSum = curSum / 10
		}

		sumStr := []string{}
		sum = append(sum, curSum)
		for i := len(sum) - 1; i >= 0; i-- {
			sumStr = append(sumStr, parse.Itos(sum[i]))
		}
		o.Stdoutln(strings.Join(sumStr, "")[:10])
	}, []*ecmodels.Execution{
		{
			Want: "5537376230",
		},
	})
}
