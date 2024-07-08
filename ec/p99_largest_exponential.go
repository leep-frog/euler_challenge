package eulerchallenge

import (
	"math"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func P99() *problem {
	return fileInputNode(99, func(lines []string, o command.Output) {
		best := maths.Largest[int, float64]()
		for idx, line := range lines {
			info := strings.Split(line, ",")
			base, exp := parse.Atoi(info[0]), parse.Atoi(info[1])
			result := math.Log10(float64(base)) * float64(exp)
			// Line number isn't 0-indexes, hence the "+1"
			best.IndexCheck(idx+1, result)
		}
		o.Stdoutln(best.BestIndex())
	}, []*execution{
		{
			args: []string{"p099.txt"},
			want: "709",
		},
	})
}
