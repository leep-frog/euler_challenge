package main

import (
	"fmt"
	"testing"

	"github.com/leep-frog/command"
)

var (
	runLongTests = false
)

func TestAll(t *testing.T) {
	for _, test := range []struct {
		name    string
		problem int
		arg     int
		want    int
		long    bool
	}{
		{
			name:    "p1 example",
			problem: 1,
			arg:     10,
			want:    23,
		},
		{
			name:    "p1",
			problem: 1,
			arg:     1000,
			want:    233168,
		},
		{
			name:    "p2",
			problem: 2,
			arg:     4000000,
			want:    4613732,
		},
		{
			name:    "p3 example",
			problem: 3,
			arg:     13195,
			want:    29,
		},
		{
			name:    "p3",
			problem: 3,
			arg:     600851475143,
			want:    6857,
		},
		{
			name:    "p4 example",
			problem: 4,
			arg:     2,
			want:    9009,
		},
		{
			name:    "p4",
			problem: 4,
			arg:     3,
			want:    906609,
		},
		{
			name:    "p5 example",
			problem: 5,
			arg:     10,
			want:    2520,
		},
		{
			name:    "p5",
			problem: 5,
			arg:     20,
			want:    232792560,
		},
		{
			name:    "p6 example",
			problem: 6,
			arg:     10,
			want:    2640,
		},
		{
			name:    "p6",
			problem: 6,
			arg:     100,
			want:    25164150,
		},
		{
			name:    "p7 example",
			problem: 7,
			arg:     6,
			want:    13,
		},
		{
			name:    "p7",
			problem: 7,
			arg:     10001,
			want:    104743,
		},
		{
			name:    "p8 example",
			problem: 8,
			arg:     4,
			want:    5832,
		},
		{
			name:    "p8",
			problem: 8,
			arg:     13,
			want:    23514624000,
		},
		{
			name:    "p9",
			problem: 9,
			arg:     1000,
			want:    31875000,
		},
		{
			name:    "p10 example",
			problem: 10,
			arg:     10,
			want:    17,
		},
		{
			name:    "p10",
			problem: 10,
			arg:     2000000,
			want:    142913828922,
			long:    true,
		},
		{
			name:    "p11",
			problem: 11,
			arg:     4,
			want:    70600674,
			long:    true,
		},
		{
			name:    "p12 example",
			problem: 12,
			arg:     5,
			want:    28,
		},
		{
			name:    "p12",
			problem: 12,
			arg:     500,
			want:    1,
		},
	} {
		if test.long == runLongTests {
			t.Run(test.name, func(t *testing.T) {
				etc := &command.ExecuteTestCase{
					Node:          node(),
					Args:          []string{itos(test.problem), itos(test.arg)},
					WantStdout:    []string{itos(test.want)},
					SkipDataCheck: true,
				}
				command.ExecuteTest(t, etc)
			})
		}
	}
}

func itos(i int) string {
	return fmt.Sprintf("%d", i)
}
