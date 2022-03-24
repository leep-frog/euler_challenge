package eulerchallenge

import (
	"fmt"
	"math"
	"math/big"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P137() *problem {
	return intInputNode(137, func(o command.Output, n int) {
		f := generator.Fibonaccis()
		g := generator.Primes()
		for i := 0; i < 25; i++ {
			fmt.Printf("%d, ", f.Nth(2*i+1)+f.Nth(2*i+2))
		}
		fmt.Println()
		return

		/*a, b := big.NewRat(1, 1), big.NewRat(1, 1)
		fibs := []*big.Rat{a}
		for i := 0; i < 500; i++ {

		}*/

		for den := int64(5); den < 6; den++ {
			for num := int64(3); num <= 4; num++ {
				//for num := (den - 1) / 2; num <= (den+1)/2; num++ {
				if generator.Coprimes(int(num), int(den), g) {
					continue
				}
				r := big.NewRat(0, 1)
				prod := big.NewRat(1, 1)
				a, b := big.NewRat(1, 1), big.NewRat(1, 1)
				for i := 0; i < 1000; i++ {
					//fmt.Println(i, f.Nth(i))
					prod.Mul(prod, big.NewRat(num, den))
					add := big.NewRat(0, 1).Mul(prod, a)
					r = big.NewRat(0, 1).Add(r, add)
					c := big.NewRat(1, 1).Add(a, b)
					a, b = b, c
				}
				flt, _ := r.Float64()
				if math.Abs(float64(int(flt))-flt) < 0.000001 {
					fmt.Println(num, den, flt)
				}
			}
		}
		o.Stdoutln(n)
	})
}
