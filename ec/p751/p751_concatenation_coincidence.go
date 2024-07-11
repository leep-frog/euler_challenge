package p751

import (
	"fmt"
	"log"
	"math/big"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
)

const PRECISION = 100

func P751() *ecmodels.Problem {
	return ecmodels.NoInputNode(751, func(o command.Output) {

		// The resulting number is non-decreasing, so we can binary search
		// until they match!
		left, right := big.NewFloat(2.0), big.NewFloat(3.0)
		theta := midpoint(left, right)
		f := opbig(theta, "")
		for f.Cmp(theta) != 0 {
			if f.Cmp(theta) > 0 {
				left = theta
			} else {
				right = theta
			}
			theta = midpoint(left, right)
			f = opbig(theta, "")
		}

		o.Stdoutln(f.Text('f', 24))

	}, &ecmodels.Execution{
		Want: "2.223561019313554106173177",
	})
}

func midpoint(a, b *big.Float) *big.Float {
	theta := new(big.Float)
	theta.SetPrec(PRECISION)
	theta.Add(a, b)
	theta.Mul(theta, big.NewFloat(0.5))
	return theta
}

func opbig(b *big.Float, s string) *big.Float {
	flrInt, _ := b.Int(nil)
	flr := new(big.Float)
	flr.SetPrec(PRECISION)
	flr.SetInt(flrInt)

	empty := len(s) == 0

	s += fmt.Sprintf(flrInt.String())
	if empty {
		s += "."
	}

	if len(s) > 28 {
		f := new(big.Float)
		f.SetPrec(PRECISION)
		f, ok := f.SetString(s)
		if !ok {
			log.Fatalf("failed to parse tao result (%s)", s)
		}
		return f
	}

	one := big.NewFloat(1.0)

	right := new(big.Float)
	right.SetPrec(PRECISION)
	right.Sub(b, flr)
	right.Add(right, one)

	nextB := new(big.Float)
	nextB.SetPrec(PRECISION)
	nextB.Mul(flr, right)
	return opbig(nextB, s)
}
