package eulerchallenge

import (
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func generate(remaining, min, squareSum int, cur []int, all *[][]int) {
	if remaining == 0 {
		if maths.IsSquare(squareSum) {
			*all = append(*all, maths.CopySlice(cur))
		}
		return
	}

	for j := min; j <= 9; j++ {
		generate(remaining-1, j, squareSum+j*j, append(cur, j), all)
	}
}

func P171() *problem {
	return intInputNode(171, func(o command.Output, n int) {
		sum := maths.Zero()
		for i := 1; i <= n; i++ {
			fmt.Println(i, "=============")

			var all [][]int
			generate(i, 1, 0, nil, &all)
			for i, numbers := range all {
				if i%10 == 0 {
					fmt.Printf("%d/%d", i, len(all))
				}
				perms := maths.Permutations(numbers)
				for _, perm := range perms {
					if perm[0] == 0 {
						continue
					}

					// sum = sum.Plus(maths.IntFromDigits(perm))
				} /**/
				_ = numbers
			}
			fmt.Println(len(all))
		}

		fmt.Println(sum)
		//fmt.Println(len(maths.Permutations(r, 3, false)))
	}, []*execution{
		{
			args: []string{"20"},
			want: "",
		},
	})
}
