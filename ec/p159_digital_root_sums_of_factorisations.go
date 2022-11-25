package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func singleDigitSum(k int) int {
	v := maths.SumSys(maths.Digits(k)...)
	for ; v >= 10; v = maths.SumSys(maths.Digits(v)...) {
	}
	return v
}

func P159() *problem {
	return intInputNode(159, func(o command.Output, n int) {
		primes := generator.Primes()
		cache := []int{0, 0}

		sum := 0
		for i := 2; i < n; i++ {
			if generator.IsPrime(i, primes) {
				cache = append(cache, singleDigitSum(i))
			} else {
				best := maths.Largest[int, int]()
				best.Check(singleDigitSum(i))
				for _, f := range generator.Factors(i, primes) {
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
	}, []*execution{
		{
			args: []string{"25"},
			want: "151",
		},
		{
			args:     []string{"1000000"},
			want:     "14489159",
			estimate: 2.5,
		},
	})
}

func recursive159(n int, primes *generator.Generator[int], cache map[int]int) int {
	return generator.CompositeCacher(n, primes, cache, func(i int) int {
		v := maths.SumSys(maths.Digits(i)...)
		for ; v >= 10; v = maths.SumSys(maths.Digits(v)...) {
		}
		return v
	}, func(primeFactor, otherFactor int) int {
		single := maths.SumSys(maths.Digits(n)...)
		fromFactor := primeFactor + recursive159(otherFactor, primes, cache)
		return maths.Max(single, fromFactor)
	})
}
