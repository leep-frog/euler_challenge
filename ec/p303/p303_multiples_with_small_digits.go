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
		fmt.Println(dfs(m, []*maths.Int{maths.NewInt(1), maths.NewInt(2)}, n))
	}, []*ecmodels.Execution{
		{
			Args: []string{"2"},
			Want: "",
		},
	})
}

// func markCompleted(m map[int]bool, n, k int, v *maths.Int) *maths.Int {
// 	sum := maths.Zero()

// 	for cur := k; cur <= n; cur *= 10 {
// 		delete(m, cur)
// 		sum = sum.Plus(v)
// 	}
// }

func dfs(m map[int]bool, opts []*maths.Int, n int) *maths.Int {

	sum := maths.Zero()

	// First, clear all the /9+/ values because they are the largest

	// for

	for len(m) > 0 {
		next := opts[0]
		opts = opts[1:]
		// fmt.Println(next)

		// maps
		// 9 12222
		// 99 1122222222
		// 999 111222222222222
		// 9999 11112222222222222222
		// 122212222222221
		// 11112222222222222222

		for _, k := range maps.Keys(m) {
			if next.ModInt(k) == 0 {
				delete(m, k)
				sum = sum.Plus(next.DivInt(k))
				fmt.Println(len(m))

				if k == 9 || k == 99 || k == 999 || k == 9999 {
					fmt.Println("nines", k, next)
				}

				for mult := k * 10; mult <= n; mult *= 10 {
					delete(m, mult)
					sum = sum.Plus(next.DivInt(k))
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
