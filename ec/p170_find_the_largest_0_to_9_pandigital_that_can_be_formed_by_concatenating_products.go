package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func containsAll170(concatenatedProduct []int) bool {
	m := make([]int, 10, 10)
	for _, d := range concatenatedProduct {
		m[d]++
	}
	for _, v := range m {
		if v != 1 {
			return false
		}
	}
	return true
}

func P170() *problem {
	return noInputNode(170, func(o command.Output) {
		options := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

		primes := generator.Primes()

		// Exclusively look for solutions that start with the best possible concatenatedProduct in order:
		// 9876543210
		// 987654321.
		// 98765432..
		// 9876543...
		// 987654....
		// etc.
		for k := len(options) - 2; k > 0; k-- {
			best := maths.Largest[int, int]()

			// The permutations of the remaining digits (dots above) will all be checked
			for _, rightDigits := range maths.Permutations(options[k:]) {
				concatenatedProduct := append(maths.CopySlice(options[:k]), rightDigits...)

				// Now try splitting the concatenated product at all places
				for split := len(options) - 2; split > 0; split-- {
					if concatenatedProduct[split] == 0 {
						continue
					}
					leftProduct, rightProduct := maths.FromDigits(concatenatedProduct[:split]), maths.FromDigits(concatenatedProduct[split:])

					// Look for factors in common
					leftFactors := primes.Factors(leftProduct)
					rightFactors := maths.NewSimpleSet(primes.Factors(rightProduct)...)
					for _, lf := range leftFactors {
						// If common factor, check if the divisors and factor make a pandigital.
						if rightFactors[lf] {
							a, b, c := lf, leftProduct/lf, rightProduct/lf
							combinedDigits := append(append(maths.Digits(a), maths.Digits(b)...), maths.Digits(c)...)
							if containsAll170(combinedDigits) {
								best.Check(maths.FromDigits(concatenatedProduct))
							}
						}
					}
				}
			}
			if best.Set() {
				o.Stdoutln(best.Best())
				return
			}
		}
	}, &execution{
		want:     "9857164023",
		estimate: 1.25,
	})
}
