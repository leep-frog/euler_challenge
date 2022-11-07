package eulerchallenge

import (
	"sort"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P90() *problem {
	return noInputNode(90, func(o command.Output) {
		squares := []string{"01", "04", "09", "16", "25", "36", "49", "64", "81"}
		opts := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
		unique := map[string]bool{}
		for _, dieOneSlice := range maths.ChooseSets(opts, 6) {
			dieOne := maths.NewSimpleSet(dieOneSlice...)
			for _, dieTwoSlice := range maths.ChooseSets(opts, 6) {
				dieTwo := maths.NewSimpleSet(dieTwoSlice...)

				hasAll := true
				for _, square := range squares {
					left := square[:1]
					right := square[1:]
					oneLeft := dieOne[left] || (left == "6" && dieOne["9"]) || (left == "9" && dieOne["6"])
					oneRight := dieOne[right] || (right == "6" && dieOne["9"]) || (right == "9" && dieOne["6"])
					twoLeft := dieTwo[left] || (left == "6" && dieTwo["9"]) || (left == "9" && dieTwo["6"])
					twoRight := dieTwo[right] || (right == "6" && dieTwo["9"]) || (right == "9" && dieTwo["6"])
					if !((oneLeft && twoRight) || (oneRight && twoLeft)) {
						hasAll = false
						break
					}
				}
				if hasAll {
					parts := []string{strings.Join(dieOneSlice, ""), strings.Join(dieTwoSlice, "")}
					sort.Strings(parts)
					key := strings.Join(parts, "_")
					unique[key] = true
				}
			}
		}
		o.Stdoutln(len(unique))
	}, &execution{
		want: "1217",
	})
}
