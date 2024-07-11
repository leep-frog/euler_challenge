package p181

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
)

func rec(depth, remB, remW, minB, minW int, cur [][]int, cache map[string]int) int {
	//fmt.Println(strings.Repeat("  ", depth), remB, remW, minB, minW)
	//fmt.Println(strings.Repeat("  ", depth), cur, remB, remW)
	code := fmt.Sprintf("%d, %d, %d, %d", remB, remW, minB, minW)
	if v, ok := cache[code]; ok {
		//fmt.Println("USED CACHE!")
		return v
	}
	if remB == 0 && remW == 0 {
		//fmt.Println("GOTIT", cur)
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
		sum += rec(depth+1, remB-minB, remW-w, minB, w, append(cur, []int{minB, w}), cache)
	}

	// Now increase blacks
	for b := minB + 1; b <= remB; b++ {
		for w := 0; w <= remW; w++ {
			sum += rec(depth+1, remB-b, remW-w, b, w, append(cur, []int{b, w}), cache)
		}
	}
	//for b := minB; b <= remB
	// First add more white ones

	cache[code] = sum
	return sum
}

func P181() *ecmodels.Problem {
	return ecmodels.IntsInputNode(181, 2, 0, func(o command.Output, ns []int) {
		fmt.Println("START")
		b := ns[0]
		w := ns[1]
		c := map[string]int{}
		o.Stdoutln(rec(0, b, w, 0, 0, nil, c))
		//fmt.Println(c)
	}, []*ecmodels.Execution{
		{
			Args:     []string{"60", "40"},
			Want:     "83735848679360680",
			Estimate: 4*60 + 30,
		},
	})
}
