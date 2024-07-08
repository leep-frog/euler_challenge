package eulerchallenge

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P139() *problem {
	return intInputNode(139, func(o command.Output, n int) {
		var count int
		tg := generator.RightTriangleGenerator().Iterator()
		for t := tg.Next(); t.GuaranteedMinimumPerimeter() < 100_000_000; t = tg.Next() {
			if t.C%(t.B-t.A) == 0 {
				// Only checking uniquely proportioned triangles.
				// Add duplicates to count
				count += 100_000_000 / t.Perimeter()
			}
		}
		o.Stdoutln(count)
	}, []*execution{
		{
			args:     []string{"1"},
			want:     "10057761",
			estimate: 20,
		},
	})
}
