package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P64() *command.Node {
	return command.SerialNodes(
		command.Description("https://projecteuler.net/problem=64"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			n := d.Int(N)
			var count int
			for k := 2; k <= n; k++ {
				if maths.IsSquare(k) {
					continue
				}
				remainder := map[int]map[int]bool{}
				as := []int{1}
				for i := 1; i*i < k; i++ {
					as[0] = i
				}
				num := 1
				den := as[0]
				for !remainder[num][den] && num != 0 {
					maths.Set(remainder, num, den, true)
					tmpDen := (k - den*den) / num
					newNum := den
					count := 0
					for ; (as[0] + newNum) >= tmpDen; newNum -= tmpDen {
						count++
					}
					as = append(as, count)
					num, den = tmpDen, -newNum
				}
				if len(as)%2 == 0 {
					count++
				}
			}
			o.Stdoutln(count)
		}),
	)
}
