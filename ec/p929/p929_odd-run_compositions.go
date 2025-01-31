package p929

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
)

const (
	mod = 1111124111
)

func P929() *ecmodels.Problem {
	return ecmodels.IntInputNode(929, func(o command.Output, n int) {
		o.Stdoutln(F(n))
	}, []*ecmodels.Execution{
		{
			Args:     []string{"2"},
			Estimate: 15,
			Want:     "57322484",
		},
	})
}

/***************************************************
*               APPROACH THREE (15s)              *
***************************************************/

/***

Basic approach is that we recursively calculate F(x) by considering the
sequences for x that start with a number, k (f(x, k))

Since there needs to be an odd number of k, then this is simply:
f(x, k) = [All sequences for x starting with 1 k]
+ [All sequences for x starting with multiple k's]

= [{k} concatenated with all sequences for (x-k) that don't start with k]
+ [{k,k} concatenated with all sequences for (x-2k) starting with k]

Therefore:
f(x, k) = F(x - k) - f(x - k, k) + f(x - 2k, k)

I made this realization and was able to implement FOne to get the solution.

The first post in the problem thread then points out that a recursive formula
can be used. I did the math on paper, but will refrain from copying that here
since it's tedious to type in code, but basically:

f(x, k) = F(x-k) - F(x-2k) + 2F(x-3k) - 3F(x-4k) + 5F(x-5k) - ...

And

F(x) = sum from k = 1 to x of f(x, k)

By combining the above two, we can then get:

F(x) = f(x, 1)
+ f(x, 2)
+ f(x, 3)
+ ...

= F(x-1) - F(x-2) + 2F(x-3) - 3F(x-4) + 5F(x-5) - ...
=          F(x-2)           -  F(x-4)             ...
=                    F(x-3)                       ...
=                              F(x-4)             ...
=                                        F(x-4)   ...

You can see that we can calculate the coefficients for each F(x-i)
instead of having to recur more deeply each time.

That's exactly what approach three now does.

*[1] It is simply to see that f(x, x) should equal 1

Given our equality assumption above, however:
f(x, x) = F(x - x) - f(0, x) + f(-x, x)
1 = F(0) - f(0, x) + f(-x, x)

It is reasonable to assume that f(k, *) for k < 0 should be 0
This, however implies that 1 = F(0)
Hence why the 0-th index is 1

***/

func F(n int) int {

	summedFibCoefs := make([]int, n+1)

	for k := 1; k <= n; k++ {
		a, b := 1, 1
		sign := 1
		for i := k; i <= n; i += k {
			fibCoef := sign * a
			a, b = b, (a+b)%mod
			summedFibCoefs[i] += fibCoef
			sign *= -1
		}
	}

	for i, v := range summedFibCoefs {
		summedFibCoefs[i] = (mod + v) % mod
	}

	F := []int{
		// We know that f(x, x) should equal 1
		// Given our equality assumption, however:
		// f(x, x) = F(x - x) - f(0, x) + f(-x, x)
		//
		//	1 = F(0) - f(0, x) + f(-x, x)
		//
		// It is reasonable to assume that f(k, *) for k < 0 should be 0
		// This, however implies that 1 = F(0)
		// Hence why the 0-th index is 1
		1,
		1,
	}
	for n >= len(F) {

		x := len(F)
		if x%1000 == 0 {
			fmt.Println(x)
		}

		var sum int
		for k := 1; k <= x; k++ {
			sum = (sum + (summedFibCoefs[k] * F[x-k])) % mod
		}
		F = append(F, sum)
	}

	return F[n]
}

/***************************************************
*                 APPROACH TWO (10m)              *
***************************************************/
func fTwo(x, start int) int {

	if x == start {
		return 1
	}

	if start > x {
		panic("Bad args")
	}

	sign := 1
	var sum int
	a, b := 1, 1
	for idx := 1; start*idx <= x; idx++ {
		fibCoef := sign * a
		a, b = b, (a+b)%mod
		sum = (sum + (fibCoef * FTwo(x-idx*start) % mod)) % mod
		sign *= -1
	}

	return sum

}

var (
	fCache = []int{
		1, // See *[1] for why F(0) is 1
		1,
	}
)

func FTwo(n int) int {
	for n >= len(fCache) {
		x := len(fCache)

		var sum int
		for i := 1; i <= x; i++ {
			fp := fTwo(x, i)
			sum = (sum + fp) % mod
		}

		fCache = append(fCache, sum)
	}
	return fCache[n]
}

/***************************************************
*                   APPROACH ONE (~1 hour)                  *
***************************************************/

var (
	cache = map[string]int{}
)

func fOne(n, mostRecent int) int {
	if n == 0 {
		return 1
	}

	if n < 0 {
		panic("Bad args")
	}

	code := fmt.Sprintf("%d-%d", n, mostRecent)
	if v, ok := cache[code]; ok {
		return v
	}

	var sum int
	for i := 1; i <= n; i++ {
		if i == mostRecent {
			continue
		}
		for cnt := 1; i*cnt <= n; cnt += 2 {
			sum += fOne(n-i*cnt, i)
		}
	}
	cache[code] = sum
	return sum
}

var (
	// Map from n to starting number to number of arrangements that add up to n
	// and start with starting number. Zero index is the total sum of arrangements
	// for n
	byStartCache = [][]int{
		// f(0)
		nil,
		// f(1)
		{
			// f(1)
			1,
			// f(1, 1)
			1,
		},
	}
)

func FOne(n int) int {
	for len(byStartCache) <= n {
		k := len(byStartCache)

		byKStart := make([]int, k+1)
		var total int
		for i := 1; i < k; i++ {
			// Add a single i to all arrangements of (k-i) that don't start with i
			v := byStartCache[k-i][0]
			if len(byStartCache[k-i]) > i {
				v = (v + mod - byStartCache[k-i][i]) % mod
			}

			// Add two i's to all arrangements of (k-i) that start with i
			if k-2*i > 0 && len(byStartCache[k-2*i]) > i {
				v = (v + byStartCache[k-2*i][i]) % mod
			}

			byKStart[i] = v
			total = (total + v) % mod
		}
		byKStart[k] = 1
		byKStart[0] = (total + 1) % mod
		byStartCache = append(byStartCache, byKStart)
	}
	return byStartCache[n][0]
}
