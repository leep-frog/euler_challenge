package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/parse"
)

func P52() *problem {
	return intInputNode(52, func(o command.Output, n int) {
		start := "1"
		end := "1"
		for {
			start += "0"
			end += "6"
			sn := parse.Atoi(start)
			en := parse.Atoi(end)
			for i := sn + 1; i <= en; i++ {
				allSame := true
				for j := 2; j <= n; j++ {
					if !sameDigits(i, i*j) {
						allSame = false
						break
					}
				}
				if allSame {
					o.Stdoutln(i)
					return
				}
			}
		}
	})
}
