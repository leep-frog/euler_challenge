package p89

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P89() *ecmodels.Problem {
	return ecmodels.FileInputNode(89, func(lines []string, o command.Output) {
		var saved int
		for _, str := range lines {
			saved += len(str) - len(maths.NumeralFromString(str).String())
		}
		o.Stdoutln(saved)
	}, []*ecmodels.Execution{
		{
			Want: "743",
		},
	})
}
