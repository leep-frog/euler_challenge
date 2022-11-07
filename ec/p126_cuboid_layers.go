package eulerchallenge

import (
	"sort"

	"github.com/leep-frog/command"
)

func P126() *problem {
	return intInputNode(126, func(o command.Output, n int) {
		maxValue := 20 * n
		C := map[int]int{}
		for a := 1; a*4 < maxValue; a++ {
			for b := a; 2*(a+b+a*b) < maxValue; b++ {
				for c := b; 2*(a*b+a*c+b*c) < maxValue; c++ {
					layer := 2 * (a*b + b*c + a*c)
					jump := 4 * (a + b + c)
					for layer < maxValue {
						C[layer]++
						layer += jump
						jump += 8
					}
				}
			}
		}

		var pairs [][]int
		for k, v := range C {
			pairs = append(pairs, []int{k, v})
		}
		sort.SliceStable(pairs, func(i, j int) bool { return pairs[i][0] < pairs[j][0] })
		for _, v := range pairs {
			if v[1] == n {
				o.Stdoutln(v[0])
				return
			}
		}
		o.Stderr("didn't find solution")
		/* Below was used to determine pattern.
		// Pattern is second order growth (difference of differences) is 8:
		// (1 x 1 x 1): [6 18 38 66 102 146 198 258]
		// (1 x 1 x 2): 10 26 50 82 122 170 226]
		// (2 x 3 x 4): [52 88 132 184 244]
		// For (1x1x1):
		//   18 -  6 = 12
		//   38 - 18 = 20
		//   66 - 38 = 28
		// Start = 2*(a*b + b*c + a*c)
		// Initial jump = 4*(a + b + c)

		maxValue := 20 * n
		C := map[int]int{}
		for a := 1; a*4 < maxValue; a++ {
			fmt.Println(a)
			for b := a; 2*(a+b+a*b) < maxValue; b++ {
				for c := b; 2*(a*b+a*c+b*c) < maxValue; c++ {

					var counts []int
					//count := 2 * (a*b + a*c + b*c)
					set := maths.NewSet[*point.Point]()

					for x := 1; x <= a; x++ {
						for y := 1; y <= b; y++ {
							for z := 1; z <= c; z++ {
								set.Add(point.NewPoint(x, y, z))
							}
						}
					}

					exposedSet := maths.NewSet[*point.Point]()
					set.For(func(p *point.Point) bool {
						exposedSet.Add(p)
						return false
					})

					//C[set.Len()]++
					//counts = append(counts, set.Len())
					for exposedSet.Len() < maxValue {
						newExposed := maths.NewSet[*point.Point]()
						exposedSet.For(func(p *point.Point) bool {
							neighbors := []*point.Point{
								point.NewPoint(p.X+1, p.Y, p.Z),
								point.NewPoint(p.X-1, p.Y, p.Z),
								point.NewPoint(p.X, p.Y-1, p.Z),
								point.NewPoint(p.X, p.Y+1, p.Z),
								point.NewPoint(p.X, p.Y, p.Z+1),
								point.NewPoint(p.X, p.Y, p.Z-1),
							}
							for _, neighbor := range neighbors {
								if !set.Contains(neighbor) {
									newExposed.Add(neighbor)
								}
							}
							return false
						})

						newExposed.For(func(p *point.Point) bool {
							set.Add(p)
							return false
						})
						exposedSet = newExposed
						C[exposedSet.Len()]++
						counts = append(counts, exposedSet.Len())
					}
					//fmt.Println(a, b, c, counts)
				}
			}
		}

		var pairs [][]int
		for k, v := range C {
			pairs = append(pairs, []int{k, v})
		}
		sort.SliceStable(pairs, func(i, j int) bool { return pairs[i][0] < pairs[j][0] })
		for _, v := range pairs {
			if v[1] == n {
				o.Stdoutln(v[0])
				return
			}
		}
		o.Stderr("didn't find solution")*/
	}, []*execution{
		{
			args: []string{"1000"},
			want: "18522",
		},
		{
			args: []string{"10"},
			want: "154",
		},
	})
}

type layeredCuboid struct {
}
