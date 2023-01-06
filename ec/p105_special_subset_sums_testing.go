package eulerchallenge

import (
	"sort"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/parse"
)

func P105() *problem {
	return fileInputNode(105, func(lines []string, o command.Output) {
		var total int
		for _, nums := range parse.ToGrid(lines, ",") {
			sort.Ints(nums)
			if verifySpecialSubsetSum(nums) {
				total += bread.Sum(nums)
			}
			//set := maths.Set(nums...)
		}
		o.Stdoutln(total)
	}, []*execution{
		{
			args: []string{"p105_sets.txt"},
			want: "73702",
		},
	})
}

func verifySpecialSubsetSum(values []int) bool {
	frontSum := values[0]
	var backSum int
	for i := 0; i < len(values)/2; i++ {
		frontSum += values[i+1]
		backSum += values[len(values)-i-1]
		if frontSum < backSum {
			return false
		}
	}

	uniqueSums := map[int]bool{}
	for _, i := range values {
		toAdd := []int{i}
		for k := range uniqueSums {
			toAdd = append(toAdd, k+i)
		}
		for _, a := range toAdd {
			if uniqueSums[a] {
				return false
			}
			uniqueSums[a] = true
		}
	}

	return true
}
