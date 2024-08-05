package p745

// TODO: Change to 745

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

const mod = 1_000_000_007

func P745() *ecmodels.Problem {
	return ecmodels.IntInputNode(745, func(o command.Output, n int) {
		o.Stdoutln(clever(maths.Pow(10, n)))
	}, []*ecmodels.Execution{
		{
			Args: []string{"1"},
			Want: "24",
		},
		{
			Args: []string{"2"},
			Want: "767",
		},
		{
			Args: []string{"14"},
			Want: "94586478",
		},
	})
}

// Assume sum is the sum of G(n), but where it's when only numbers
// from 1 through k are considered. When considering (k+1)
func clever(n int) int {
	p := generator.Primes()

	sum := n

	mapper := map[int]map[int]int{
		// 1: map[int]int{
		// 	1: 1,
		// },
	}

	for k := 2; k*k <= n; k++ {

		if k%10_000 == 0 {
			fmt.Println(k)
		}

		// 4 => Rm ones
		// 9 => Rm ones
		// 16 => Rm every fourth four
		// 25 => Rm ones
		// 36 => Rm

		// 04 08 12 16 20 24 28 32 36 40 44 48 52 56 60 64 68 72
		//  4  4  4 16  4  4  4 16 36  4  4 16  4  4  4 16  4 36
		//  4  4  4 16  4  4  4 16 36  4  4 16  4  4  4 16  4 36

		// 2 => Rm ones
		// 3 => Rm ones
		// 4 => Rm every other two
		// 5 => Rm ones
		// 6 (2,3) => Rm one third of remaining twos and one half of threes
		// 8 => Rm every other 4
		// 9 => Rm every other 3
		// 10 => Rm one half of fives and one fifth of twos
		// 30 =>

		// 2, 6, 10, 14, 18, 22, 26, 30
		// 34, 38, 42, 46, 50,

		// 60 => 2, 4

		// Whenever we get to a number
		// 1. Add all the counts for that number
		// 2. Subtract all the factor pairs for that number

		// At 4 * 9, subtract 4 values and 9 values
		// At 4 * 9 * 25, we have (4*9, 25) (4, 9*25), (4*25, 9)
		// Already would have subtracted all 4s and 9s (from 4*9)
		// Already would have subtracted all 25s (from 4*25 or 9*25)
		// Need to subtract the following values: (4*9), (4*25), (9*25)
		// So, iterate ver factors that are a square of a prime

		// What about with multiple values:
		// 4 * 4 * 9
		// At 4 * 9, subtract 4 values and 9 values
		// At 4 * 4, subtract 4 values
		// Need to subtract (4*4) values and (4*9) values

		// 36=2*2*3*3

		// The number of times we see k
		t := n / (k * k)

		// Only subtract ones
		if p.Contains(k) {
			sum = (sum + t*(k*k-1)) % mod
			mapper[k] = map[int]int{
				// k * k: 1,
				1: -1,
			}
			continue
		}

		// 144 -> 1
		// 4 -> 4
		// 9 -> 4 + 9 - 1
		// 36 -> 4 + 9 - 1 + (36 - 4 - 9 + 1)
		// 36 -> 36
		//

		// 12=12*12 = 2*2*2*2*3*3*3*3
		// 2*2*3*3 => (2, 223), (3 223)
		// 2*2*3*3 => (22, 33f), (3 223)

		offset := k * k
		pfs := p.PrimeFactors(k)

		if true {
			nm, fm := map[int]int{}, map[int]int{
				// k * k: 1,
			}

			for f := range pfs {
				c := k / f
				for k, v := range mapper[c] {
					nm[k] -= v
				}
			}

			for sq, v := range nm {
				if v > 1 {
					offset += sq
					fm[sq] = 1
				} else if v < -1 {
					offset -= sq
					fm[sq] = -1
				}
			}

			for f := range pfs {
				c := k / f
				fm[c*c]--
				offset -= c * c
			}

			for _, k := range fm {
				if fm[k] == 0 {
					delete(fm, k)
				}
			}

			mapper[k] = fm
			// fmt.Println(mapper, offset)

			// fmt.Println("+++++")
			// for _, i := range []int{1, 2, 3, 4, 6, 8, 12, 24} {
			// 	if v, ok := mapper[i]; ok {
			// 		fmt.Println(i, v)
			// 	}
			// }

		} else {

			// first := true
			// var onesRemoved int
			rmFactorCnt := map[int]int{}
			for f := range pfs {
				c := k / f
				offset -= c * c
				// if !first {
				if p.Contains(c) {
					rmFactorCnt[1]++
					// onesRemoved++
				} else {
					for sf := range p.PrimeFactors(c) {
						rmFactorCnt[sf]++
					}
				}
			}

			// 16*9, 36*4
			// 4*3 6*2

			for rf, cnt := range rmFactorCnt {
				if cnt > 1 {
					offset = offset + rf*rf*(cnt-1)
				}
			}
		}

		/*if onesRemoved > 1 {
			if k == 12 {
				fmt.Println("REM", offset, onesRemoved)
			}
			offset = offset + (onesRemoved - 1)
			if k == 12 {
				fmt.Println("REMF", offset)
			}
		}*/

		sum = (sum + t*(offset)) % mod

		// for _, f := range p.Factors(k) {

		// Subtract the

		// f = 9, p = 4*25
		// We've counted 9 before that we should not have
		// }

		// 4 * 9 * 25
		// 2 * 3 * 5
		// 1, 4, 9, 25, 4*9, 4*25, 9*25
		//
	}
	// var sum int
	// sum
	//
	return sum
}

func brute(n int) int {
	var sum int
	p := generator.Primes()
	for i := 1; i <= n; i++ {
		coef := 1
		for f, cnt := range p.PrimeFactors(i) {
			coef = coef * maths.Pow(f, cnt-(cnt%2))
		}
		sum += coef
	}
	return sum
}
