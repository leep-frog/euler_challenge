package p188

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P188() *ecmodels.Problem {
	return ecmodels.IntInputNode(188, func(o command.Output, n int) {
		o.Stdoutln(clever(1777, 1855, maths.Pow(10, 8)))
	}, []*ecmodels.Execution{
		{
			Args: []string{"2"},
			Want: "95962097",
		},
	})
}

func brute(a, b int) *maths.Int {
	if b == 1 {
		return maths.NewInt(a)
	}

	prev := brute(a, b-1)
	r := maths.BigPow(a, prev.ToInt())
	return r
}

func clever(a, b, mod int) int {
	if b == 1 {
		return a % mod
	}

	cycle := cycleLen(a, mod)
	modCycle := clever(a, b-1, len(cycle))
	return cycle[modCycle%len(cycle)]
}

func cycleLen(a, mod int) []int {
	cycle := []int{1}
	has := map[int]bool{
		1: true,
	}

	for cur := a; !has[cur]; cur = (cur * a) % mod {
		cycle = append(cycle, cur)
		has[cur] = true
	}

	return cycle
}
