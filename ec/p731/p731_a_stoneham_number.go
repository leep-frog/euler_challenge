package p731

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P731() *ecmodels.Problem {
	return ecmodels.IntInputNode(731, func(o command.Output, n int) {
		digitLength := 11

		// n is now the right most digit we need
		n = maths.Pow(10, n) + digitLength

		var val int
		for i, decimalOffset := 1, 3; decimalOffset <= n; i, decimalOffset = i+1, decimalOffset*3 {
			patternLen := maths.Max(1, decimalOffset/9)
			patternStart := (n - decimalOffset) % patternLen

			var curVal int
			// Extra buffer in case of addition overflow
			for j := 15; j >= 0; j-- {
				offset := (patternStart - j) % patternLen
				if offset < 0 {
					offset += patternLen
				}
				curVal = 10*curVal + maths.NthDigit(decimalOffset, offset+1)
			}
			val += curVal
		}
		o.Stdoutln((val / maths.Pow(10, digitLength-8)) % 10_000_000_000)
	}, []*ecmodels.Execution{
		{
			Args: []string{"2"},
			Want: "4938271604",
		},
		{
			Args: []string{"8"},
			Want: "2584642393",
		},
		{
			Args: []string{"16"},
			Want: "6086371427",
		},
	})
}

func pattern(k int) []int {
	var pattern []int
	visited := map[string]bool{}
	for rem := 1; !visited[fmt.Sprintf("%d", rem)]; {
		visited[fmt.Sprintf("%d", rem)] = true
		pattern = append(pattern, rem/k)
		rem = (rem % k) * 10
	}
	return pattern[1:]
}

func patternUpTo(k, maxSize int) []int {
	var pattern []int
	visited := map[string]bool{}
	for rem := 1; !visited[fmt.Sprintf("%d", rem)] && len(pattern) <= maxSize; {
		visited[fmt.Sprintf("%d", rem)] = true
		pattern = append(pattern, rem/k)
		rem = (rem % k) * 10
	}
	return pattern[1:]
}
