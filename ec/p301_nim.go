package eulerchallenge

import (
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
	"golang.org/x/exp/slices"
)

// Used the function at the bottom to evaluate for each k (where stacks are (k, 2k, 3k)).
// Noticed that numbers that had consecutive 1's in their biary
// string representation. Then realized that the number of binary
// strings with consecutive 1s, which are also less than or equal to 2^n
// is simply the fibonacci sequence at term (n+1)
func P301() *problem {
	return intInputNode(301, func(o command.Output, n int) {
		o.Stdoutln(generator.Fibonaccis().Nth(n + 1))
	}, []*execution{
		{
			args: []string{"3"},
			want: "5",
		},
		{
			args: []string{"4"},
			want: "8",
		},
		{
			args: []string{"30"},
			want: "2178309",
		},
	})
}

// Explore the game tree of nim and return whether player1 wins.
func nim(player1 bool, stacks []int, cache map[string]bool) bool {
	copy := maths.CopySlice(stacks)
	slices.Sort(copy)
	code := fmt.Sprintf("%v", copy)
	if v, ok := cache[code]; ok {
		if player1 {
			return v
		}
		return !v
	}
	for i, originalSize := range stacks {
		for size := 0; size < originalSize; size++ {
			stacks[i] = size
			player1Wins := nim(!player1, stacks, cache)
			if player1Wins && player1 {

				stacks[i] = originalSize
				cache[code] = true
				return true
			}
			if !player1 && !player1Wins {
				stacks[i] = originalSize
				cache[code] = true
				return false
			}
		}
		stacks[i] = originalSize
	}
	cache[code] = false
	return !player1
}
