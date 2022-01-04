package eulerchallenge

import (
	"fmt"
	"sort"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/parse"
)

var (
	letterMap = map[string]int{}
	letters   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func wordScore(s string) int {
	if len(letterMap) == 0 {
		for i := range letters {
			letterMap[letters[i:i+1]] = i + 1
		}
	}

	var sum int
	for i := range s {
		if v, ok := letterMap[s[i:i+1]]; ok {
			sum += v
		} else {
			panic(fmt.Sprintf("unknown letter: %q", s[i:i+1]))
		}
	}
	return sum
}

func P22() *command.Node {
	return command.SerialNodes(
		command.Description("Sorts the names in the provided file and computes the total of the name scores"),
		command.StringNode("FILE", ""),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			namesStr := parse.ReadFileLines(d.String("FILE"))[0]

			names := strings.Split(strings.ReplaceAll(namesStr, `"`, ""), ",")
			sort.Strings(names)

			var score int
			for i, name := range names {
				score += wordScore(name) * (i + 1)
			}
			o.Stdoutln(score)
		}),
	)
}
