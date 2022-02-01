package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P94() *problem {
	return noInputNode(94, func(o command.Output) {
		perimSum := 0
		for k := 2; ; k++ {
			for _, j := range []int{k - 1, k + 1} {
				fourH2 := 4*k*k - j*j
				if !maths.IsSquare(fourH2) {
					continue
				}
				twoH := maths.Sqrt(fourH2)
				if twoH*twoH != fourH2 || twoH <= 0 || fourH2 <= 0 {
					panic("oh")
				}

				fourArea := twoH * j
				perim := k*2 + j
				if fourArea%4 == 0 && perim <= 1_000_000_000 {
					perimSum += perim
				}

				if perim > 1_000_000_010 {
					o.Stdoutln(perimSum)
					return
				}
			}
		}
	})
}
