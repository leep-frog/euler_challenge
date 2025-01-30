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
		// var sum int

		// for i := 1; i <= n; i++ {
		// 	fmt.Println("B", i, f(i, -1), byStart(i))
		// 	fmt.Println(byStartCache)
		// }
		fmt.Println(byStart(n))
	}, []*ecmodels.Execution{
		{
			Args: []string{"2"},
			Want: "",
		},
	})
}

var (
	cache = map[string]int{}
)

func f(n, mostRecent int) int {
	if n == 0 {
		return 1
	}

	if n < 0 {
		panic("NO")
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
			sum += f(n-i*cnt, i)
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

func byStart(n int) int {
	for len(byStartCache) <= n {
		k := len(byStartCache)

		if k%100 == 0 {
			fmt.Println("CALCING FOR", k)
		}

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
	// var sum int
}

// var (
// 	// Map from n to starting number to number of arrangements that add up to n
// 	// and start with starting number. Zero index is the total sum of arrangements
// 	// for n
// 	byStartMapCache = map[string]int{
// 		"1-0": 1,
// 		"1-1": 1,
// 	}
// )

// func solve(n int) int {
// 	for i := 1; i <= n; i++ {

// 	}
// }

// func byStartMap(n, start int) int {
// 	code := fmt.Sprintf("%d-%d", n, start)
// 	if v, ok := byStartMapCache[code]; ok {
// 		return v
// 	}

// 	if start > n {
// 		panic("AH")
// 	}

// 	// Add a single i to all arrangements of (k-i) that don't start with i
// 	v := byStartMap(n-start, 0)
// 	if start <= n-start {
// 		v = (v + mod - byStartMap(n-start, start)) % mod
// 	}

// 	// Add two start's to all arrangements of (n-start) that start with start
// 	if n-2*start > 0 && n-2*start >= start {
// 		v = (v + byStartMap(n-2*start, start)) % mod
// 	}

// 	byStartMapCache[code] = v
// 	return v
// }

//

// 1 1 4 4 10 19 33 59 113 210

// []
// [  1]
// [  1  0]
// [  4  2  1]
// [  4  2  0  1]
// [ 10  4  3  1  1]
// [ 19  8  5  3  1  1]
// [ 33 15  8  3  4  1  1]
// [ 59 26 14  9  3  4  1  1]
// [113 48 28 17  9  4  4  1  1]
// [210 91 50 31 18  9  4  4  1  1]

// []
// [  1]
// [  1  0]
// [  4  2  1]
// [  4  2  0  1]
// [ 10  4  3  1  1]
// [ 19  8  5  3  1  1]
// [ 33 15  8  3  4  1  1]
// [ 59 26 14  9  3  4  1  1]
// [113 48 28 17  9  4  4  1  1]
// [210 91 50 31 18  9  4  4  1  1]
