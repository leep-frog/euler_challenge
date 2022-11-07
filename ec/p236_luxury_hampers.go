package eulerchallenge

import (
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
	"golang.org/x/exp/slices"
)

func P236() *problem {
	return noInputNode(236, func(o command.Output) {
		supply := [][]int{
			{5248, 640},
			{1312, 1888},
			{2624, 3776},
			{5760, 3776},
			{3936, 5664},
			/*
				[5248 640 ] (2^7 * 41)        (2^7 * 5)
				[1312 1888] (2^5 * 41)        (2^5 * 59)
				[2624 3776] (2^6 * 41)        (2^6 * 59)
				[5760 3776] (2^7 * 3^2 * 5^1) (2^6 * 59)
				[3936 5664] (2^5 * 3 * 41)    (2^5 * 3 * 59)

				[1476 1475] (2^2 * 3^2 * 41)  (5^2 * 59)

				m = (n_b / d_b) / (n_a / d_a)
				  = (n_b * d_a) / (d_b * n_a)

			*/
		}

		g := generator.Primes()
		fmt.Println("[1476 1475]", generator.PrimeFactors(1476, g), generator.PrimeFactors(1475, g))
		for _, product := range supply {
			fmt.Println(product, generator.PrimeFactors(product[0], g), generator.PrimeFactors(product[1], g))
		}

		return

		slices.SortFunc(supply, func(a, b []int) bool {
			return a[0] < b[0]
		})

		total_a, total_b := 0, 0
		for _, product := range supply {
			total_a += product[0]
			total_b += product[1]
		}

		// n_b * m = n_a

		// d_a (d for denominator) = number of products for a
		// d_b (d for denominator) = number of products for b
		d_a, d_b := 5760, 3776

		// n_a = number of spoiled for a
		// n_b = number of spoiled for b
		for n_a := 1; n_a <= d_a; n_a++ {
			// r_a = spoilage rate for a
			// r_b = spoilage rate for b
			// spoilage rate for b should be higher:
			// n_b / d_b > n_a / d_a
			// n_b > (d_b * n_a) / d_a
			for n_b := (d_b * n_a) / d_a; n_b <= d_b; n_b++ {
				// First, verify the total can have m
				/*total_a

				cnt_a, cnt_b := 0, 0
				// m = (n_b / d_b) / (n_a / d_a)
				//   = (n_b / d_b) * (d_a / n_a)
				//   = (n_b * d_a) / (d_b * n_a)
				num := (n_b * d_a)
				den := (d_b * n_a)*/
			}
		}

		//fmt.Println(cnt)
	}, &execution{
		want: "0",
		skip: "TODO",
	})
}
