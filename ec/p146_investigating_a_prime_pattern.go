package eulerchallenge

import (
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P146() *problem {
	return intInputNode(146, func(o command.Output, n int) {
		p := generator.Primes()
		offsets := []int{1, 3, 7, 9, 13, 27}
		var sum int
		max := 1_000_000
		// Can't be divisble by
		// Don't be divisble by 3
		mod3 := 1
		for i := 10; i < 150_000_000; i, mod3 = i+10, (mod3+1)%3 {
			if mod3 == 0 {
				continue
			}
			if i > max {
				fmt.Println(max)
				max += 1_000_000
			}
			sq := i * i
			valid := true
			for _, o := range offsets {
				if !generator.IsPrime(sq+o, p) {
					valid = false
					break
				}
			}
			if valid {
				fmt.Println(i)
				sum += i
			}
		}
		o.Stdoutln(sum)
	})
}
