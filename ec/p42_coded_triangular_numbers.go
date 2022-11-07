package eulerchallenge

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P42() *problem {
	return fileInputNode(42, func(lines []string, o command.Output) {
		words := strings.Split(strings.ReplaceAll(lines[0], "\"", ""), ",")
		var count int
		for _, w := range words {
			if generator.IsTriangular(wordScore(w)) {
				count++
			}
		}
		o.Stdoutln(count)
	}, []*execution{
		{
			args: []string{"words.txt"},
			want: "162",
		},
	})
}
