package eulerchallenge

import (
	"math/big"

	"github.com/leep-frog/command/command"
)

func P121() *problem {
	return intInputNode(121, func(o command.Output, n int) {
		r := blueDiscs(int64(n), int64((n/2)+1), 1)
		o.Stdoutln(r.Denom().Int64() / r.Num().Int64())
	}, []*execution{
		{
			args:     []string{"15"},
			want:     "2269",
			estimate: 0.2,
		},
		{
			args:     []string{"4"},
			want:     "10",
			estimate: 0.2,
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
