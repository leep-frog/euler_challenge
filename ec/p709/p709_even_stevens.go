package p709

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P709() *ecmodels.Problem {
	return ecmodels.IntInputNode(709, func(o command.Output, n int) {
		o.Stdoutln(cle(n))
	}, []*ecmodels.Execution{
		{
			Args: []string{"8"},
			Want: "1385",
		},
		{
			Args:     []string{"24680"},
			Want:     "773479144",
			Estimate: 20,
		},
	})
}

var (
	cache = map[string]int{}
	mod   = 1_020_202_009
)

// counts is map from bag with number of counts, to number of bags with that
func dp(n int, count int) int {

	if n == 0 {
		return 1
	}

	code := fmt.Sprintf("%d-%d", n, count)
	if v, ok := cache[code]; ok {
		return v
	}

	var sum int

	// Put in an empty bag
	sum = (sum + dp(n-1, count+1)) % mod

	for i := 2; i <= count; i += 2 {
		v := dp(n-1, count-i+1)
		// Number of permutations with those bags
		coef := maths.Choose(count, i).ModInt(mod)

		t := (coef * v) % mod
		sum = (sum + t) % mod
	}

	cache[code] = sum
	return sum
}

var (
	ats = []int{
		1,
		1,
	}

	invs = []int{
		0,
	}
)

// Use this to cache since maths.PowMod is super slow
func getInv(k int) int {
	for len(invs) <= k {
		n := len(invs)
		invs = append(invs, maths.PowMod(n, -1, mod))
	}
	return invs[k]
}

func cle(at int) int {
	// 2*a(n+1) = Sum_{k=0..n} binomial(n, k)*a(k)*a(n-k)
	for len(ats) <= at {
		var sum int
		nPlusOne := len(ats)
		n := nPlusOne - 1

		nChooseK := 1
		for k := 0; k < nPlusOne; k++ {
			v := nChooseK
			v = (v * ats[k]) % mod
			v = (v * ats[n-k]) % mod
			sum = (sum + v) % mod

			// Update nChooseK
			nChooseK = (nChooseK * (n - k)) % mod
			inv := getInv(k + 1)
			nChooseK = (nChooseK * inv) % mod
		}

		inv := getInv(2)
		sum = (sum * inv) % mod
		ats = append(ats, sum)
	}

	return ats[at]
}
