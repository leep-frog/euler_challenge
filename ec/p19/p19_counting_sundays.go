package p19

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
)

var (
	months = []int{
		31, // January
		28, // February
		31, // March
		30, // April
		31, // May
		30, // June
		31, // July
		31, // August
		30, // September
		31, // October
		30, // November
		31, // December
	}
)

func P19() *ecmodels.Problem {
	return ecmodels.NoInputNode(19, func(o command.Output) {
		day := 1
		var count int
		for year := 1900; year <= 2000; year++ {
			for month := 0; month < 12; month++ {
				if day == 0 && year != 1900 {
					count++
				}
				if month == 1 && year%4 == 0 && (year%100 != 0 || year%400 == 0) {
					day = (day + 29) % 7
				} else {
					day = (day + months[month]) % 7
				}
			}
		}
		o.Stdoutln(count)
	}, &ecmodels.Execution{
		Want: "171",
	})
}
