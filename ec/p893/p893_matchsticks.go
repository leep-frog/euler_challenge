package p893

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P893() *ecmodels.Problem {
	return ecmodels.IntInputNode(893, func(o command.Output, n int) {

		count := []int{
			//  0, 1, 2, 3, 4, 5, 6, 7, 8, 9
			0 + 6, 2, 5, 5, 4, 5, 6, 3, 7, 6,
		}

		valsAddOnly := []int{
			count[0],
			count[1],
		}
		valsMulOnly := []int{
			count[0],
			count[1],
		}

		p := generator.Primes()
		for len(valsAddOnly) <= n {
			k := len(valsAddOnly)

			// Check the best possible values
			bestMulOnly := maths.Smallest[int, int]()
			bestAddOnly := maths.Smallest[int, int]()

			// Check the digit on its own
			digitsValue := 0
			for _, d := range maths.Digits(k) {
				digitsValue += count[d]
			}
			bestMulOnly.Check(digitsValue)
			bestAddOnly.Check(digitsValue)

			// Multiply it by things
			for _, f := range p.Factors(k) {
				if f == 1 || f == k {
					continue
				}
				bestMulOnly.Check(2 + valsMulOnly[f] + valsMulOnly[k/f])
			}

			// Add to things
			// TODO: Optimize
			for i := 1; i <= k/2; i++ {
				bestAddOnly.Check(2 + valsMulOnly[i] + valsMulOnly[k-i])
				bestAddOnly.Check(2 + valsMulOnly[i] + valsAddOnly[k-i])
				bestAddOnly.Check(2 + valsAddOnly[i] + valsMulOnly[k-i])
				bestAddOnly.Check(2 + valsAddOnly[i] + valsAddOnly[k-i])
			}

			valsAddOnly = append(valsAddOnly, bestAddOnly.Best())
			valsMulOnly = append(valsMulOnly, bestMulOnly.Best())
		}

		var sum int
		for i := 1; i <= n; i++ {
			if valsAddOnly[i] < valsMulOnly[i] {
				sum += valsAddOnly[i]
			} else {
				sum += valsMulOnly[i]
			}
		}

		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args: []string{"100"},
			Want: "916",
		},
		{
			Args:     []string{"1000000"},
			Want:     "26688208",
			Estimate: 60 * 60,
		},
	})
}
