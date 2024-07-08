package eulerchallenge

import (
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func P103() *problem {
	return intInputNode(103, func(o command.Output, n int) {
		best := maths.Smallest[string, int]()
		FindSpecialSet(n, 1, 0, []int{}, map[int]bool{}, map[int]bool{}, best)

		o.Stdoutln(best.BestIndex())
	}, []*execution{
		{
			args:     []string{"7"},
			want:     "20313839404245",
			estimate: 1.5,
		},
		{
			args: []string{"6"},
			want: "111819202225",
		},
		{
			args: []string{"5"},
			want: "69111213",
		},
		{
			args: []string{"4"},
			want: "3567",
		},
		{
			args: []string{"3"},
			want: "234",
		},
		{
			args: []string{"2"},
			want: "12",
		},
		{
			args: []string{"1"},
			want: "1",
		},
	})
}

func FindSpecialSet(remaining, start, curSum int, values []int, curSet map[int]bool, notAllowed map[int]bool, best *maths.Bester[string, int]) {
	if remaining == 0 {
		frontSum := values[0]
		var backSum int
		for i := 0; i < len(values)/2; i++ {
			frontSum += values[i+1]
			backSum += values[len(values)-i-1]
			if frontSum < backSum {
				return
			}
		}

		best.IndexCheck(strings.Join(parse.IntsToStrings(values), ""), curSum)
		return
	}

	for j := start; j < 100; j++ {
		if len(values) > 2 && j > values[0]+values[1] {
			return
		}
		if best.Set() && curSum+remaining*j > best.Best() {
			break
		}
		if notAllowed[j] {
			continue
		}

		toAdd := []int{j}
		valid := true
		for k := range notAllowed {
			toAdd = append(toAdd, k+j)
			if notAllowed[k+j] {
				valid = false
				break
			}
		}
		if !valid {
			continue
		}
		for _, a := range toAdd {
			notAllowed[a] = true
		}
		curSet[j] = true
		values = append(values, j)
		FindSpecialSet(remaining-1, j+1, curSum+j, values, curSet, notAllowed, best)
		values = values[:len(values)-1]
		for _, a := range toAdd {
			delete(notAllowed, a)
		}
		delete(curSet, j)
	}
}
