package p862

import (
	"time"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/combinatorics"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P862() *ecmodels.Problem {
	return ecmodels.IntInputNode(862, func(o command.Output, n int) {
		s := maths.Zero()
		for zs := 0; zs < n; zs++ {
			s = s.Plus(dp(zs, n-zs, 1, nil))
		}
		o.Stdoutln(s, time.Now())
	}, []*ecmodels.Execution{
		{
			Args: []string{"3"},
			Want: "1701",
		},
		{
			Args: []string{"12"},
			Want: "6111397420935766740",
		},
	})
}

// Count the number
func dp(numZeroes, remDigits, min int, digitCount []int) *maths.Int {
	if len(digitCount) > 9 {
		return maths.Zero()
	}
	if remDigits == 0 {
		return numDigits(numZeroes, digitCount)
	}

	s := maths.Zero()
	for i := min; i <= remDigits; i++ {
		s = s.Plus(dp(numZeroes, remDigits-i, i, append(digitCount, i)))
	}
	return s
}

func numDigits(numZeroes int, digitCount []int) *maths.Int {
	// Determine the number of numbers we can create with the provided digit count (including zeroes)
	permCnt := combinatorics.PermutationFromCount(append(digitCount, numZeroes))

	// Remove numbers with leading zeroes
	if numZeroes > 0 {
		leadingZeroPerms := combinatorics.PermutationFromCount(append(digitCount, numZeroes-1))
		permCnt = permCnt.Minus(leadingZeroPerms)
	}

	// There will permCnt numbers. Consider these numbers from lowest to highest:
	// p_1, p_2, ..., p_permCnt
	// Then, consider the values for T(p_i):
	// T(p_1) = 0, T(p_2) = 1, ..., T(p_permCnt) = permCnt
	// So, the sum of these values is (1 + 2 + ... + permCnt) = permCnt*(permCnt-1)/2
	sumOfTs := permCnt.Times(permCnt.MinusInt(1)).DivInt(2)

	// Next, consider the number of unique set of numbers
	// For example 112, 221, 113, 331, etc. (note we *don't* count permutations here)
	f := maths.Factorial(9).Div(maths.Factorial(9 - len(digitCount)))

	// Divide redundant digits (e.g. if we have two digit counts of 3, then those are interchangeable,
	// so we divide by 3!
	dm := map[int]int{}
	for _, c := range digitCount {
		dm[c]++
	}
	for _, freq := range dm {
		if freq > 1 {
			f = f.Div(maths.Factorial(freq))
		}
	}

	// Finally, return (number of unique set of numbers) * (sum of T(x) values)
	return f.Times(sumOfTs)
}
