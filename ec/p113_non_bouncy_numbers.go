package eulerchallenge

import (
	"github.com/leep-frog/command/command"
)

func P113() *problem {
	return intInputNode(113, func(o command.Output, n int) {
		var sum int
		for j := 1; j <= n; j++ {
			d := getDecreasing(9, j, 0, map[int]map[int]int{})
			i := getIncreasing(1, j, map[int]map[int]int{})
			sum += d + i - 9
		}
		o.Stdoutln(sum)
	}, []*execution{
		{
			args: []string{"100"},
			want: "51161058134250",
		},
		{
			args: []string{"10"},
			want: "277032",
		},
		{
			args: []string{"6"},
			want: "12951",
		},
		{
			args: []string{"2"},
			want: "99",
		},
		{
			args: []string{"1"},
			want: "9",
		},
	})
}

func getIncreasing(min, rem int, m map[int]map[int]int) int {
	if rem == 0 {
		return 1
	}
	if min > 9 {
		return 0
	}

	if m1, ok := m[min]; ok {
		if v, ok2 := m1[rem]; ok2 {
			return v
		}
	}

	v := getIncreasing(min+1, rem, m) + getIncreasing(min, rem-1, m)
	if m[min] == nil {
		m[min] = map[int]int{}
	}
	m[min][rem] = v
	return v
}

func getDecreasing(max, rem, length int, m map[int]map[int]int) int {
	if rem == 0 {
		return 1
	}
	if max < 0 || length == 0 && max == 0 {
		return 0
	}

	if m1, ok := m[max]; ok {
		if v, ok2 := m1[rem]; ok2 {
			return v
		}
	}

	v := getDecreasing(max-1, rem, length, m) + getDecreasing(max, rem-1, length+1, m)
	if m[max] == nil {
		m[max] = map[int]int{}
	}
	m[max][rem] = v
	return v
}
