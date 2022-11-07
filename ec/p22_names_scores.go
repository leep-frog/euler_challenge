package eulerchallenge

import (
	"fmt"
	"sort"
	"strings"

	"github.com/leep-frog/command"
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

func P22() *problem {
	return fileInputNode(22, func(lines []string, o command.Output) {
		namesStr := lines[0]

		names := strings.Split(strings.ReplaceAll(namesStr, `"`, ""), ",")
		sort.Strings(names)

		var score int
		for i, name := range names {
			score += wordScore(name) * (i + 1)
		}
		o.Stdoutln(score)
	}, []*execution{
		{
			args: []string{"p22.txt"},
			want: "871198282",
		},
	})
}
