package p164

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func brute164(n int) int {
	cnt := 0
	max := maths.Pow(10, n)
	for i := maths.Pow(10, n-1); i < max; i++ {
		ds := maths.Digits(i)
		valid := true
		for j := 0; j < len(ds); j++ {
			end := maths.Min(j+3, len(ds))
			if bread.Sum(ds[j:end]) > 9 {
				valid = false
				break
			}
		}
		if valid {
			cnt++
		}
	}
	return cnt
}

func elegant164(n int) int {
	if n == 1 {
		return 9
	}
	// Keep track of number of valid ways we can make a number that starts with ab
	// This two dimensional slice stores:
	// a -> b -> number of valid ways we can make a number that starts with ab
	var startCounts [][]int
	for a := 0; a <= 9; a++ {
		var row []int
		for b := 0; b <= 9; b++ {
			if a+b <= 9 {
				row = append(row, 1)
			} else {
				row = append(row, 0)
			}
		}
		startCounts = append(startCounts, row)
	}

	for i := 2; i < n; i++ {
		// Create an empty newStartCount
		var newStartCount [][]int
		for j := 0; j <= 9; j++ {
			newStartCount = append(newStartCount, make([]int, 10, 10))
		}

		// Shift existing startCount to be bc and see what a's work
		for a := 0; a <= 9; a++ {
			for b, row := range startCounts {
				for c, cnt := range row {
					if a+b+c <= 9 {
						newStartCount[a][b] += cnt
					}
				}
			}
		}
		startCounts = newStartCount
	}

	sum := 0
	for i, row := range startCounts {
		if i != 0 {
			sum += bread.Sum(row)
		}
	}
	return sum
}

func P164() *ecmodels.Problem {
	return ecmodels.IntInputNode(164, func(o command.Output, n int) {
		//o.Stdoutln(brute164(n))
		o.Stdoutln(elegant164(n))
	}, []*ecmodels.Execution{
		{
			Args: []string{"1"},
			Want: "9",
		},
		{
			Args: []string{"2"},
			Want: "45",
		},
		{
			Args: []string{"3"},
			Want: "165",
		},
		{
			Args: []string{"4"},
			Want: "990",
		},
		{
			Args: []string{"5"},
			Want: "5445",
		},
		{
			Args: []string{"20"},
			Want: "378158756814587",
		},
	})
}
