package eulerchallenge

import (
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P153() *problem {
	return intInputNode(153, func(o command.Output, n int) {
		p := generator.Primes()
		_ = p
		var ss int
		counts := make([]int, n/2+1, n/2+1)
		for i := 1; i <= n/2; i++ {
			if i%1_000_000 == 0 {
				fmt.Println(i)
			}
			//fmt.Println(i, generator.Factors(i, p))
			for _, k := range generator.Factors(i, p) {
				ss += k
			}
			counts[i] = ss
			//ss +=
		}

		var actualTotal int
		for i := 1; i <= n/2; i++ {
			// Number of integers with i as a factor = n/i
			actualTotal += (n / i) * i
			//fmt.Println(i, n/i, i*(n/i))
		}
		// Numbers greater than n/2 only occur once
		start := (n / 2) + 1
		end := n
		// start + (start+1) + (start+2) + ... + (end-2) + (end-1) + end
		// = start*(end-start+1) + 1 + 2 + 3 + ... + (end - start)
		// = start*(end-start+1) + (end - start)*(end-start+1)/2
		actualTotal += start*(end-start+1) + (end-start)*(end-start+1)/2
		_ = start

		fmt.Println("ss", ss)

		var iSums int
		for i := 1; i*i <= n; i++ {
			for j := i; j*j+i*i <= n; j++ {
				if generator.Coprimes(i, j, p) && (i != 1) && (j != 1) {
					continue
				}
				//hit := i*i + j*j
				val := 2 * (i + j)
				upTo := n / (j*j + i*i)
				offset := counts[upTo] * val
				if i == j {
					offset /= 2
				}
				//fmt.Println("offset", i, j, counts[upTo], upTo, offset)
				iSums += offset
			}
		}
		fmt.Println(iSums, iSums+actualTotal)

		return
		var sum int
		for i := 1; i <= n/2; i++ {
			// Number of integers with i as a factor = n/i
			sum += (n / i) * i
			//fmt.Println(i, n/i, i*(n/i))
		}
		fmt.Println(sum)
		// Numbers greater than n/2 only occur once
		start = (n / 2) + 1
		end = n
		// start + (start+1) + (start+2) + ... + (end-2) + (end-1) + end
		// = start*(end-start+1) + 1 + 2 + 3 + ... + (end - start)
		// = start*(end-start+1) + (end - start)*(end-start+1)/2
		sum += start*(end-start+1) + (end-start)*(end-start+1)/2
		_ = start

		var iSum int

		/*for i := 1; i*i <= n; i++ {
			for j := i + 1; j*j+i*i <= n; j++ {
				if generator.Coprimes(i, j, p) {
					continue
				}
				hit := j*j + i*i
				val := j + i
				//count :=
			}
		}
		/*squares := generator.SmallPowerGenerator(2)

		for iter, i := squares.Start(1); i <= n; i = iter.Next() {
			irt := maths.Sqrt(i)
			fmt.Println("irt", irt)
			for jter, j := squares.Start(i); j+i <= n; j = jter.Next() {
				jrt := maths.Sqrt(j)
				fmt.Println("jrt", jrt)
				number := i + j
				count := n / number
				if i == j {
					iSum += count * (irt + jrt)
					fmt.Println("One", irt, jrt, count, irt+jrt, count*(irt+jrt))
				} else {
					iSum += count * 2 * (irt + jrt)
					fmt.Println("Two", irt, jrt, count, 2*(irt+jrt), count*(2*irt+jrt))
				}
			}
		}*/

		fmt.Println(sum, iSum)
	})
}
