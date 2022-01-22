package eulerchallenge

import (
	"strconv"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func P30() *problem {
	return intInputNode(30, func(o command.Output, n int) {
		v := 10
		for i := 0; i < n; i++ {
			v *= 10
		}

		var superSum int
		for i := 2; i < v; i++ {
			str := strconv.Itoa(i)
			var sum int
			for j := 0; j < len(str); j++ {
				digit := maths.Pow(parse.Atoi(str[j:j+1]), n)
				sum += digit
			}
			if sum == i {
				superSum += i
			}
		}
		o.Stdoutln(superSum)
	})
}
