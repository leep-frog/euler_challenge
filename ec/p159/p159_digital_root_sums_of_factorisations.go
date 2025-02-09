package p159

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func singleDigitSum(k int) int {
	v := bread.Sum(maths.Digits(k))
	for ; v >= 10; v = bread.Sum(maths.Digits(v)) {
	}
	return v
}

func P159() *ecmodels.Problem {
	return ecmodels.IntInputNode(159, func(o command.Output, n int) {
		primes := generator.Primes()
		cache := []int{0, 0}

		sum := 0
		for i := 2; i < n; i++ {
			if primes.Contains(i) {
				cache = append(cache, singleDigitSum(i))
			} else {
				best := maths.Largest[int, int]()
				best.Check(singleDigitSum(i))
				for _, f := range primes.Factors(i) {
					if f == 1 || f == i {
						continue
					}
					best.Check(singleDigitSum(f) + cache[i/f])
				}
				cache = append(cache, best.Best())
			}
			sum += cache[i]
		}
		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args: []string{"25"},
			Want: "151",
		},
		{
			Args:     []string{"1000000"},
			Want:     "14489159",
			Estimate: 2.5,
		},
	})
}

func recursive159(n int, primes *generator.Prime, cache map[int]int) int {
	return generator.CompositeCacher(primes, n, cache, func(i int) int {
		v := bread.Sum(maths.Digits(i))
		for ; v >= 10; v = bread.Sum(maths.Digits(v)) {
		}
		return v
	}, func(primeFactor, otherFactor int) int {
		single := bread.Sum(maths.Digits(n))
		fromFactor := primeFactor + recursive159(otherFactor, primes, cache)
		return maths.Max(single, fromFactor)
	})
}
