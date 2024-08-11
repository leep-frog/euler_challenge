package p725

import (
	"fmt"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/combinatorics"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
	"golang.org/x/exp/slices"
)

func P725() *ecmodels.Problem {
	return ecmodels.IntInputNode(725, func(o command.Output, n int) {
		// o.Stdoutln(n)
		// p := generator.Primes()
		// for i := 2; i <= n; i++ {
		// 	fmt.Println(i, s2(i)/i)
		// }
		fmt.Println(s2(n))

	}, []*ecmodels.Execution{
		{
			Args: []string{"2"},
			Want: "",
		},
	})
}

var (
	mod    = maths.Pow(10, 16)
	a, b   = map[int]bool{}, map[int]bool{}
	cache  = map[string]int{}
	cache2 = map[string]map[int]*maths.Int{}
)

func s(n int) int {
	var sum int
	for i := 1; i <= 9; i++ {
		sum = (sum + gen(n-1, i, 0, []int{i})) % mod
	}
	return sum
}

func gen(remDigits, remValue, min int, cur []int) int {
	code := fmt.Sprintf("%d-%d-%d-%v", remDigits, remValue, min, cur)
	if v, ok := cache[code]; ok {
		return v
	}

	if remDigits == 0 {
		if remValue != 0 {
			return 0
		}

		var sum int
		sum = combinatorics.PermutationCount(cur).ToInt()
		// for _, p := range combinatorics.Permutations(cur) {
		// 	sum = (sum + maths.Join(p)) % mod
		// 	a[maths.Join(p)] = true
		// }
		cache[code] = sum
		return sum
	}

	var sum int
	for i := min; i <= remValue && i <= 9; i++ {
		sum = (sum + gen(remDigits-1, remValue-i, i, append(cur, i))) % mod
	}
	cache[code] = sum
	return sum
}

func s2(n int) int {
	m := map[int]*maths.Int{}
	for i := 1; i <= 9; i++ {
		for k, v := range gen2(n-1, i, 0, []int{i}) {
			if c, ok := m[k]; ok {
				m[k] = c.Plus(v)
			} else {
				m[k] = v
			}
		}
	}

	s := maths.Zero()
	for k, v := range m {
		num := maths.MustIntFromString(strings.Repeat(fmt.Sprintf("%d", k), n))
		s = s.Plus(num.Times(v))
	}

	return s.DivInt(n).ModInt(mod)
}

func gen2(remDigits, remValue, min int, cur []int) map[int]*maths.Int {
	code := fmt.Sprintf("%d-%d-%d-%v", remDigits, remValue, min, cur)
	// if v, ok := cache2[code]; ok {
	// 	return v
	// }

	if remDigits == 0 {
		if remValue != 0 {
			return map[int]*maths.Int{}
		}

		combos := combinatorics.PermutationCount(cur)
		m := map[int]*maths.Int{}
		for _, c := range cur {
			if v, ok := m[c]; ok {
				m[c] = v.Plus(combos)
			} else {
				m[c] = combos
			}
		}
		cache2[code] = m
		return m
	}

	m := map[int]*maths.Int{}
	for i := min; i <= remValue && i <= 9; i++ {
		for k, v := range gen2(remDigits-1, remValue-i, i, append(cur, i)) {
			if c, ok := m[k]; ok {
				m[k] = c.Plus(v)
			} else {
				m[k] = v
			}
		}
	}
	cache2[code] = m
	return m
}

func brute(n int) int {
	p := maths.Pow(10, n)
	var sum int
	for i := 1; i <= p; i++ {
		ds := maths.Digits(i)
		slices.Sort(ds)
		first := ds[len(ds)-1]
		if bread.Sum(ds) == 2*first {
			// fmt.Println(i)
			b[i] = true
			sum += i
		}
	}
	return sum
}
