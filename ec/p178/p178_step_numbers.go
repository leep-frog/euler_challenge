package p178

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
)

func rec178(rem int, cur int, numberCount []int, numberPresence []bool, distinctNumberCount int, cache map[string]int) int {
	code := fmt.Sprintf("%d %d %v", rem, cur, numberPresence)
	if v, ok := cache[code]; ok {
		return v
	}
	var cnt int
	if distinctNumberCount == 10 {
		cnt++
	}
	if rem == 0 {
		return cnt
	}

	// The number below
	if cur > 0 {
		if numberCount[cur-1] == 0 {
			distinctNumberCount++
			numberPresence[cur-1] = true
		}
		numberCount[cur-1]++
		cnt += rec178(rem-1, cur-1, numberCount, numberPresence, distinctNumberCount, cache)
		numberCount[cur-1]--
		if numberCount[cur-1] == 0 {
			distinctNumberCount--
			numberPresence[cur-1] = false
		}
	}

	// The number above
	if cur < 9 {
		if numberCount[cur+1] == 0 {
			distinctNumberCount++
			numberPresence[cur+1] = true
		}
		numberCount[cur+1]++
		cnt += rec178(rem-1, cur+1, numberCount, numberPresence, distinctNumberCount, cache)
		numberCount[cur+1]--
		if numberCount[cur+1] == 0 {
			distinctNumberCount--
			numberPresence[cur+1] = false
		}
	}
	cache[code] = cnt
	return cnt
}

func P178() *ecmodels.Problem {
	return ecmodels.IntInputNode(178, func(o command.Output, n int) {
		var sum int
		cache := map[string]int{}
		// Starting with 1 and 8 are symmetric
		for i := 9; i >= 5; i-- {
			nc := make([]int, 10, 10)
			np := make([]bool, 10, 10)
			nc[i]++
			np[i] = true
			v := rec178(n-1, i, nc, np, 1, cache)
			// 9 reflects to zero, but numbers can't start with 0
			if i != 9 {
				v *= 2
			}
			sum += v
		}
		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args: []string{"8"},
			Want: "0",
		},
		{
			Args: []string{"9"},
			Want: "0",
		},
		{
			Args: []string{"10"},
			Want: "1",
		},
		{
			Args: []string{"40"},
			Want: "126461847755",
		},
	})
}
