package p22

import (
	"fmt"
	"sort"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
)

var (
	letterMap = map[string]int{}
)

func WordScore(s string) int {
	if len(letterMap) == 0 {
		for i := range ecmodels.Letters {
			letterMap[ecmodels.Letters[i:i+1]] = i + 1
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

func P22() *ecmodels.Problem {
	return ecmodels.FileInputNode(22, func(lines []string, o command.Output) {
		namesStr := lines[0]

		names := strings.Split(strings.ReplaceAll(namesStr, `"`, ""), ",")
		sort.Strings(names)

		var score int
		for i, name := range names {
			score += WordScore(name) * (i + 1)
		}
		o.Stdoutln(score)
	}, []*ecmodels.Execution{
		{
			Args: []string{"p22.txt"},
			Want: "871198282",
		},
	})
}
