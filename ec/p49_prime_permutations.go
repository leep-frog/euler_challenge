package eulerchallenge

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P49() *problem {
	return noInputNode(49, func(o command.Output) {
		primes := generator.Primes()

		var fourDig []int
		for i, p := 0, primes.Nth(0); p < 10_000; i, p = i+1, primes.Nth(i+1) {
			if p >= 1000 {
				fourDig = append(fourDig, p)
			}
		}

		for i := 0; i < len(fourDig); i++ {
			pi := fourDig[i]
			for j := i + 1; j < len(fourDig) && 2*fourDig[j]-pi < 10_000; j++ {
				pj := fourDig[j]
				pk := 2*pj - pi
				if !sameDigits(pi, pj) {
					continue
				}
				if !sameDigits(pi, pk) {
					continue
				}
				if primes.Contains(pk) {
					o.Stdoutf("%d%d%d ", pi, pj, pk)
				}
			}
		}
		o.Stdoutln()
	}, &execution{
		want: "148748178147 296962999629 ",
	})
}

func sameDigits(this, that int) bool {
	return mapEq(digitMap(this), digitMap(that))
}

func mapEq(this, that map[int]int) bool {
	if len(this) != len(that) {
		return false
	}

	for k, v := range this {
		if that[k] != v {
			return false
		}
	}
	return true
}

func digitMap(n int) map[int]int {
	m := map[int]int{}
	for ; n > 0; n /= 10 {
		m[n%10]++
	}
	return m
}
