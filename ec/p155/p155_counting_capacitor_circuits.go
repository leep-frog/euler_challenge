package p155

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/fraction"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

// Find D(18) where D(n) is the number of distinct total capacitance values we
// can obtain when using up to n equal-valued capacitors.

func P155() *ecmodels.Problem {
	return ecmodels.IntInputNode(155, func(o command.Output, n int) {

		// Initially tried with float64, but the precision wasn't good enough.
		// Using fractions ensures that we are perfectly precise.

		// List of circuit values that can be made wih k circuits
		circuitValues := []*maths.Set[*fraction.Fraction]{
			nil,
			maths.NewSet(fraction.New(1, 1)),
		}

		primes := generator.Primes()
		uniqueCs := maths.NewSet(fraction.New(1, 1))
		for i := 2; i <= n; i++ {
			circuitValues = append(circuitValues, maths.NewSet[*fraction.Fraction]())
			for j := 1; j <= i/2; j++ {
				// Assume we have all possible capacitances for D(i-1), D(i-2), etc.
				// Then, to get the capacitance values for D(i), we just combine all
				// values, D(a) and D(b) s.t. a + b = i.
				circuitValues[j].For(func(c1 *fraction.Fraction) bool {
					circuitValues[i-j].For(func(c2 *fraction.Fraction) bool {
						parallel := (c1.Reciprocal().Plus(c2.Reciprocal())).Reciprocal()
						parallel = fraction.Simplify(parallel.N, parallel.D, primes)

						series := c1.Plus(c2)
						series = fraction.Simplify(series.N, series.D, primes)

						circuitValues[i].Add(parallel, series)
						uniqueCs.Add(parallel, series)
						return false
					})
					return false
				})
			}
		}
		o.Stdoutln(uniqueCs.Len())
	}, []*ecmodels.Execution{
		{
			Args: []string{"1"},
			Want: "1",
		},
		{
			Args: []string{"2"},
			Want: "3",
		},
		{
			Args: []string{"3"},
			Want: "7",
		},
		{
			Args: []string{"7"},
			Want: "179",
		},
		{
			Args:     []string{"18"},
			Want:     "3857447",
			Estimate: 150,
		},
	})
}
