package eulerchallenge

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/generator"
)

/* This solution uses an ordering of 2^t * 3^h values.
 * If the term is v_1 + v_2 + v_3 + ... + v_k where v_i = 2^(t_i) * 3^(h_i),
 * then t_i > t_(i+1) and h_i < h_(i+1) for all values of i.
 *
 * This solution dynamically calculates all values that can be generated
 * when the first term is 2^t * 3^k. The `bigMap333` stores
 * that set of values, which can they be used as a cache when calculating
 * all values that can be generated when starting at higher-order terms (i.e. earlier in the ordering).
 */

type bigMap333 map[int]map[int]map[int]int

func (bm bigMap333) increment(t, h, v, cnt int) {
	if bm[t] == nil {
		bm[t] = map[int]map[int]int{}
	}
	if bm[t][h] == nil {
		bm[t][h] = map[int]int{}
	}
	bm[t][h][v] += cnt
}

func (bm bigMap333) getMap(t, h int) map[int]int {
	// no nil checks because all sub-nodes should be set
	return bm[t][h]
}

// generates all possible elements that can come after 2^num2s * 3^num3s (see ordering description at top of file).
func generateChildren333(num2s, num3s, max3s int) [][]int {
	var r [][]int
	for h := num3s + 1; h <= max3s; h++ {
		for t := num2s - 1; t >= 0; t-- {
			r = append(r, []int{t, h})
		}
	}
	return r
}

func P333() *problem {
	return intInputNode(333, func(o command.Output, n int) {
		// Map from number of threes to number of twos to a number to the number
		// of times that number can be made when the first term is 2^t * 3^h
		m := bigMap333{}

		start := 0
		threeVal := 1
		for ; threeVal <= n; start, threeVal = start+1, threeVal*3 {
		}
		start--
		threeVal /= 3

		for h := start; h >= 0; h, threeVal = h-1, threeVal/3 {
			for t, twoVal := 0, 1; twoVal*threeVal <= n; t, twoVal = t+1, twoVal*2 {
				m.increment(t, h, twoVal*threeVal, 1)
				for _, child := range generateChildren333(t, h, start) {
					var incrs [][]int
					for childVal, cnt := range m.getMap(child[0], child[1]) {
						incrs = append(incrs, []int{twoVal*threeVal + childVal, cnt})

					}

					for _, incr := range incrs {
						m.increment(t, h, incr[0], incr[1])
					}
				}
			}

		}

		// Construct final values
		finMap := map[int]int{}
		for _, tm := range m {
			for _, hm := range tm {
				for k, cnt := range hm {
					finMap[k] += cnt
				}
			}
		}

		ps := generator.Primes()
		sum := 0
		for g, p := ps.Start(0); p <= n; p = g.Next() {
			if finMap[p] == 1 {
				sum += p
			}
		}

		o.Stdoutln(sum)
	}, []*execution{
		{
			args: []string{"100"},
			want: "233",
		},
		{
			args:     []string{"1000000"},
			want:     "3053105",
			estimate: 8,
		},
	})
}
