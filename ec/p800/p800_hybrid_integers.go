package p800

import (
	"math"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
)

func P800() *ecmodels.Problem {
	return ecmodels.NoInputWithExampleNode(800, func(o command.Output, ex bool) {
		pr := generator.Primes()

		// Evaluate in logs to eliminate need for exponentiation time
		max := 800800 * math.Log10(800800)
		if ex {
			max = 800 * math.Log10(800)
		}

		var sum int
		for i := 0; ; i++ {
			p := float64(pr.Nth(i))

			var count int
			for j := i + 1; ; j++ {
				q := float64(pr.Nth(j))
				if (q*math.Log10(p) + p*math.Log10(q)) > max {
					break
				}
				count++
			}
			if count == 0 {
				break
			}
			sum += count
		}
		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args: []string{"-x"},
			Want: "10790",
		},
		{
			Want:     "1412403576",
			Estimate: 30,
		},
	})
}
