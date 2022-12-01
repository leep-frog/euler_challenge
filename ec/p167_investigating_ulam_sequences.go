package eulerchallenge

import (
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

type UlamSequence struct {
	a, b  int
	vs    []int
	diffs []int
	// TODO: make slice?
	contains []bool
	m        map[int]int
	mod      int
	modMap   [][]int
}

func NewUlam(a, b int) *UlamSequence {
	c := make([]bool, b+1, b+1)
	c[a] = true
	c[b] = true
	mod := 1234
	modMap := make([][]int, mod, mod)
	modMap[a%mod] = []int{a}
	modMap[b%mod] = []int{b}
	return &UlamSequence{
		a,
		b,
		[]int{a, b},
		[]int{b - a},
		c,
		map[int]int{a + b: 1},
		mod,
		modMap,
	}
}

func (u *UlamSequence) generateNext() int {
	for i := u.vs[len(u.vs)-1] + 1; ; i++ {
		cnt := 0
		wantMod := i % u.mod
		for modA := 0; modA < u.mod; modA++ {
			modB := (wantMod - modA + u.mod) % u.mod
			if modB < modA {
				continue
			}
			for _, a := range u.modMap[modA] {
				for _, b := range u.modMap[modB] {
					if modA == modB && b < a {
						continue
					}
					if a == b {
						continue
					}
					if a+b == i {
						cnt++
					}
					// bs are sorted in increasing order
					if a+b > i {
						break
					}
					if cnt == 2 {
						goto DONZO
					}
				}
			}
		}
	DONZO:
		if cnt == 1 {
			u.diffs = append(u.diffs, i-u.vs[len(u.vs)-1])
			u.vs = append(u.vs, i)
			u.modMap[wantMod] = append(u.modMap[wantMod], i)
			return i
		}
	}
}

func (u *UlamSequence) midGenerateNext() int {

	for i := u.vs[len(u.vs)-1] + 1; ; i++ {
		cnt := 0
		for j := 0; u.vs[j]*2 < i; j++ {
			if u.contains[i-u.vs[j]] {
				cnt++
				if cnt == 2 {
					break
				}
			}
		}
		if cnt == 1 {
			u.diffs = append(u.diffs, i-u.vs[len(u.vs)-1])
			u.vs = append(u.vs, i)
			return i
		}
	}
}

func (u *UlamSequence) oldgenerateNext() int {
	for i := u.vs[len(u.vs)-1]; ; i++ {
		cnt := u.m[i]
		delete(u.m, i)
		if cnt == 1 {
			for _, v := range u.vs {
				u.m[v+i]++
			}
			u.diffs = append(u.diffs, i-u.vs[len(u.vs)-1])
			u.vs = append(u.vs, i)
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
	for evenCount < 2 {
		if u.at(start)%2 == 0 {
			evenCount++
		}
		start++
	}

	for patternLength := 25; ; patternLength++ {
		if patternLength%10_000 == 0 {
			fmt.Println("PL", u.b, patternLength)
		}
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

		fmt.Println("START")
		// fmt.Println("U", NewUlam(2, 5).generateNext())
		fmt.Println("U", NewUlam(2, 5).at(50))
		// return

		sum := 0
		for k := 2; k <= 10; k++ {
			v := NewUlam(2, 2*k+1).kth(maths.Pow(10, 11) - 1)
			fmt.Println(k, v)
			sum += v
		}
		fmt.Println(sum)

		return
	}, []*execution{
		{
			args: []string{"1"},
			want: "",
		},
	})
}
