package p387

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P387() *ecmodels.Problem {
	return ecmodels.IntInputNode(387, func(o command.Output, n int) {
		p := generator.Primes()

		// Create initial Harshad numbers
		var hNums []int
		for i := 1; i < 10; i++ {
			hNums = append(hNums, i)
		}

		var sum int
		for i := 0; i < n; i++ {
			// Determine which Harshad numbers are right truncatable, strong Harshad numbers
			for _, hn := range hNums {
				// Skip if not right truncatable or strong
				if !rtStrongHarshadNumber(hn, p) {
					continue
				}

				// See if any concatenations are prime
				digits := maths.Digits(hn)
				for offset := 1; offset <= 9; offset += 2 {
					opt := maths.FromDigits(append(digits, offset))
					if p.Contains(opt) {
						sum += opt
					}
				}
			}

			// Create next set of rt Harshad numbers
			var nextHNums []int
			for _, hn := range hNums {
				digits := maths.Digits(hn)
				for offset := 0; offset <= 9; offset++ {
					opt := maths.FromDigits(append(digits, offset))
					if harshadNumber(opt) {
						nextHNums = append(nextHNums, opt)
					}
				}
			}
			hNums = nextHNums
		}

		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args: []string{"3"},
			Want: "90619",
		},
		{
			Args:     []string{"13"},
			Want:     "696067597313468",
			Estimate: 3,
		},
	})
}

func rtStrongHarshadNumber(k int, p *generator.Prime) bool {
	return rtHarshadNumber(k) && p.Contains(k/maths.NewInt(k).DigitSum())
}

func harshadNumber(k int) bool {
	sum := maths.NewInt(k).DigitSum()
	return k%sum == 0
}

func rtHarshadNumber(k int) bool {
	if k == 0 {
		return true
	}
	return harshadNumber(k) && rtHarshadNumber(k/10)
}
