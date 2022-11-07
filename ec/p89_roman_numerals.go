package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P89() *problem {
	return fileInputNode(89, func(lines []string, o command.Output) {
		var saved int
		for _, str := range lines {
			saved += len(str) - len(maths.NumeralFromString(str).String())
		}
		o.Stdoutln(saved)
	}, []*execution{
		{
			args: []string{"p89.txt"},
			want: "743",
		},
	})
}
