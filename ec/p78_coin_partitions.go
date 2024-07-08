package eulerchallenge

import (
	"github.com/leep-frog/command/command"
)

func P78() *problem {
	return intInputNode(78, func(o command.Output, n int) {
		totals := map[int]int{
			0: 1,
			1: 1,
		}

		for i := 2; ; i++ {
			totals[i] = 0
			var jump int
			oddOffset, evenOffset := 1, 1
			var indicator, positive bool
			var sum int
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
					sum += totals[j]
				} else {
					sum += 1_000_000 - totals[j]
				}
			}
			sum = sum % 1_000_000
			totals[i] = sum
			if sum == 0 {
				o.Stdoutln(i)
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
	}, []*execution{
		{
			args: []string{"1000000"},
			//want: "55374 36325300925435785930832331577396761646715836173633893227071086460709268608053489541731404543537668438991170680745272159154493740615385823202158167635276250554555342115855424598920159035413044811245082197335097953570911884252410730174907784762924663654000000",
			want:     "55374",
			estimate: 0.45,
		},
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
