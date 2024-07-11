package p121

import (
	"math/big"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
)

func P121() *ecmodels.Problem {
	return ecmodels.IntInputNode(121, func(o command.Output, n int) {
		r := blueDiscs(int64(n), int64((n/2)+1), 1)
		o.Stdoutln(r.Denom().Int64() / r.Num().Int64())
	}, []*ecmodels.Execution{
		{
			Args:     []string{"15"},
			Want:     "2269",
			Estimate: 0.2,
		},
		{
			Args:     []string{"4"},
			Want:     "10",
			Estimate: 0.2,
		},
	})
}

func blueDiscs(turns, needBlue, numRed int64) *big.Rat {
	if needBlue == 0 {
		return big.NewRat(1, 1)
	}
	if needBlue > turns {
		return big.NewRat(0, 1)
	}

	blueDraw := big.NewRat(0, 1).Mul(big.NewRat(1, numRed+1), blueDiscs(turns-1, needBlue-1, numRed+1))
	redDraw := big.NewRat(0, 1).Mul(big.NewRat(numRed, numRed+1), blueDiscs(turns-1, needBlue, numRed+1))
	r := big.NewRat(0, 1).Add(redDraw, blueDraw)
	return r
}
