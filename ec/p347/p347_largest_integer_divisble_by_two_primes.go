package p347

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P347() *ecmodels.Problem {
	return ecmodels.IntInputNode(347, func(o command.Output, n int) {
		o.Stdoutln(s(n))
	}, []*ecmodels.Execution{
		{
			Args: []string{"100"},
			Want: "2262",
		},
		{
			Args:     []string{"10_000_000"},
			Want:     "11109800204052",
			Estimate: 1.5,
		},
	})
}

func s(max int) int {
	p := generator.Primes()

	var sum int
	for i := 0; p.Nth(i)*p.Nth(i+1) <= max; i++ {
		for j := i + 1; p.Nth(i)*p.Nth(j) <= max; j++ {
			sum += m(p.Nth(i), p.Nth(j), max)
		}
	}
	return sum
}

func m(a, b, max int) int {
	if a*b > max {
		return 0
	}

	best := maths.Largest[int, int]()
	for bee := b; a*bee <= max; bee *= b {
		for ab := a * bee; ab <= max; ab *= a {
			best.Check(ab)
		}
	}

	return best.Best()
}
