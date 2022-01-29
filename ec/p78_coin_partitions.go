package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P78() *problem {
	return intInputNode(78, func(o command.Output, n int) {
		totals := map[int]*maths.Int{
			0: maths.One(),
			1: maths.One(),
		}

		for i := 2; ; i++ {
			totals[i] = maths.Zero()
			var jump int
			oddOffset, evenOffset := 1, 1
			var indicator, positive bool
			sum := maths.Zero()
			for j := i; j >= 0; j -= jump {
				if indicator {
					jump = oddOffset
					oddOffset++
					positive = !positive
				} else {
					jump = evenOffset
					evenOffset += 2
				}
				indicator = !indicator
				if positive {
					sum = sum.Plus(totals[j])
				} else {
					sum = sum.Minus(totals[j])
				}
			}
			totals[i] = sum
			if totals[i].Part(0)%1000000 == 0 {
				o.Stdoutln(i, sum)
				return
			}
		}

		/*
			// Map from number, k, to max_value to number of unique ways
			// k can be made from values less than or equal to max_value.
			byMax := map[int]map[int]*maths.Int{}
			maths.Set(byMax, 1, 1, maths.One())

			// Map from number to number of ways that can be uniquely made
			totals := map[int]*maths.Int{}
			totals[1] = maths.One()

			for i := 2; ; i++ {
				curTotal := maths.One() // For j == i
				maths.Set(byMax, i, 1, maths.One())
				for j := 2; j < i; j++ {
					if j > i/2 {
						curTotal = curTotal.Plus(totals[i-j])
					} else {
						curTotal = curTotal.Plus(byMax[i-j][j])
					}
					//o.Stdoutln(i, j, curTotal)
					maths.Set(byMax, i, j, curTotal)
				}
				curTotal = curTotal.Plus(maths.One())
				maths.Set(byMax, i, i, curTotal)
				totals[i] = curTotal
				if totals[i].Negative() {
					o.Stdoutln("oops", i, totals[i])
					return
				}
				o.Stdoutln(i, totals[i])
				if _, m := totals[i].Divide(maths.NewInt(10000)); m.IsZero() {
					//o.Stdoutln(i, totals[i])
					return
				}
			}
			var ks []int
			for k := range totals {
				ks = append(ks, k)
			}
			sort.Ints(ks)
			for _, k := range ks {
				o.Stdoutln(k, totals[k])
			}
		*/
	})
}

func dfs78(remaining, value int) (map[int]int, map[int]int) {
	m := map[int]int{}
	maxM := map[int]int{}
	dfs_78(remaining, value, 0, value, m, maxM)
	return m, maxM
}

func dfs_78(remaining, value, curLen, max int, m, maxM map[int]int) {
	if remaining == 0 {
		m[curLen]++
		maxM[max]++
		return
	}
	for i := value; i <= remaining; i++ {
		if i > max {
			max = i
		}
		dfs_78(remaining-i, i, curLen+1, max, m, maxM)
	}
}
