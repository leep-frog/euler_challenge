package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/maths"
)

// Related to problem 175

func rec169(rem *maths.Int, cache map[string]int) int {
	if v, ok := cache[rem.String()]; ok {
		return v
	}

	if rem.IsZero() {
		return 1
	}

	var r []int
	if rem.ModInt(2) == 1 {
		// We need exactly one of the current power of 2
		r = append(r, rec169(rem.Minus(maths.NewInt(1)).DivInt(2), cache))
	} else {
		// We can have either zero or two of the current power of 2
		r = append(r, rec169(rem.DivInt(2), cache))
		r = append(r, rec169(rem.Minus(maths.NewInt(2)).DivInt(2), cache))
	}
	s := bread.Sum(r)
	cache[rem.String()] = s
	return s
}

func powTwoBelow(k *maths.Int) (*maths.Int, *maths.Int) {
	two := maths.NewInt(2)
	for two.LT(k) {
		two = two.Times(maths.NewInt(2))
	}
	sr := two.Times(maths.NewInt(2)).Minus(maths.NewInt(2))
	two = two.DivInt(2)
	return two, sr
}

func P169() *problem {
	return intInputNode(169, func(o command.Output, n int) {
		// Noticed the following pattern:
		// f(x), f(2x), f(4x) make a linear pattern
		// let d = f(2x) - f(x)
		// f(x * 2^k) = f(x) + k * d
		// Let y = 10^n, then f(y) = f(5^n * 2^n) = f(5^n) + n * d
		// Let x = 5^n
		// f(y) = f(x) + n * d
		// f(y) = f(x) + n * (f(2x) - f(x))
		x := maths.BigPow(5, n)
		twoX := x.Times(maths.NewInt(2))
		cache := map[string]int{}
		fx, f2x := rec169(x, cache), rec169(twoX, cache)
		d := f2x - fx
		// Although, my DP approach works really well:
		// rec169(maths.BigPow(10, n), cache) runs basically as quickly as this optimized approach.
		o.Stdoutln(fx + n*d)
	}, []*execution{
		{
			args: []string{"1"},
			want: "5",
		},
		{
			args: []string{"2"},
			want: "19",
		},
		{
			args: []string{"25"},
			want: "178653872807",
		},
	})
}
