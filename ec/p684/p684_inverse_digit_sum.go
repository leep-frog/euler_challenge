package p684

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

const (
	mod = uint64(1_000_000_007)
)

func P684() *ecmodels.Problem {
	return ecmodels.NoInputWithExampleNode(684, func(o command.Output, ex bool) {

		var seq []uint64
		if ex {
			seq = append(seq, 20)
		} else {
			f := generator.CustomFibonacci(0, 1)
			for i := 2; i <= 90; i++ {
				seq = append(seq, uint64(f.Nth(i)))
			}
		}

		var sum uint64
		for _, s := range seq {
			sum = (sum + S(s)) % mod
		}
		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args: []string{"-x"},
			Want: "1074",
		},
		{
			Want: "922058210",
		},
	})
}

func S(n uint64) uint64 {

	// Let N = floor(n/9)
	N := int(n / 9)

	// S(n) = [1 + 2 + ... + 9  ] + [19 + 29 + ... + 99          ] + [199 + ...] + ...
	//      = [(sum from 1 to 9)] + [10 * (sum from 1 to 9) + 9*9] + [100 * (sum from 1 to 9) + 9*99] + ...

	// sum from 1 to 9 = 45

	//      = 45 * (1 + 10 + 100) + 9 * (9 + 99 + ...)
	//      = 45 * (1 + 10 + 100) + 81 * (1 + 11 + ...)
	//      = 45 * (sum from k=0 to N-1 {10^k}) + 81 * (sum from j=1 to N-1 {sum from i=0 to j-1 {10^k}})

	// First sum closed form: https://www.wolframalpha.com/input?i=sum+from+k+%3D+0+to+n-1+of+10%5Ek
	// Second sum closed form: https://www.wolframalpha.com/input?i=sum+from+n%3D1+to+n%3Dm+%28sum+from+0+to+n-1+of+10%5Ek%29

	//      = 45 * 1/9 * (10^N - 1) + 81 * [(1/81) * (-9N + 10^N - 1)]

	//      = 45* 1/9 * (10^N - 1) + -9N + 10^N - 1
	//      = 5 * (10^N - 1) + -9N + 10^N - 1
	//      = 5*10^N - 5 + -9N + 10^N - 1
	//      = 6*10^N - 9N - 6

	tenPow := tenPowValue(n / 9)

	term := maths.NewInt(int(tenPow)).TimesInt(6).MinusInt(N*9 + 6)

	// The remainder is 199 + 299 + ... + 599 (if n % 9 were 5)
	//                = 200 + 300 + ... + 600 - (n % 9)
	//                = sum from i = 1 to n % 9 of 100 + i*100 - (n % 9)
	//                = 10^N*(sum from i = 1 to n % 9 of 1 + i) - (n % 9)
	//                = (10^N)*(n%9) + 10^N*(sum from i = 1 to n % 9 of i) - (n % 9)

	remCount := int(n % 9)
	var remProd int
	for i := 1; i <= remCount; i++ {
		remProd += i + 1
	}

	rem := maths.NewInt(int(tenPow)).TimesInt(remProd).MinusInt(remCount)

	return uint64(term.Plus(rem).ModInt(int(mod)))
}

// tenPowValue returns 10^n mod 1_000_000_007 by recursively
// getting half the value and multiplying by 2
func tenPowValue(n uint64) uint64 {
	// Base cases
	if n == 0 {
		return 1
	} else if n == 1 {
		return 10
	}

	// Recursively get half the value
	half := tenPowValue(n / 2)
	full := (half * half) % mod

	if n%2 == 1 {
		return (full * 10) % mod
	}

	return full
}
