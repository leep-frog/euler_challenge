package p42

import (
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/ec/p22"
	"github.com/leep-frog/euler_challenge/generator"
)

func P42() *ecmodels.Problem {
	return ecmodels.FileInputNode(42, func(lines []string, o command.Output) {
		words := strings.Split(strings.ReplaceAll(lines[0], "\"", ""), ",")
		var count int
		for _, w := range words {
			if generator.IsTriangular(p22.WordScore(w)) {
				count++
			}
		}
		o.Stdoutln(count)
	}, []*ecmodels.Execution{
		{
			Args: []string{"words.txt"},
			Want: "162",
		},
	})
}
