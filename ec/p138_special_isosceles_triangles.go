package eulerchallenge

import (
	"math/big"

	"github.com/leep-frog/command"
)

func P138() *problem {
	return intInputNode(138, func(o command.Output, n int) {
		sizes := [][]*big.Int{
			{big.NewInt(8), big.NewInt(17)},
			{big.NewInt(136), big.NewInt(305)},
		}

		/*
			Calculated first few by brute force (just iterating over A = 1, 2, 3, ...
			where B = A +- 1). Noticed pattern:
				Lines represent A and C of a right triangle
				8 17
				136 305 (136 = 8 * 17)
				2448 5473 (2448 = (8 + 136) * 17)
				43920 98209 (43920 = (8 + 136) * 305)
				788120 1762289 (788120 = (136 + 2448) * 305)
				14142232 31622993 (14142232 = (136 + 2448) * 5473)
		*/
		for idx := 0; len(sizes) < n; idx++ {
			nextPos := big.NewInt(1).Mul(big.NewInt(0).Add(sizes[idx][0], sizes[idx+1][0]), sizes[idx+1][1])
			nextNeg := big.NewInt(1).Mul(big.NewInt(0).Add(sizes[idx][0], sizes[idx+1][0]), sizes[idx][1])

			first := big.NewInt(1).Mul(big.NewInt(5), big.NewInt(1).Mul(nextPos, nextPos))
			second := big.NewInt(1).Mul(big.NewInt(4), nextPos)
			third := big.NewInt(1)
			pos := big.NewInt(1).Add(big.NewInt(1).Add(first, second), third)

			first = big.NewInt(1).Mul(big.NewInt(5), big.NewInt(1).Mul(nextNeg, nextNeg))
			second = big.NewInt(1).Mul(big.NewInt(4), nextNeg)
			third = big.NewInt(1)
			neg := big.NewInt(1).Add(big.NewInt(1).Sub(first, second), third)

			posRoot := big.NewInt(1).Sqrt(pos)
			negRoot := big.NewInt(1).Sqrt(neg)

			sizes = append(sizes, [][]*big.Int{
				{nextNeg, negRoot},
				{nextPos, posRoot},
			}...)
		}
		sum := big.NewInt(0)
		for _, s := range sizes {
			sum = big.NewInt(0).Add(sum, s[1])
		}
		o.Stdoutln(sum)
	})
}
