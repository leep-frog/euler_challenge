package p893

import (
	"fmt"

	"github.com/google/btree"
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

type node struct {
	number int
	value  int
}

func P893() *ecmodels.Problem {
	return ecmodels.IntInputNode(893, func(o command.Output, n int) {

		count := []int{
			//  0, 1, 2, 3, 4, 5, 6, 7, 8, 9
			0 + 6, 2, 5, 5, 4, 5, 6, 3, 7, 6,
		}

		fmt.Println("A")

		valsAddOnly := []int{
			count[0],
			count[1],
		}
		valsMulOnly := []int{
			count[0],
			count[1],
		}
		addTree := btree.NewG[node](2, func(a, b node) bool {
			if a.value == b.value {
				return a.number < b.number
			}
			return a.value < b.value
		})
		addTree.ReplaceOrInsert(node{1, count[1]})

		p := generator.Primes()

		for len(valsAddOnly) <= n {
			k := len(valsAddOnly)

			if k%1_000 == 0 {
				fmt.Println(k)
			}

			// Check the best possible values
			bestMulOnly := maths.Smallest[int, int]()
			bestAddOnly := maths.Smallest[int, int]()

			// Check the digit on its own
			digitsValue := 0
			for _, d := range maths.Digits(k) {
				digitsValue += count[d]
			}
			bestMulOnly.Check(digitsValue)
			bestAddOnly.Check(digitsValue)

			// Multiply it by things
			for _, f := range p.Factors(k) {
				if f == 1 || f == k {
					continue
				}
				bestMulOnly.Check(2 + valsMulOnly[f] + valsMulOnly[k/f])
			}

			// (Brute force): Add to things, but only consider values that are the best
			// for i := 1; i <= k/2; i++ {
			// 	if bestAddOnly.Check(2+valsMulOnly[i]+valsMulOnly[k-i]) || bestAddOnly.Check(2+valsMulOnly[i]+valsAddOnly[k-i]) || bestAddOnly.Check(2+valsAddOnly[i]+valsMulOnly[k-i]) || bestAddOnly.Check(2+valsAddOnly[i]+valsAddOnly[k-i]) {
			// 		m[i]++
			// 		if k == 2980 {
			// 			fmt.Println(k, i, k-i)
			// 		}
			// 	}
			// }

			// Try adding pairs together
			addTree.Ascend(func(item node) bool {
				betterR := maths.Min(valsAddOnly[k-item.number], valsMulOnly[k-item.number])
				bestAddOnly.Check(2 + item.value + betterR)
				// Only consider the smaller of the two values for optimization
				return item.value <= bestAddOnly.Best()/2
			})

			valsAddOnly = append(valsAddOnly, bestAddOnly.Best())
			valsMulOnly = append(valsMulOnly, bestMulOnly.Best())
			addTree.ReplaceOrInsert(node{
				k, maths.Min(bestAddOnly.Best(), bestMulOnly.Best()),
			})
		}

		var sum int
		for i := 1; i <= n; i++ {
			if valsAddOnly[i] < valsMulOnly[i] {
				sum += valsAddOnly[i]
			} else {
				sum += valsMulOnly[i]
			}
		}

		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args: []string{"100"},
			Want: "916",
		},
		{
			Args:     []string{"1000000"},
			Want:     "26688208",
			Estimate: 60 * 60,
		},
	})
}
