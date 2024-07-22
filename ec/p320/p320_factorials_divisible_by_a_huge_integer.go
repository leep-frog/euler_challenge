package p320

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

const (
	coef       = 1234567890
	mod  int64 = 1_000_000_000_000_000_000
)

// Related to p549
func P320() *ecmodels.Problem {
	return ecmodels.IntInputNode(320, func(o command.Output, n int) {
		// The solution here is to calculate N(i) for i = 9.
		// Then, determine the number of factors that change when going to i+1 and determine
		// if the maximum number for the factorial changes or not.

		p := generator.Primes()

		// Create the initial set of factors
		factorMaxes := map[int]*factorMax{}
		for f, cnt := range p.MutablePrimeFactors(maths.FactorialI(9)) {
			factorMaxes[f] = factorCount(int64(f), int64(cnt))
		}

		// Keep track of the largest factorial required
		prevBest := maths.Largest[any, int64]()
		for _, fm := range factorMaxes {
			prevBest.Check(fm.number)
		}

		var sum int64
		for i := 10; i <= maths.Pow(10, n); i++ {

			// Get the next set of prime factors to consider, and iterate over them
			fs := p.MutablePrimeFactors(i)
			for f, cnt := range fs {

				// Create the new required number of factors
				iCnt := int64(cnt)
				if fm, ok := factorMaxes[f]; ok {
					iCnt += fm.factorCount
				}
				fm := factorCount(int64(f), iCnt)
				factorMaxes[f] = fm

				// Check if the factor max is the biggest number required
				prevBest.Check(fm.number)
			}
			sum = (sum + prevBest.Best()) % mod
		}

		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args: []string{"3"},
			Want: "614538266565663",
		},
		{
			Args:     []string{"6"},
			Want:     "278157919195482643",
			Estimate: 4,
		},
	})
}

type factorMax struct {
	// The number of factors required
	factorCount int64
	// Smallest number that includes the number of factor counts
	number int64
}

func (fm *factorMax) String() string {
	return fmt.Sprintf("(%d: %v)", fm.factorCount, fm.number)
}

// factorCount returns the smallest number (k) such that k! is divisble by (f^noCoefCnt)
func factorCount(f, noCoefCnt int64) *factorMax {
	power := int64(f)
	factorCount := int64(1)
	coefCnt := noCoefCnt * coef

	// Get the biggest power needed
	for factorCount*f+1 < coefCnt {
		power *= int64(f)
		factorCount = factorCount*f + 1
	}

	var number int64
	for tmpCnt := coefCnt; tmpCnt > 0; tmpCnt, factorCount = tmpCnt%factorCount, (factorCount-1)/f {

		// Increment
		times := tmpCnt / factorCount
		number = (number + ((power*times)%mod)%mod)

		// Decrement power
		power /= int64(f)
	}

	return &factorMax{noCoefCnt, number}
}
