package p146

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P146() *ecmodels.Problem {
	return ecmodels.IntInputNode(146, func(o command.Output, n int) {
		p := generator.Primes()
		offsets := []int{1, 3, 7, 9, 13, 27}
		offsetM := make([]bool, 28)
		for _, o := range offsets {
			offsetM[o] = true
		}

		var sum int

		// allowedMods is a map from modulo-er, k, to valid values of (square % k)
		allowedModsMap := map[int][]bool{}
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
			allowedModsMap[k] = setToSlice(possibleMods)
		}

		allowedMods := setToSlice(allowedModsMap)

		// Noticed that all answers were divisble by 10, so increment by that much
		for i := 10; i < n; i += 10 {
			sq := i * i
			valid := true
			for k, vs := range allowedMods {
				if len(vs) == 0 {
					continue
				}
				md := sq % k
				if (md >= len(vs) || !vs[sq%k]) && sq > k {
					valid = false
					break
				}
			}
			if !valid {
				continue
			}

			valid = true
			var needToConfirm []int

			// Check optimistic prime test
			for j := 1; valid && j <= 27; j += 2 {

				// Confirm it's a prime
				if offsetM[j] {
					if !p.FermatContains(sq+j, 10) {
						valid = false
						// Don't need to confirm, as we know this value is not a prime
					} else {
						// Need to absolutely confirm it's prime because we might have just missed
						needToConfirm = append(needToConfirm, j)
					}
					continue
				}

				// If we think it's a prime when it shouldn't be, we still need to confirm that
				if p.FermatContains(sq+j, 10) {
					needToConfirm = append(needToConfirm, j)
				}
			}

			// Check if it's actually prime
			if valid {
				for _, j := range needToConfirm {
					if p.Contains(sq+j) != offsetM[j] {
						valid = false
						break
					}
				}
			}

			if valid {
				sum += i
			}
		}
		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args: []string{"1000000"},
			Want: "1242490",
		},
		{
			Args:     []string{"150000000"},
			Want:     "676333270",
			Estimate: 30,
		},
	})
}

func setToSlice[T any](m map[int]T) []T {
	max := maths.Largest[int, int]()
	for k := range m {
		max.Check(k)
	}
	r := make([]T, max.Best()+1)
	for k, v := range m {
		r[k] = v
	}
	return r
}
