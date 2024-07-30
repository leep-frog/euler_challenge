package p869

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P869() *ecmodels.Problem {
	return ecmodels.IntInputNode(869, func(o command.Output, n int) {

		p := generator.Primes()

		var primes []int
		for i := 0; p.Nth(i) <= n; i++ {
			primes = append(primes, p.Nth(i))
		}

		o.Stdoutf("%0.8f\n", ev(primes))
	}, []*ecmodels.Execution{
		{
			Args: []string{"10"},
			Want: "2.00000000",
		},
		{
			Args: []string{"30"},
			Want: "2.90000000",
		},
		{
			Args:     []string{"100000000"},
			Want:     "14.97696693",
			Estimate: 30,
		},
	})
}

func ev(opts []int) float64 {
	if len(opts) == 0 {
		return 0
	}

	var zeroOpts, oneOpts []int
	for _, o := range opts {
		if o < 1 {

		} else if o%2 == 1 {
			oneOpts = append(oneOpts, o/2)
		} else {
			zeroOpts = append(zeroOpts, o/2)
		}
	}

	bestGuess := float64(maths.Max(len(oneOpts), len(zeroOpts)))
	return (bestGuess + (float64(len(zeroOpts)) * ev(zeroOpts)) + (float64(len(oneOpts)) * ev(oneOpts))) / float64(len(opts))
}
