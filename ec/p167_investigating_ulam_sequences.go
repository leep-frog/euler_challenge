package eulerchallenge

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/maths"
)

type UlamSequence struct {
	a, b     int
	vs       []int
	diffs    []int
	contains []bool
	evens    []int
}

func NewUlam(a, b int) *UlamSequence {
	c := make([]bool, b+1, b+1)
	c[a] = true
	c[b] = true
	return &UlamSequence{
		a,
		b,
		[]int{a, b},
		[]int{b - a},
		c,
		[]int{a},
	}
}

func (u *UlamSequence) generateNext() int {
	for i := u.vs[len(u.vs)-1] + 1; ; i++ {
		cnt := 0
		if len(u.evens) == 2 {
			for _, e := range u.evens {
				if u.contains[i-e] && i-e != e {
					cnt++
				}
			}
		} else {
			for j := 0; u.vs[j]*2 < i; j++ {
				if u.contains[i-u.vs[j]] {
					cnt++
					if cnt == 2 {
						break
					}
				}
			}
		}
		u.contains = append(u.contains, cnt == 1)
		if cnt == 1 {
			u.diffs = append(u.diffs, i-u.vs[len(u.vs)-1])
			u.vs = append(u.vs, i)
			if i%2 == 0 {
				u.evens = append(u.evens, i)
			}
			return i
		}
	}
}

func (u *UlamSequence) at(i int) int {
	for i >= len(u.vs) {
		u.generateNext()
	}
	return u.vs[i]
}

func (u *UlamSequence) kth(k int) int {
	evenCount := 0
	var start int
	var evens []int
	for evenCount < 2 {
		if u.at(start)%2 == 0 {
			evenCount++
			evens = append(evens, u.at(start))
		}
		start++
	}
	u.evens = evens

	for patternLength := 25; ; patternLength++ {
		valid := true
		sum := 0
		for i := 0; i < patternLength; i++ {
			u.at(start + patternLength*2)
			sum += u.diffs[start+i]
			if u.diffs[start+i] != u.diffs[start+i+patternLength] {
				valid = false
				break
			}
		}

		if valid {
			value := u.at(start) + ((k-start)/patternLength)*sum
			for i := ((k - start) % patternLength); i > 0; i-- {
				value += u.at(start+i) - u.at(start+i-1)
			}
			return value
		}
	}
}

func P167() *problem {
	return intInputNode(167, func(o command.Output, n int) {

		sum := 0
		for k := 2; k <= 10; k++ {
			sum += NewUlam(2, 2*k+1).kth(maths.Pow(10, 11) - 1)
		}
		o.Stdoutln(sum)
	}, []*execution{
		{
			args: []string{"1"},
			want: "3916160068885",
		},
	})
}
