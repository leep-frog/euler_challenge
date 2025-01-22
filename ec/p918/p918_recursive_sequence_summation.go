package p918

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P918() *ecmodels.Problem {
	return ecmodels.IntInputNode(918, func(o command.Output, n int) {
		o.Stdoutln(s(maths.Pow(10, n)))
	}, []*ecmodels.Execution{
		{
			Args: []string{"1"},
			Want: "-13",
		},
		{
			Args: []string{"12"},
			Want: "-6999033352333308",
		},
	})
}

var (
	dCache = map[int]int{
		1: 1,
	}
)

func a(k int) int {
	if k < 0 {
		panic("Unexpected negative number")
	}

	if k == 0 {
		return 0
	}

	if v, ok := dCache[k]; ok {
		return v
	}

	if k%2 == 0 {
		v := 2 * a(k/2)
		dCache[k] = v
	} else {
		n := (k - 1) / 2
		v := a(n) - 3*a(n+1)
		dCache[k] = v
	}

	return dCache[k]
}

func s(n int) int {
	if n == 0 {
		return 0
	}

	q := n / 4

	// Consider the sequence
	// a_1, a_2, ..., a_q, ..., a_2q
	//
	// And
	// a_(2q+1) + a_(2q+2) + a_(2q+3) + ... + a_(4q - 2) + a_(4q - 1) + a_4q
	//  { odd }   { even }   { odd }           { even }     { odd }   { even }
	// = [ a_q - 3 * a_(q+1) ] + [ 2 * a_(q+1) ] + [ a_(q+1) - 3 * a_(q + 2) ] + ... + [ 2 * a_(2q - 1) ] + [ a_(2q - 1) - 3 * a_2q ] + [ 2 * a_2q ]
	// Lots of things cancel (because -3, +2, +1)
	// = a_q - a_2q
	//
	// Therefore,
	// a_1 + ... + a_4q = (a_1 + ... + a_2q) + (a_(2q+1) + ... + a_4q)
	//                  = (     s(2q)      ) + a_q - a_2q

	sum := s(2*q) + a(q) - a(2*q)

	for k := q*4 + 1; k <= n; k++ {
		sum += a(k)
	}

	return sum
}
