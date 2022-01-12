package eulerchallenge

import (
	"fmt"
	"sort"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P66() *command.Node {
	return command.SerialNodes(
		command.Description("https://projecteuler.net/problem=66"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			n := d.Int(N)

			best := maths.Largest()
			p := generator.Primes()
			_ = p
			need := map[int]bool{}
			for D := 2; D <= n; D++ {
				if !maths.IsSquare(D) {
					need[D] = true
				}
			}

			sols := map[int][]int{}

			/*for x := 2; len(need) > 0; x++ {
				got := map[int]bool{}
				for D := range need {
					// x*x - D*y*y = 1
					// y = SQRT[(x*x - 1)/D]
					y2 := (x*x - 1) / D
					if x*x-D*y2 == 1 && maths.IsSquare(y2) {
						fmt.Println(D, x)
						best.IndexCheck(D, x)
						got[D] = true
					}
				}

				for D := range got {
					delete(need, D)
				}
				if len(got) > 0 {
					fmt.Println("NEED:", len(need), "| X:", x)
				}
			}*/
			for x := 2; len(need) > 1; x++ {
				got := map[int]bool{}
				for D := range need {
					// x*x - D*y*y = 1
					// y = SQRT[(x*x - 1)/D]
					y2 := (x*x - 1) / D
					if x*x-D*y2 == 1 && maths.IsSquare(y2) {
						y := maths.Sqrt(y2)
						best.IndexCheck(D, x)
						got[D] = true
						o.Stdoutf("Got D: %d^2 - %d * %d^2", x, D, y)
						sols[D] = []int{x, y}
						factors := generator.PrimeFactors(D, p)
						var parts []int
						for k, v := range factors {
							for ; v >= 2; v -= 2 {
								parts = append(parts, k*k)
							}
						}
						for _, set := range maths.Sets(parts) {
							d := D
							for _, s := range set {
								d /= s
							}
							if got[d] {
								o.Stdoutln("huzzah: ", D, d)
								delete(got, d)
							}
						}
						o.Stdoutln("PARTS: ", parts)
					}
				}

				for D := range got {
					delete(need, D)
				}
				if len(got) > 0 {
					fmt.Println("NEED:", len(need), "| X:", x)
				}
			}

			if len(need) == 1 {
				o.Stdoutln(need)
			} else {
				o.Stdoutln(best.BestIndex(), best.Best())
			}

			/*for y := 2; len(need) > 0; y++ {
				got := map[int]bool{}
				for D := range need {
					// x*x - D*y*y = 1
					// x = SQRT[1 + D*y*y]
					x2 := 1 + D*y*y
					if x2-D*y*y == 1 && maths.IsSquare(x2) {
						fmt.Println(D, x2)
						best.IndexCheck(D, maths.Sqrt(x2))
						got[D] = true
						o.Stdoutln("Got D", D)
						sols[D] = []int{maths.Sqrt(x2), y}
						factors := generator.PrimeFactors(D, p)
						var parts []int
						for k, v := range factors {
							for ; v >= 2; v -= 2 {
								parts = append(parts, k*k)
							}
						}
						for _, set := range maths.Sets(parts) {
							d := D
							for _, s := range set {
								d /= s
							}
							o.Stdoutln("Subset: ", D, d)
							if got[d] {
								o.Stdoutln("huzzah: ", D)
								delete(got, d)
								return
							}
						}
						o.Stdoutln("PARTS: ", parts)
					}
				}

				for D := range got {
					delete(need, D)
				}
				if len(got) > 0 {
					fmt.Println("NEED:", len(need), "| Y:", y)
				}
			}*/

			var keys []int
			for k := range sols {
				keys = append(keys, k)
			}
			sort.Ints(keys)
			for _, k := range keys {
				v := sols[k]
				o.Stdoutf("%d^2 - %d * %d^2 = 1", v[0], k, v[1])
			}
			/*for D := 2; D <= n; D++ {
				if maths.IsSquare(D) {
					continue
				}

				for x := 2; ; x++ {
					// x*x - D*y*y = 1
					// y = SQRT[(x*x - 1)/D]
					y2 := (x*x - 1) / D
					if x*x-D*y2 == 1 && maths.IsSquare(y2) {
						fmt.Println(D, y2)
						best.IndexCheck(D, x)
						goto found
					}
				}
			found:
			}*/
			o.Stdoutln(best.Best(), best.BestIndex())
		}),
	)
}

/*
x*x - D*y*y = 1
D = (x*x - 1)/(y*y)

/*func altR(n int) {
	p := generator.Primes()
	need := map[int]bool{}
	for D := 2; D <= n; D++ {
		if !maths.IsSquare(D) {
			need[D] = true
		}
	}
	for x := 2; ; x++ {
		factors := generator.PrimeFactors(x*x - 1)
	}
}
*/
