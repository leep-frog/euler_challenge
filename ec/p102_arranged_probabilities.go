package eulerchallenge

import (
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func P102() *problem {
	return fileInputNode(102, func(lines []string, o command.Output) {
		var count int
		for _, line := range lines {
			nums := strings.Split(line, ",")
			var is []int
			for _, n := range nums {
				is = append(is, parse.Atoi(n))
			}

			one := maths.CrossProductSign(is[0], is[1], is[2], is[3])
			two := maths.CrossProductSign(is[2], is[3], is[4], is[5])
			three := maths.CrossProductSign(is[4], is[5], is[0], is[1])
			if one == two && one == three && one != 0 {
				count++
			}
		}
		o.Stdoutln(count)
	}, []*execution{
		{
			args: []string{"p102_triangles.txt"},
			want: "228",
		},
	})
}
