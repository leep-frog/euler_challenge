package eulerchallenge

import (
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/fraction"
	"github.com/leep-frog/euler_challenge/generator"
	"golang.org/x/exp/slices"
)

type ctx236 struct {
	m      *fraction.Fraction
	totalA int
	totalB int
	primes *generator.Generator[int]
	checks int
	path   [][]int
}

func check236(depth int, ctx *ctx236, spoiledA, spoiledB int, supply [][]int) bool {
	ctx.checks++
	if ctx.checks > 100_000 {
		return false
	}
	if len(supply) == 0 {
		// true if spoiled rate of a is worse:
		// (spoiled_a / total_a) / (spoiled_b / total_b) = m_n / m_d
		// m_d * (spoiled_a / total_a) = m_n * (spoiled_b / total_b)
		// m_d * spoiled_a * total_b = m_n * spoiled_b * total_a
		if ctx.m.D*spoiledA*ctx.totalB == ctx.m.N*spoiledB*ctx.totalA {
			fmt.Println(append(ctx.path, []int{spoiledA, spoiledB}))
			return true
		}
		return false
	}

	// m = m_n / m_d = (n_b * d_a) / (d_b * n_a)
	// n_a and n_b must be integers:
	// m_n / m_d = (n_b * d_a) / (d_b * n_a)
	// (d_b * n_a * m_n) / (m_d * d_a) = n_b
	d_a, d_b := supply[0][0], supply[0][1]
	f := fraction.Simplify(d_b*ctx.m.N, ctx.m.D*d_a, ctx.primes)
	for n_a := f.D; n_a <= d_a; n_a += f.D {
		n_b := (d_b * n_a * ctx.m.N) / (ctx.m.D * d_a)
		/*if depth == 0 {
			fmt.Println("D", ctx.m, fraction.New(n_a, d_a), fraction.New(n_b, d_b))
		}*/

		ctx.path = append(ctx.path, []int{n_a, n_b})
		if check236(depth+1, ctx, spoiledA+n_a, spoiledB+n_b, supply[1:]) {
			return true
		}
		ctx.path = ctx.path[:(len(ctx.path) - 1)]
	}
	return false
}

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

				m = m_n / m_d > 1
				m_n > m_d

				m = spoilage_b / spoilage_a
				  = (n_b / d_b) / (n_a / d_a)

				m_n / m_d = (n_b / d_b) / (n_a / d_a)

				// Assume we know all values, except n_b and n_a

				(m_n * d_a) / (m_d * n_a) = (n_b / d_b)
				(m_n * d_a * d_b) / (m_d * n_a) = n_b

				// n_b must be an integer so
				// (m_n * d_a * d_b) % (m_d * n_a) == 0
				// (m_n * d_a * d_b) / (m_d * n_a) = k

			*/
		}

		slices.SortFunc(supply, func(a, b []int) bool {
			return a[0] < b[0]
		})

		primes := generator.Primes()

		totalA, totalB := 0, 0
		for _, s := range supply {
			totalA += s[0]
			totalB += s[1]
			fmt.Println(s, generator.PrimeFactors(s[0], primes), generator.PrimeFactors(s[1], primes))
		}
		fmt.Println(totalA, totalB, generator.PrimeFactors(totalA, primes), generator.PrimeFactors(totalB, primes))
		fmt.Println("M", generator.PrimeFactors(1476, primes), generator.PrimeFactors(1475, primes))

		checked := map[string]bool{}
		valid := map[string]bool{}

		d_a := supply[0][0]
		d_b := supply[0][1]
		best := fraction.New(1476, 1475)
		for n_a := 1; n_a <= d_a; n_a++ {
			fmt.Println("NA", n_a)
			// m = m_n / m_d > 1
			// m = spoilage_b / spoilage_a = (n_b / d_b) / (n_a / d_a)
			//   (n_b / d_b) / (n_a / d_a) > 1
			//   (n_b / d_b) > (n_a / d_a)
			//   n_b > (n_a * d_b) / d_a
			start := ((n_a * d_b) / d_a) + 1
			if (n_a*d_b)%d_a == 0 {
				start = n_a * d_b / d_a
			}

			for n_b := start; n_b <= d_b; n_b++ {
				m := fraction.Simplify(n_b*d_a, d_b*n_a, primes)
				if checked[m.String()] {
					continue
				}
				checked[m.String()] = true
				if !best.LT(m) {
					continue
				}
				// m = (n_b / d_b) / (n_a / d_a)
				//   = (n_b / d_b) * (d_a / n_a)
				//   = (n_b * d_a) / (d_b * n_a)

				// m = (s_a / t_a) / (s_b / t_b)
				//   = (s_a / t_a) * (t_b / s_b)
				//   = (s_a * t_b) / (t_a * s_b)

				ctx := &ctx236{m, totalA, totalB, primes, 0, nil}

				// Ensure all supplies can divide it
				canDivide := true
				for _, s := range supply[1:] {
					sd_a, sd_b := s[0], s[1]
					// Smallest value of n_a
					f := fraction.Simplify(sd_b*ctx.m.N, ctx.m.D*sd_a, ctx.primes)
					smallest_sn_a := f.D
					if smallest_sn_a > sd_a {
						canDivide = false
						break
					}
					smallest_sn_b := (sd_b * smallest_sn_a * ctx.m.N) / (ctx.m.D * sd_a)
					if smallest_sn_b > sd_b {
						canDivide = false
						break
					}
				}

				//fmt.Println("CD", canDivide)
				//fmt.Println("TM", m)

				ctx.path = append(ctx.path, []int{n_a, n_b})
				if canDivide && check236(0, ctx, n_a, n_b, supply[1:]) {
					fmt.Println("VALID", m)
					valid[m.String()] = true
					best = m
				}
			}
		}

		fmt.Println(valid)
		return

		g := generator.Primes()
		fmt.Println("[1476 1475]", generator.PrimeFactors(1476, g), generator.PrimeFactors(1475, g))
		for _, product := range supply {
			fmt.Println(product, generator.PrimeFactors(product[0], g), generator.PrimeFactors(product[1], g))
		}

		return

		total_a, total_b := 0, 0
		for _, product := range supply {
			total_a += product[0]
			total_b += product[1]
		}

		// n_b * m = n_a

		// d_a (d for denominator) = number of products for a
		// d_b (d for denominator) = number of products for b
		//d_a, d_b := 5760, 3776

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
		want: "123/59",
		skip: "initial approach printed out a few values. I tried the value above and it worked!",
	})
}
