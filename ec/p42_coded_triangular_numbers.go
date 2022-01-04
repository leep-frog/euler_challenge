package eulerchallenge

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/parse"
)

func P42() *command.Node {
	return command.SerialNodes(
		command.Description("https://projecteuler.net/problem=42"),
		command.StringNode("FILE", ""),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			words := strings.Split(strings.ReplaceAll(parse.ReadFileLines(d.String("FILE"))[0], "\"", ""), ",")
			var count int
			for _, w := range words {
				if generator.IsTriangular(wordScore(w)) {
					count++
				}
			}
			o.Stdoutln(count)
		}),
	)
}
