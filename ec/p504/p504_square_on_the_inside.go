package p504

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P504() *ecmodels.Problem {
	return ecmodels.IntInputNode(504, func(o command.Output, m int) {
		var cache [][]int
		for b := 0; b < m; b++ {
			var row []int
			for a := 0; a <= b; a++ {
				row = append(row, -1)
			}
			cache = append(cache, row)
		}

		var sum int
		for a := 1; a <= m; a++ {
			for b := 1; b <= m; b++ {
				for c := 1; c <= m; c++ {
					for d := 1; d <= m; d++ {
						latticePoints := 1 + count(a, b, cache) + count(b, c, cache) + count(c, d, cache) + count(d, a, cache)
						// Each line is overcounted once, but the point at the end shouldn't be included since it's on the line
						latticePoints -= a + b + c + d - 4

						if maths.IsSquare(latticePoints) {
							sum++
						}
					}
				}
			}
		}
		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args: []string{"4"},
			Want: "42",
		},
		{
			Args:     []string{"100"},
			Want:     "694687",
			Estimate: 2,
		},
	})
}

// count calculates the number of lattice points strictly contained in the triangle formed by (0, 0), (0, a), (b, 0) (not including the origin)
func count(a, b int, cache [][]int) int {
	// Always have a be smaller
	if a > b {
		a, b = b, a
	}

	// Check the cache
	v := cache[b-1][a-1]
	if v != -1 {
		return v
	}

	// Strictly contained so not on line
	// First, get the number of points in the 90-45-45 triangle
	// (1, 1) = 0
	// (2, 2) = 2         = 1 + (1)
	// (3, 3) = 2 + 3     = 2 + (1 + 2)
	// (4, 4) = 2 + 3 + 4 = 3 + (1 + 2 + 3)
	// (x, x)             = (x-1) + x*(x-1)/2
	//                    = (2(x-1) + x*(x-1))/2
	//                    = (x+2)*(x-1)/2
	count := (a + 2) * (a - 1) / 2 // This will be zero if a is 1, so no if check needed for that

	// Then, get the rest manually
	// Start at the left
	for x := 0; x <= b; x++ {
		// Get the point just below the triangle
		// In y = mx + b, b = a, m = -a/b
		y := -(a*x)/b + a

		// Simce m is negative, then we default to on the line (if a mod x is zero) or the point above it
		y--

		count += y + 1 - maths.Max(0, a-x)
	}

	// Update the cache and return
	cache[b-1][a-1] = count
	return count
}
