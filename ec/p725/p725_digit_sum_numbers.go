package p725

import (
	"fmt"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/combinatorics"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/functional"
	"golang.org/x/exp/slices"
)

func P725() *ecmodels.Problem {
	return ecmodels.IntInputNode(725, func(o command.Output, n int) {
		o.Stdoutln(s(n))
	}, []*ecmodels.Execution{
		{
			Args: []string{"3"},
			Want: "63270",
		},
		{
			Args: []string{"7"},
			Want: "85499991450",
		},
		{
			Args: []string{"2020"},
			Want: "4598797036650685",
		},
	})
}

var (
	mod = maths.Pow(10, 16)
)

// I realized that each number will show up the same number of times in each spot.
// So, the solution is to count the number of times each number shows up in a given spot,
// and then multiply that number by ddddd... for digit d
func s(n int) int {

	// m[d] is the number of times digit d shows up
	var m []int
	for len(m) <= 9 {
		m = append(m, 0)
	}

	// Iterate over each digit-sum number (e.g. 9 in the number 945) and update m
	counts := make([]int, 10, 10)
	for i := 1; i <= 9; i++ {
		counts[i]++
		got := dp(n, n-1, i, 0, counts)
		for i, v := range got {
			m[i] = (m[i] + v) % mod
		}
		counts[i]--
	}

	// Now, add up all the results
	res := maths.Zero()
	for k, v := range m {
		num := maths.MustIntFromString(strings.Repeat(fmt.Sprintf("%d", k), n))
		res = res.Plus(num.TimesInt(v))
	}

	return res.ModInt(mod)
}

func dp(n, remDigits, remValue, min int, counts []int) []int {
	if remDigits == 0 {
		if remValue != 0 {
			return nil
		}

		// Determine the number of combinations it makes.
		combos := combinatorics.PermutationFromCount(bread.Copy(counts))

		// Create the result based on the existing counts
		return functional.Map(counts, func(i int) int {
			// We need to divide by n because a number is in each spot an equal number of times
			// Consider 123, 132, 213, 312, 231, 321, Count of 1 is 1, combos is 6, but it's in each spot 1*6/2=3 times
			// Consider 448, 484, 844. Count of 4 is 2, combos is 3, but it's in each spot 2*3/3 = 2 times
			// We also need to divide here because dividing above isn't always correct, when combos % n != 0
			return combos.TimesInt(i).DivInt(n).ModInt(mod)
		})
	}

	// Create the new sums
	m := []int{}
	for len(m) <= 9 {
		m = append(m, 0)
	}
	for i := min; i <= remValue && i <= 9; i++ {
		for cnt := 1; i*cnt <= remValue && cnt <= remDigits; cnt++ {
			counts[i] += cnt
			for k, v := range dp(n, remDigits-cnt, remValue-(i*cnt), i+1, counts) {
				m[k] = (m[k] + v) % mod
			}
			counts[i] -= cnt
		}
	}
	return m
}

func brute(n int) int {
	p := maths.Pow(10, n)
	var sum int
	for i := 1; i <= p; i++ {
		ds := maths.Digits(i)
		slices.Sort(ds)
		first := ds[len(ds)-1]
		if bread.Sum(ds) == 2*first {
			sum += i
		}
	}
	return sum
}
