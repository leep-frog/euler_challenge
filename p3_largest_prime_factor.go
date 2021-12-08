package main

import (
	"github.com/leep-frog/command"
)

/*type Generator struct {
	values []int

	next func()
}*/

type Primer struct {
	Primes []int
}

func (p *Primer) Last() int {
	return p.Primes[len(p.Primes)-1]
}

func (p *Primer) Nth(i int) int {
	for len(p.Primes) <= i {
		p.Next()
	}
	return p.Primes[i]
}

func (p *Primer) Next() int {
	if len(p.Primes) == 0 {
		p.Primes = append(p.Primes, 2)
		return 2
	}

	for i := p.Primes[len(p.Primes)-1] + 1; ; i++ {
		newPrime := true
		for _, p := range p.Primes {
			if i%p == 0 {
				newPrime = false
				break
			}
		}
		if newPrime {
			p.Primes = append(p.Primes, i)
			return i
		}
	}
}

// TODO: move all of these to helper directory
func getPrimeFactors(n int, p *Primer) map[int]int {
	r := map[int]int{}
	for i := 0; ; i++ {
		pi := p.Nth(i)
		for n%pi == 0 {
			r[pi]++
			n = n / pi
			if n == 1 {
				return r
			}
		}
	}
}

func p3() *command.Node {
	return command.SerialNodes(
		command.Description("Find the largest prime factor of N"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) error {
			factors := getPrimeFactors(d.Int(N), &Primer{})

			max := 0
			for f := range factors {
				if f > max {
					max = f
				}
			}

			o.Stdoutf("%d", max)
			return nil
		}),
	)
}
