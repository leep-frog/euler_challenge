package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P111() *problem {
	return intInputNode(111, func(o command.Output, n int) {
		/*start := maths.Pow(10, n-1)
		//g := generator.NewPrimes(maths.Pow(10, n))
		g := generator.Primes()
		// Map from digit to number of that digit to list of primes
		m := map[int]map[int][]int{}
		var maxLen int

		for p := start + 1; p < start*10; p++ {
			if !generator.IsPrime(p, g) {
				continue
			}
			//for i := 0; ; i++ {
			//p := g.Nth(i)
			digits := maths.Digits(p)

			if 4*p/5 > maxLen {
				maxLen = p
				fmt.Println(maxLen)
			}

			if len(digits) < n {
				continue
			}
			if len(digits) > n {
				break
			}

			counts := map[int]int{}
			for _, d := range digits {
				counts[d]++
			}

			for k, v := range counts {
				if m[k] == nil {
					m[k] = map[int][]int{}
				}
				m[k][v] = append(m[k][v], p)
			}
		}

		var sum int
		for d := 0; d <= 9; d++ {
			var max int
			for k := range m[d] {
				max = maths.Max(max, k)
			}
			//fmt.Println(d, max, len(m[d][max]), bread.Sum(m[d][max]))
			sum += bread.Sum(m[d][max])
		}

		o.Stdoutln(sum)*/

		var sum int
		g := generator.Primes()
		for d := 0; d <= 9; d++ {
			for j := n; j >= 1; j-- {
				var os []int
				p110Generator(d, n, j, []int{}, &os, g)
				if len(os) > 0 {
					sum += bread.Sum(os)
					break
				}
			}
		}
		o.Stdoutln(sum)
	}, []*execution{
		{
			args: []string{"10"},
			want: "612407567715",
		},
		{
			args:     []string{"4"},
			want:     "273700",
			estimate: 0.25,
		},
	})
}

func p110Generator(d, remainingLen, remainingDs int, cur []int, opts *[]int, g *generator.Prime) {
	if remainingLen == 0 {
		v := maths.FromDigits(cur)
		if remainingDs == 0 && g.Contains(v) {
			*opts = append(*opts, v)
		}
		return
	}

	if remainingDs > remainingLen {
		return
	}

	if remainingDs > 0 {
		if d > 0 || len(cur) != 0 {
			cur = append(cur, d)
			p110Generator(d, remainingLen-1, remainingDs-1, cur, opts, g)
			cur = cur[:len(cur)-1]
		}
	}

	start := 1
	if len(cur) > 0 {
		start = 0
	}
	for i := start; i <= 9; i++ {
		if i == d {
			continue
		}
		cur = append(cur, i)
		p110Generator(d, remainingLen-1, remainingDs, cur, opts, g)
		cur = cur[:len(cur)-1]
	}
}
