package eulerchallenge

import (
	"math/big"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P143() *problem {
	return intInputNode(143, func(o command.Output, n int) {
		// According to https://en.wikipedia.org/wiki/Fermat_point
		// AN, BM, and CO all intersect at 60 degrees.
		// Therefore, we can find all valid ATB triangles by iterating over
		// side lengths p and q (and calculating A given p and q have a 120 degree
		// angle between them)
		// Then, we just need to find three ATB triangles that create a cycle of
		// side lengths.
		triMap := map[int]map[int]bool{}

		for p := 1; p <= n; p++ {
			for q := p; q+p <= n; q++ {
				// x^2 = q^2 + p^2 + 2pq*cos(120)
				// cos(120) = -1/2
				x2 := q*q + p*p + p*q
				if maths.IsSquare(x2) {
					maths.Insert(triMap, q, p, true)
					maths.Insert(triMap, p, q, true)
				}
			}
		}

		contains := map[int]bool{}
		for p, m := range triMap {
			for q := range m {
				for r := range m {
					if triMap[q][r] && p+q+r <= n {
						contains[p+q+r] = true
					}
				}
			}
		}
		var sum int
		for t := range contains {
			sum += t
		}
		o.Stdoutln(sum)
	})
}

func ratAdd(a, b *big.Rat) *big.Rat {
	return big.NewRat(1, 1).Add(a, b)
}

func ratSub(a, b *big.Rat) *big.Rat {
	return big.NewRat(1, 1).Sub(a, b)
}

func ratMul(a, b *big.Rat) *big.Rat {
	return big.NewRat(1, 1).Mul(a, b)
}

func ratSqrt(r *big.Rat) (*big.Rat, bool) {
	numRoot, nok := bigSqrt(r.Num())
	denRoot, dok := bigSqrt(r.Denom())
	if !nok || !dok {
		return nil, false
	}
	return big.NewRat(1, 1).SetFrac(numRoot, denRoot), true
}

func bigSqrt(i *big.Int) (*big.Int, bool) {
	r := big.NewInt(1).Sqrt(i)
	return r, big.NewInt(1).Mul(r, r).Cmp(i) == 0
}
