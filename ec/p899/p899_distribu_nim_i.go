package p899

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P899() *ecmodels.Problem {
	return ecmodels.IntInputNode(899, func(o command.Output, n int) {
		o.Stdoutln(clever(maths.Pow(7, n)))
	}, []*ecmodels.Execution{
		{
			Args: []string{"1"},
			Want: "21",
		},
		{
			Args: []string{"2"},
			Want: "221",
		},
		{
			Args: []string{"3"},
			Want: "2512",
		},
		{
			Args: []string{"17"},
			Want: "10784223938983273",
		},
	})
}

func clever(n int) int {

	var sum int

	// First, sum up the squares that are totally full
	twoPow, two := 1, 2
	for ; n >= two-1; twoPow, two = twoPow+1, two*2 {
		if twoPow == 1 {
			sum++
		} else if twoPow == 2 {
			sum += 5
		} else {
			// The dots, which take the pattern:
			// 1 3 1 7 1 3 1 15 1 3 1 ...
			// This can be transformed to:
			// 1 (1+2) 1 (1+2+4), etc.
			//
			// So, there are (2^k-1) 1s, (2^(k-1)-1) 2s, (2^(k-2)-1) 4s, etc.
			//
			// which can be further simplified to:
			// 2^0 * (2^k-1) + 2^1 * (2^(k-1) - 1) + ... + 2^(k-1) * (2^1 - 1)
			// = (2^k - 2^0)  +     (2^k - 2^1)     + ... + (2^k - 2^(k-1))
			// = k * 2^k - (2^0 + 2^1 + ... + 2^(k-1))
			// = k * 2^k - (2^k - 1)
			// = k * 2^k - 2^k + 1
			// = (k - 1) * 2^k + 1

			// But we need to offset our coefficients (by two).
			// k = twoPow - 2   | 2^k = twoPow / 2^2
			// k-1 = twoPow - 3 | 2^k = twoPow / 4
			k := twoPow - 3
			// Also, multiply by two since the dots are symmetric on the left and top axes (see below)
			sum += 2 * (k*(two/4) + 1)

			// The square parts, which is just the half-perimeter of the square:
			sum += two*2 - 3
		}
	}
	twoPow--
	two /= 2

	remainingSections := (n - two + 1) / 2
	// Sections are in the pattern:
	// 1 3 1 7 1 3 1 15 1 3 1 ...
	// This can be transformed to:
	// 1 (1+2) 1 (1+2+4), etc.
	// So there are r 1s, r/2 2s, r/4 4s, etc.
	incr := 1
	for remainingSections > 0 {
		// Multiply by 2 since the grid is symmetric on the top and left axes
		sum += 2 * remainingSections * incr
		remainingSections /= 2
		incr *= 2
	}

	return sum
}

/*

  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63
  .  X  .  X  .  X  .  X  .  X  .  X  .  X  .  X  .  X  .  X  .  X  .  X  .  X  .  X  .  X  .  X  .  X  .  X  .  X  .  X  .  X  .  X  .  X  .  X  .  X  .  X  .  X  .  X  .  X  .  X  .  X  .
  X  X  .  X  X  X  .  X  X  X  .  X  X  X  .  X  X  X  .  X  X  X  .  X  X  X  .  X  X  X  .  X  X  X  .  X  X  X  .  X  X  X  .  X  X  X  .  X  X  X  .  X  X  X  .  X  X  X  .  X  X  X  .
  .  .  .  X  X  X  .  X  X  X  .  X  X  X  .  X  X  X  .  X  X  X  .  X  X  X  .  X  X  X  .  X  X  X  .  X  X  X  .  X  X  X  .  X  X  X  .  X  X  X  .  X  X  X  .  X  X  X  .  X  X  X  .
  X  X  X  X  X  X  .  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  .
  .  X  X  X  X  X  .  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  .
  X  X  X  X  X  X  .  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  .
  .  .  .  .  .  .  .  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  .
  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  .  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  .  .  .  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  .  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  .  .  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  .  .  .  .  .  .  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  .  .  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  .  .  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  .  .  .  .  .  .  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  .  .  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  .  .  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  .  .  .  .  .  .  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  .  .  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  .  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  .
  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .

*/

// This is the brute force solution used to determine the pattern (and print out the above grid)
func brute(n int) int {
	var cnt int

	for a := 1; a <= n; a++ {
		fmt.Printf("%3d", a)
	}
	fmt.Println()

	for a := 1; a <= n; a++ {
		for b := 1; b <= n; b++ {
			if !play(a, b, true) {
				cnt++
				fmt.Printf("  .")
			} else {
				fmt.Printf("  X")
			}
		}
		fmt.Println()
	}
	return cnt
}

var (
	cache = map[string]bool{}
)

func play(a, b int, player1 bool) bool {
	if a > b {
		a, b = b, a
	}

	code := fmt.Sprintf("%d-%d", a, b)
	if v, ok := cache[code]; ok {
		return player1 == v
	}

	if a == 1 && b == 1 {
		// Player loses
		return !player1
	}

	minTakeLeft := 0
	if a == b {
		minTakeLeft = 1
	}
	maxTakeLeft := a - 1

	for takeLeft := minTakeLeft; takeLeft <= maxTakeLeft; takeLeft++ {
		p1Wins := play(a-takeLeft, b-(a-takeLeft), !player1)
		if p1Wins && player1 {
			cache[code] = true
			return true
		} else if !p1Wins && !player1 {
			cache[code] = true
			return false
		}
	}

	// Here and we could not win, so we must have all losing scenarios
	return !player1
}
