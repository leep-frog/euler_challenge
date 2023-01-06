package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/maths"
)

func generate171(remaining, min, squareSum int, cur []int, all *[][]int) {
	if remaining == 0 {
		if maths.IsSquare(squareSum) {
			*all = append(*all, bread.Copy(cur))
		}
		return
	}

	for j := min; j <= 9; j++ {
		generate171(remaining-1, j, squareSum+j*j, append(cur, j), all)
	}
}

func brute171(all [][]int) *maths.Int {
	sum := maths.Zero()

	for _, numbers := range all {
		perms := maths.Permutations(bread.Copy(numbers))
		for _, perm := range perms {
			sum = sum.Plus(maths.IntFromDigits(perm))
		}
	}
	return sum
}

// Fun fact: each number will be in each digit spot the same number of times!
// Even with zeros included:
// 44200 ==> 16 + 16 + 4 = 36 = 6^2
// 00244
// 04240
// etc.
func P171() *problem {
	return intInputNode(171, func(o command.Output, n int) {

		numDigits := 9

		var all [][]int
		generate171(n, 0, 0, nil, &all)

		// For every number, each digit will appear in each spot the same number of times.
		// Let digitSum be the sum of every number in a single spot, then the solution is:
		// digitSum + 10*digitSum + 100*digitSum + ...
		digitSum := maths.Zero()
		for _, numbers := range all {
			checked := map[int]bool{}
			for i, digit := range numbers {
				if checked[digit] || digit == 0 {
					continue
				}
				checked[digit] = true

				parts := append(bread.Copy(numbers[:i]), bread.Copy(numbers[i+1:])...)
				cnt := maths.PermutationCount(parts)
				digitSum = digitSum.Plus(cnt.TimesInt(digit))
			}
			digitSum = digitSum.TrimDigits(numDigits)
		}

		sum := maths.Zero()
		for i := 0; i < n; i++ {
			sum = sum.Plus(digitSum)
			digitSum = digitSum.TimesInt(10)
		}
		o.Stdoutln(sum.TrimDigits(numDigits))
	}, []*execution{
		{
			args: []string{"2"},
			want: "726",
		},
		{
			args: []string{"3"},
			want: "28083",
		},
		{
			args:     []string{"20"},
			want:     "142989277",
			estimate: 25,
		},
	})
}
