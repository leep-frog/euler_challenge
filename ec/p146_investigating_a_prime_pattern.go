package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P146() *problem {
	return intInputNode(146, func(o command.Output, n int) {
		p := generator.Primes()
		offsets := []int{1, 3, 7, 9, 13, 27}
		offsetM := map[int]bool{}
		for _, o := range offsets {
			offsetM[o] = true
		}

		var sum int

		// allowedMods is a map from modulo-er, k, to valid values of (square % k)
		allowedMods := map[int]map[int]bool{}
		for k := 3; k < 251; k += 2 {
			possibleMods := map[int]bool{}
			for i := 0; i < k; i++ {
				possibleMods[i] = true
			}
			for _, o := range offsets {
				// square + o can't equal mod 0
				diff := k - o
				for ; diff < 0; diff += k {
				}
				delete(possibleMods, diff)
			}
			allowedMods[k] = possibleMods
		}

		for i := 10; i < n; i += 10 {
			sq := i * i
			valid := true
			for k, vs := range allowedMods {
				if !vs[sq%k] && sq > k {
					valid = false
					break
				}
			}
			if !valid {
				continue
			}

			valid = true
			for j := 1; j <= 27; j += 2 {
				if p.Contains(sq+j) != offsetM[j] {
					valid = false
					break
				}
			}
			if valid {
				sum += i
			}
		}
		o.Stdoutln(sum)
	}, []*execution{
		{
			args:     []string{"150000000"},
			want:     "676333270",
			estimate: 100,
		},
		{
			args: []string{"1000000"},
			want: "1242490",
		},
	})
}
