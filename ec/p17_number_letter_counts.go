package eulerchallenge

import (
	"github.com/leep-frog/command"
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

func P17() *command.Node {
	return command.SerialNodes(
		command.Description("Sum the letter counts of the numbers up to N"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			n := d.Int(N)

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
		}),
	)
}
