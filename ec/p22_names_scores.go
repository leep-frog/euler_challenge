package eulerchallenge

import (
	"sort"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/parse"
)

func P22() *command.Node {
	return command.SerialNodes(
		command.Description("Sorts the names in the provided file and computes the total of the name scores"),
		command.StringNode("FILE", ""),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			namesStr := parse.ReadFileLines(d.String("FILE"))[0]

			names := strings.Split(strings.ReplaceAll(namesStr, `"`, ""), ",")
			sort.Strings(names)

			letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
			letterMap := map[string]int{}
			for i := 0; i < len(letters); i++ {
				letterMap[letters[i:i+1]] = i + 1
			}

			var score int
			for i, name := range names {
				var curScore int
				for j := 0; j < len(name); j++ {
					curScore += letterMap[name[j:j+1]]
				}
				score += curScore * (i + 1)
			}
			o.Stdoutln(score)
		}),
	)
}
