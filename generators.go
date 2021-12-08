package main

type Generator struct {
	values []int

	next func(*Generator) int
}

func (g *Generator) Last() int {
	return g.values[len(g.values)-1]
}

func (g *Generator) Nth(i int) int {
	for len(g.values) <= i {
		g.Next()
	}
	return g.values[i]
}

func (g *Generator) Next() int {
	i := g.next(g)
	g.values = append(g.values, i)
	return i
}

func NewGenerator(start int, f func(*Generator) int) *Generator {
	return &Generator{
		values: []int{start},
		next: func(g *Generator) int {
			if len(g.values) == 0 {
				return start
			}
			return f(g)
		},
	}
}

func Primer() *Generator {
	return NewGenerator(2, func(g *Generator) int {
		for i := g.Last() + 1; ; i++ {
			newPrime := true
			for _, p := range g.values {
				if i%p == 0 {
					newPrime = false
					break
				}
			}
			if newPrime {
				return i
			}
		}
	})
}

func Fibonaccis() *Generator {
	a := 1
	b := 2
	return NewGenerator(1, func(g *Generator) int {
		r := b
		b = a + b
		a = r
		return a
	})
}

func Triangulars() *Generator {
	i := 0
	return NewGenerator(1, func(g *Generator) int {
		i++
		return g.Last() + i
	})
}
