package p17

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
)

var (
	twoDigits = map[int]string{
		2: "twenty",
		3: "thirty",
		4: "forty",
		5: "fifty",
		6: "sixty",
		7: "seventy",
		8: "eighty",
		9: "ninety",
	}
	letterNumberMap = map[int]string{
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "fifteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
	}
)

func P17() *ecmodels.Problem {
	return ecmodels.IntInputNode(17, func(o command.Output, n int) {
		var counts []string
		for i := 1; i < 100; i++ {
			if i < 20 {
				counts = append(counts, letterNumberMap[i])
			} else {
				counts = append(counts, twoDigits[i/10]+letterNumberMap[i%10])
			}
		}

		for i := 1; i <= 9; i++ {
			curHun := letterNumberMap[i] + "hundred"
			counts = append(counts, curHun)
			for i := 0; i < 99; i++ {
				counts = append(counts, curHun+"and"+counts[i])
			}
		}

		counts = append(counts, "onethousand")

		var sum int
		for i := 0; i < n; i++ {
			sum += len(counts[i])
		}

		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args: []string{"5"},
			Want: "19",
		},
		{
			Args: []string{"1000"},
			Want: "21124",
		},
	})
}
