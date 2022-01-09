package eulerchallenge

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/parse"
)

func P13() *command.Node {
	return command.SerialNodes(
		command.Description("Find a large sum"),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			// TODO: store branch value in data?
			lines := parse.ReadFileLines("p13.txt")

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
			o.Stdout(strings.Join(sumStr, "")[:10])
		}),
	)
}
