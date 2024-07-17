package p686

import (
	"fmt"
	"math"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P686() *ecmodels.Problem {
	return ecmodels.IntInputNode(686, func(o command.Output, n int) {
		o.Stdoutln(jumpAndLog(n))
	}, []*ecmodels.Execution{
		{
			Args: []string{"45"},
			Want: "12710",
		},
		{
			Args: []string{"678910"},
			Want: "193060223",
		},
	})
}

/********************
* Log+Jump Approach *
*********************/

func jumpAndLog(n int) int {
	target := float64(123)

	// Noticed that all jumps are one of the following values
	opts := []float64{
		196,
		289,
		485,
	}

	// Get the first power that works
	power := float64(1)
	for ; !leadsWithLogApproach(power, target); power++ {
	}

	// Start count at one from initialization above
	count := 1
	for count < n {

		// Iterate over the jumps (in ascending order)
		for _, o := range opts {
			if leadsWithLogApproach(power+o, target) {
				count++
				power += o
				break
			}
		}
	}
	return int(power)
}

/***************
* Log Approach *
****************/

func logApproach(n int) int {
	var power, count int
	for count < n {
		power++
		if leadsWithLogApproach(float64(power), 123) {
			if count%100 == 0 {
				fmt.Println(count, power)
			}
			count++
		}

	}
	return power
}

func leadsWithLogApproach(power, target float64) bool {

	lg2 := math.Log10(2)
	lgLowerBound := math.Log10(target)
	lgUpperBound := math.Log10(target + 1)

	// Initial:   2^k = 123[0-9]*
	//            123000. <= 2^k < 124000.
	//            123 * 10^x <= 2^k < 124 * 10^x
	// (base 10)  lg(123) + x <= k*lg(2) < lg(124) + x

	// x <= k*lg(2) - lg(123)
	// Since x is an integer, just take the floor of this to get x
	x := math.Floor(float64(power)*lg2 - lgLowerBound)

	// k*lg(2) - lg(124) < x
	// For some reason '<=' worked (probably due to precision limits)
	return float64(power)*lg2-lgUpperBound <= x
}

/****************
* Jump Approach *
*****************/

type powOfTwo struct {
	pow int
	v   *maths.Int
}

func newPowOfTwo(k, prec int) *powOfTwo {
	return &powOfTwo{k, maths.MustIntFromString(maths.BigPow(2, k).String()[:prec])}
}

func jump(n int) int {
	want := 123
	digitLen := len(maths.Digits(want))

	// First instance is at 90
	power := 90
	prod := maths.BigPow(2, 90)

	// Noticed that all jumps are either 196, 289, or 485
	prec := 15
	opts := []*powOfTwo{
		newPowOfTwo(196, prec),
		newPowOfTwo(289, prec),
		newPowOfTwo(485, prec),
	}

	limit := maths.MustIntFromString("1" + strings.Repeat("0", prec))
	count := 1
	for count < n {

		// Iterate over the jumps (in ascending order)
		for _, o := range opts {
			next := prod.Times(o.v)
			ds := next.Digits()
			head := maths.FromDigits(ds[:maths.Min(len(ds), digitLen)])
			if head == want {
				count++
				prod = next
				power += o.pow
				if count%100 == 0 {
					fmt.Println(count)
				}

				for prod.GT(limit) {
					prod = prod.DivInt(10)
				}
				break
			}
		}
	}
	return power
}
