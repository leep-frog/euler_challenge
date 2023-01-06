package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/combinatorics"
	"github.com/leep-frog/euler_challenge/maths"
)

type p93op int

func (op p93op) apply(a, b float64) (float64, bool) {
	switch op {
	case add:
		return a + b, true
	case subtract:
		return a - b, true
	case reverseSubtract:
		return b - a, true
	case multiply:
		return a * b, true
	case divide:
		if b == 0 {
			return 0, false
		}
		return a / b, true
	case reverseDivide:
		if a == 0 {
			return 0, false
		}
		return b / a, true
	}
	panic("unknown op")
}

const (
	add p93op = iota
	subtract
	reverseSubtract
	divide
	reverseDivide
	multiply
)

var (
	p93ops = []p93op{
		add,
		subtract,
		reverseSubtract,
		divide,
		multiply,
	}
)

func P93() *problem {
	return intInputNode(93, func(o command.Output, n int) {
		opSets := combinatorics.GenerateCombos(&combinatorics.Combinatorics[p93op]{
			Parts:            p93ops,
			MinLength:        3,
			MaxLength:        3,
			AllowReplacement: true,
			OrderMatters:     true,
		})
		best := maths.Largest[[]float64, int]()
		for d := 0.0; d <= float64(n); d++ {
			for c := d - 1; c >= 0; c-- {
				for b := c - 1; b >= 0; b-- {
					for a := b - 1; a >= 0; a-- {
						values := map[int]bool{}
						for _, order := range combinatorics.Permutations([]float64{a, b, c, d}) {
							m := maths.NewSimpleSet(order...)
							bad := false
							for k := range m {
								if m[-k] {
									bad = true
									break
								}
							}
							if bad {
								continue
							}
							for _, opSet := range opSets {
								value := order[0]
								valid := true
								for i, op := range opSet {
									value, valid = op.apply(value, order[i+1])
									if !valid {
										break
									}
								}
								if valid && float64(int(value)) == value {
									values[int(value)] = true
								}

								// Then, see if combining pairs of ops
								value, valid = opSet[0].apply(order[0], order[1])
								if !valid {
									continue
								}
								value2, valid2 := opSet[1].apply(order[2], order[3])
								if !valid2 {
									continue
								}
								value, valid = opSet[2].apply(value, value2)
								if !valid || float64(int(value)) != value {
									continue
								}
								values[int(value)] = true
							}
						}
						k := 1
						for ; values[k] || values[-k]; k++ {
						}
						best.IndexCheck([]float64{a, b, c, d}, k-1)
					}
				}
			}
		}
		o.Stdoutln(maths.Join(best.BestIndex(), ""), best.Best())
	}, []*execution{
		{
			args: []string{"9"},
			want: "1258 51",
		},
		{
			args: []string{"4"},
			want: "1234 28",
		},
	})
}
