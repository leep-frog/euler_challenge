package p692

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/series"
)

func P692() *ecmodels.Problem {
	return ecmodels.IntInputNode(692, func(o command.Output, n int) {

		// After doing brute force, I noticed that brute(fibonacci(i)) = A055244(i)
		s := series.A055244()

		f := generator.Fibonaccis()
		idx := 0
		for ; f.Nth(idx) != n; idx++ {
		}

		o.Stdoutln(s.Get(idx))
	}, []*ecmodels.Execution{
		{
			Args: []string{"13"},
			Want: "43",
		},
		{
			Args: []string{"23416728348467685"},
			Want: "842043391019219959",
		},
	})
}

var (
	forceWin = map[string]bool{}
)

func brute(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		c := h(i)
		sum += c
	}
	return sum
}

func h(n int) int {
	for i := 1; i <= n; i++ {
		if !canForceAWin(i, n-i) {
			return i
		}
	}
	return n
}

func canForceAWin(prev, n int) bool {
	code := fmt.Sprintf("%d-%d", n, prev)
	if v, ok := forceWin[code]; ok {
		return v
	}

	if n == 0 || n <= prev*2 {
		forceWin[code] = true
		return true
	}

	var canWin bool
	for i := 1; i <= 2*prev && i <= n; i++ {
		canWin = canWin || !canForceAWin(i, n-i)
	}

	forceWin[code] = canWin

	return canWin
}
