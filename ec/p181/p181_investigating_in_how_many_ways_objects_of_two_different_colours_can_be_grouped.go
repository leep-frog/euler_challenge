package p181

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
)

func rec(remB, remW, minB, minW int, cur [][]int, cache []int, codeFunc func(b1, b2, w1, w2 int) int) int {
	code := codeFunc(remB, minB, remW, minW)
	if cache[code] >= 0 {
		return cache[code]
	}

	if remB == 0 && remW == 0 {
		return 1
	}

	if minB > remB {
		return 0
	}

	var sum int
	// Ordering is black is strictly increasing,
	// while white is strictly increasing for the same black.

	// First do the same number of blacks, then white must be >=
	for w := minW; w <= remW; w++ {
		if minB == 0 && w == 0 {
			continue
		}
		sum += rec(remB-minB, remW-w, minB, w, append(cur, []int{minB, w}), cache, codeFunc)
	}

	// Now increase blacks
	for b := minB + 1; b <= remB; b++ {
		for w := 0; w <= remW; w++ {
			sum += rec(remB-b, remW-w, b, w, append(cur, []int{b, w}), cache, codeFunc)
		}
	}

	cache[code] = sum
	return sum
}

func P181() *ecmodels.Problem {
	return ecmodels.IntsInputNode(181, 2, 0, func(o command.Output, ns []int) {
		b := ns[0]
		w := ns[1]
		if b > w {
			b, w = w, b
		}

		c := make([]int, (b+1)*(b+1)*(w+1)*(w+1))
		for i := range c {
			c[i] = -1
		}
		codeFunc := func(b1, b2, w1, w2 int) int {
			return (b+1)*(w+1)*(w+1)*b1 + (w+1)*(w+1)*b2 + (w+1)*w1 + w2
		}

		o.Stdoutln(rec(b, w, 0, 0, nil, c, codeFunc))
	}, []*ecmodels.Execution{
		{
			Args:     []string{"60", "40"},
			Want:     "83735848679360680",
			Estimate: 15,
		},
	})
}
