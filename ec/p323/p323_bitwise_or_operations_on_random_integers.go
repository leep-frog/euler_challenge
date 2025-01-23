package p323

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/fraction"
)

func P323() *ecmodels.Problem {
	return ecmodels.IntInputNode(323, func(o command.Output, n int) {
		evs := []*fraction.Rational{
			fraction.NewRational(0, 1),
		}

		for k := 1; k <= n; k++ {
			evs = calcEV(evs, k)
		}

		o.Stdoutf("%0.10f\n", evs[n].Float64())
	}, []*ecmodels.Execution{
		{
			Args: []string{"1"},
			Want: "2.0000000000",
		},
		{
			Args: []string{"2"},
			Want: "2.6666666667",
		},
		{
			Args: []string{"32"},
			Want: "6.3551758451",
		},
	})
}

func calcEV(evs []*fraction.Rational, k int) []*fraction.Rational {
	// Note (P(x flip out of k) = 1/2^k * (k choose x)
	//
	// ev(k) = P(0 flip)*(ev if 0 flip) + P(1 flips)*(ev if 1 flips) + ... + P(k - 1 flips)*(ev if k - 1 flips) + P(all flip)*(ev if all flip)
	//       = 1/2^k*(k choose 0)*(1+ev(k)) + 1/2^k*(k choose 1)*(1+ev(k-1)) + ... + 1/2^k*(k choose k-1)*(1+ev(1)) + 1/2^k*(k choose k)*(1+ev(0))
	//       = 1/2^k * [(k choose 0) * (1 + ev(k)) + (k choose 1) * (1 + ev(k-1)) + ... + (k choose k-1) * (1 + ev(1)) + (k choose k) * (1 + ev(0))]
	//
	//       Separate out 1s and ev(x) values
	//       = 1/2^k * [(k choose 0) + (k choose 1) + ... + (k choose k-1) + (k choose k)] + 1/2^k * [(k choose 0) * ev(k) + (k choose 1) * (ev(k-1)) + ... + (k choose k-1) * (ev(1)) + (k choose k) * (ev(0))]
	//
	//       First term (1/2^k * sum of chooses) is equal to 1
	//       = 1 + 1/2^k * [(k choose 0) * ev(k) + (k choose 1) * (ev(k-1)) + ... + (k choose k-1) * (ev(1)) + (k choose k) * (ev(0))]
	//
	//       All ev(x) for x < k are known, so we can rearrange to solve for ev(k):
	// ev(k) = 1 + {1/2^k * (k choose 0) * ev(k)} + 1/2^k * [(k choose 1) * (ev(k-1)) + ... + (k choose k-1) * (ev(1)) + (k choose k) * (ev(0))]
	//       = 1 + {1/2^k * ev(k)} + 1/2^k * [(k choose 1) * (ev(k-1)) + ... + (k choose k-1) * (ev(1)) + (k choose k) * (ev(0))]
	// ev(k) (1 - 1/2^k) = 1 + 1/2^k * [(k choose 1) * (ev(k-1)) + ... + (k choose k-1) * (ev(1)) + (k choose k) * (ev(0))]

	// ev(1) = 1/2 * 1 + 1/2 * (1 + ev(1))
	// ev(1) = 1/2 + 1/2 + ev(1)/2
	// ev(1) (1 - 1/2) = 1
	// ev(1) / 2 = 1
	// ev(1) = 2

	// E2 = 1/4 * 1 + 1/4 * (1 + E2) + 1/2 * (1 + E1)
	// E2 = 1/4 + 1/4 + E2/4 + 3/2
	// (3/4) *E2 = 1/2 + 3/2
	// (3/4) *E2 = 2
	// (3/4) *E2 = 8/3

	summation := fraction.NewRational(0, 1)
	for i := 1; i <= k; i++ {
		summation = summation.Plus(choose(k, i).Times(evs[k-i]))
	}

	summation = summation.Div(pow(2, k)).Plus(fraction.NewRational(1, 1))

	divisor := fraction.NewRational(1, 1).Minus(fraction.NewRational(1, 1).Div(pow(2, k)))

	return append(evs, summation.Div(divisor))
}

func pow(a, b int) *fraction.Rational {
	r := fraction.NewRational(1, 1)
	for i := 0; i < b; i++ {
		r = r.Times(fraction.NewRational(a, 1))
	}
	return r
}

func choose(n, k int) *fraction.Rational {
	r := fraction.NewRational(1, 1)

	if k > n-k {
		k = n - k
	}

	for i := 1; i <= k; i++ {
		r = r.Times(fraction.NewRational(n-i+1, i))
	}
	return r
}

/*func brute(n int) {
	var sum int
	times := 1_000_000
	for i := 0; i <= times; i++ {

		ys := n
		var cnt int
		for ys > 0 {
			cnt++

			var mins int
			for j := 0; j < ys; j++ {
				if rand.Uint()%2 == 1 {
					mins++
				}
			}
			ys -= mins

		}
		sum += cnt
	}
	fmt.Println("BRUTE", float64(sum)/float64(times))
}*/
