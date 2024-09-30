package p622

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P622() *ecmodels.Problem {
	return ecmodels.IntInputNode(622, func(o command.Output, n int) {
		o.Stdoutln(clever(n))
	}, []*ecmodels.Execution{
		{
			Args: []string{"8"},
			Want: "412",
		},
		{
			Args: []string{"60"},
			Want: "3010983666182123972",
		},
	})
}

// After running brute approach, realized that each n requires the following number of shuffles:
// https://oeis.org/A002326
//
// Therefore, we iterate over factors of (2^k)-1 and include values
// that don't divide any smaller number of the form (2^x)-1
func clever(k int) int {

	// Generate factors
	p := generator.Primes()
	max := maths.Pow(2, k)
	fs := p.Factors(max - 1)

	var sum int
	for _, f := range fs {

		// Ignore factors that divide a smaller (2^x)-1
		for two := 2; two < max; two *= 2 {
			if (two-1)%f == 0 {
				goto INVALID
			}
		}
		sum += f + 1
	INVALID:
	}

	return sum
}

// Simply iterate over deck sizes and shuffle each deck until it is the same as it started,
// summing the sizes that need n shuffles.
func brute(n int) int {
	var sum int
	final := maths.Pow(2, n)
	for k := 2; k <= final; k += 2 {
		var deck []int
		for i := 0; i < k; i++ {
			deck = append(deck, i)
		}

		if shufflesRequired(deck) == n {
			sum += k
		}
	}
	return sum
}

func shuffle(deck []int) []int {
	size := len(deck)
	var shuffled []int

	for i := 0; i < size; i++ {
		if i%2 == 0 {
			shuffled = append(shuffled, deck[i/2])
		} else {
			shuffled = append(shuffled, deck[(size/2)+(i/2)])
		}
	}
	return shuffled
}

func ordered(deck []int) bool {
	for i, v := range deck {
		if i != v {
			return false
		}
	}
	return true
}

func shufflesRequired(deck []int) int {
	cnt := 1
	for deck = shuffle(deck); !ordered(deck); deck = shuffle(deck) {
		cnt++
	}
	return cnt
}
