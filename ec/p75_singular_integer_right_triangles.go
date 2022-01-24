package eulerchallenge

import (
	"github.com/leep-frog/command"
)

func P75() *problem {
	return intInputNode(75, func(o command.Output, n int) {
		m := map[int][][]int{}
		// c^2 = x^2 + y^2 > x^2 + x^2 = 2x^2
		// c = Sqrt(2) * x
		// L = x + y + c > x + x + c > x + x + x*Sqrt(2) = 3.4141 * x > 3.4 * x
		//   3.4x < L -> x < L / 3.4 = 10 * L / 34
		for x := 1; x <= 10*n/34; x++ {
			// x^2 + y^2 = (y + k)^2
			// x^2 + y^2 = y^2 + 2yk + k^2
			// x^2 = 2yk + k^2
			// [k=1] x^2 = 2y + 1
			// [k=1] x^2 = 4y + 4
			// [k=1] x^2 = 6y + 9
			// [k=1] x^2 = 8y + 16
			// [var] x^2 = ay + b

			// 2y = (x^2 / k) - k
			// L = x + 2*y + k
			//   = x + (2x^2 / k) - k + k
			//   = x + 2x^2 / k
			// k = (x + 2x^2) / L
			// k >= (x + 2x^2) / L_max
			//    = (x + 2x^2) / n // didn't work last time
			x2 := x * x
			for k := 1; k < x; k++ {
				a, b := 2*k, k*k
				y := (x2 - b) / a
				L := x + y + y + k
				if L > n || y <= x {
					continue
				}
				if x2 < b {
					// this would make a negative
					break
				}
				rem := x2 - b
				if rem%a == 0 {
					m[L] = append(m[L], []int{x, y, y + k})
				}
			}
		}
		total := 0
		for _, v := range m {
			if len(v) == 1 {
				total++
			}
		}
		o.Stdoutln(total)
	})
}
