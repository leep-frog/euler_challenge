package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P75() *problem {
	return intInputNode(75, func(o command.Output, L int) {
		counts := map[int]int{}
		tg := generator.RightTriangleGenerator().Iterator()
		for t := tg.Next(); 2*t.M*t.M+2*t.M <= L; t = tg.Next() {
			p := t.Perimeter()
			for l := p; l <= L; l += p {
				counts[l]++
			}

		}
		var count int
		for _, v := range counts {
			if v == 1 {
				count++
			}
		}
		o.Stdoutln(count)
	})
}
