package eulerchallenge

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P80() *problem {
	return noInputNode(80, func(o command.Output) {
		bigNum := maths.MustIntFromString("1" + strings.Repeat("0", 100))
		var sum int
		for n := 2; n <= 100; n++ {
			if maths.IsSquare(n) {
				continue
			}
			start, period := maths.SquareRootPeriod(n)
			den := maths.One()
			startIdx := 200
			num := maths.NewInt(int64(period[startIdx%len(period)]))
			for idx := startIdx - 1; idx >= 0; idx-- {
				tmp := den
				den = num
				num = tmp.Plus(den.Times(maths.NewInt(int64(period[idx%len(period)]))))
			}
			tmp := den
			den = num
			num = tmp.Plus(den.Times(maths.NewInt(int64(start))))
			//remainder := num.Minus(den.Times(maths.NewInt(int64(start))))
			remainder := num
			digits := maths.MustIntFromString(remainder.Times(bigNum).Div(den).String()[:100])
			sum += digits.DigitSum()
		}
		o.Stdoutln(sum)
	})
}
