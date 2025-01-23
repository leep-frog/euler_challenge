package p303

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
	"golang.org/x/exp/maps"
)

func P303() *ecmodels.Problem {
	return ecmodels.IntInputNode(303, func(o command.Output, n int) {

		m := map[int]bool{}
		for i := 1; i <= n; i++ {
			m[i] = true
		}
		o.Stdoutln(dfs(m, []*maths.Int{maths.NewInt(1), maths.NewInt(2)}, n))
	}, []*ecmodels.Execution{
		{
			Args: []string{"2"},
			Want: "11363107",
		},
		{
			Args: []string{"4"},
			Want: "1111981904675169",
		},
	})
}

func markCompleted(m map[int]bool, n, k int, v *maths.Int) *maths.Int {
	sum := maths.Zero()

	for cur := k; cur <= n; cur *= 10 {
		delete(m, cur)
		// This intentionally uses k because the ratio stays the same:
		// (v / k) == (10v / 10k) == (100v / 100k) == ...
		sum = sum.Plus(v.DivInt(k))
	}
	return sum
}

func dfs(m map[int]bool, opts []*maths.Int, n int) *maths.Int {

	sum := maths.Zero()

	// First, clear all the /9+/ values because they require the largest values
	// but create a very simple pattern (k ones followed by 4k twos)
	for nines, ones, twos := 9, []int{1}, []int{2, 2, 2, 2}; nines <= n; nines, ones, twos = nines*10+9, append(ones, 1), append(twos, 2, 2, 2, 2) {
		sum = sum.Plus(markCompleted(m, n, nines, maths.IntFromDigits(append(ones, twos...))))
	}

	for len(m) > 0 {
		next := opts[0]
		opts = opts[1:]

		for _, k := range maps.Keys(m) {
			if next.ModInt(k) == 0 {

				sum = sum.Plus(markCompleted(m, n, k, next))
				fmt.Println(len(m))

				if k == 9 || k == 99 || k == 999 || k == 9999 {
					fmt.Println("nines", k, next)
				}

				if len(m) < 20 {
					fmt.Println(sum, "|", k, next, m)
				}

			}
		}

		opts = append(opts, next.TimesInt(10), next.TimesInt(10).PlusInt(1), next.TimesInt(10).PlusInt(2))
	}
	return sum
}
