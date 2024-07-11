package p140

import (
	"math/big"
	"sort"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
)

func P140() *ecmodels.Problem {
	return ecmodels.IntInputNode(140, func(o command.Output, n int) {
		// After brute force trial and error, realized that fractions for golden
		// nuggets relate to the following fibonacci sequences:
		// odds: 1, 1, 2, 3, 5, 8, 13, ...
		// evens: 2, 5, 7, 12, 19, 31, ...
		// Odds solutions: 1/2, 3/5, 8/13, ...
		// Evens solutions: 2/5, 7/12, 19/31, ...
		odds := generator.Fibonaccis()
		evens := generator.CustomFibonacci(2, 5)
		var solns []*big.Int
		// Add extra solutions because integer solution isn't ordered with fractions
		for count := 0; count < n+2; count++ {
			// f(x) = x + 3x^2 + xf(x) + x^2f(x)
			// Solve for f(x)
			// f(x)(1 - x - x^2) = x + 3x^2
			// f(x) = (x + 3x^2)/(1 - x - x^2)
			// f(x) = x(1 + 3x)/(1 - x - x^2)

			// Let x = N/D
			// f(x) = [N/D + 3(N/D)^2] / [1 - N/D - (N/D)^2]
			//      = [(ND + 3N^2) / D^2] / [(D^2 - ND - N^2)/D^2]
			//      = (ND + 3N^2)  / (D^2 - ND - N^2)
			var N, D *big.Int
			if count%2 == 0 {
				N, D = big.NewInt(int64(odds.Nth(count+1))), big.NewInt(int64(odds.Nth(count+2)))
			} else {
				N, D = big.NewInt(int64(evens.Nth(count-1))), big.NewInt(int64(evens.Nth(count)))
			}

			NN := big.NewInt(1).Mul(N, N)
			ND := big.NewInt(1).Mul(N, D)
			DD := big.NewInt(1).Mul(D, D)
			top := big.NewInt(1).Add(ND, big.NewInt(1).Mul(big.NewInt(3), NN))
			bottom := big.NewInt(1).Sub(DD, big.NewInt(1).Add(ND, NN))
			solns = append(solns, top.Div(top, bottom))
		}
		sort.SliceStable(solns, func(i, j int) bool { return solns[i].Cmp(solns[j]) < 0 })
		sum := big.NewInt(0)
		for i := 0; i < n; i++ {
			sum.Add(sum, solns[i])
		}
		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args: []string{"30"},
			Want: "5673835352990",
		},
	})
}

func rationalQuadratic(a, b, c int) (*big.Rat, bool) {
	ba, bb, bc := big.NewInt(int64(a)), big.NewInt(int64(b)), big.NewInt(int64(c))
	detOne := big.NewInt(1).Mul(bb, bb)
	detTwo := big.NewInt(1).Mul(big.NewInt(-4), big.NewInt(1).Mul(ba, bc))
	det := big.NewInt(1).Add(detOne, detTwo)
	if det.Cmp(big.NewInt(0)) < 0 {
		return nil, false
	}
	root := big.NewInt(1).Sqrt(det)
	if big.NewInt(1).Mul(root, root).Cmp(det) != 0 {
		return nil, false
	}
	xNum := big.NewInt(1).Sub(root, bb)
	xDen := big.NewInt(1).Mul(ba, big.NewInt(2))
	return big.NewRat(xNum.Int64(), xDen.Int64()), true
}
