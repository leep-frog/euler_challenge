package series

import "github.com/leep-frog/euler_challenge/maths"

// Series is a type that represents a recursive series (i.e. a series where S(n)
// is dependent on S(n-1), ..., S(n-k)).
type Series[T any] struct {
	// at is the index of the series this is currently at
	at int
	// elements are the last k elements of the series, where
	// elements[0] is Series(at), elements[1] is Series(at-1), etc.
	elements []T

	// generateNext generates the next element in the series
	// where nOffset(-k) returns the current series element minus k
	generateNext func(n int, nOffset func(int) T) T
}

// New generates a new Series with the provided starting elements.
// generateNext is a function that generates the next element in the series
// where nOffset(-k) returns the current series element minus k
func New[T any](elements []T, generateNext func(n int, nOffset func(int) T) T) *Series[T] {
	return &Series[T]{len(elements) - 1, elements, generateNext}
}

func (s *Series[T]) Get(k int) T {
	for s.at <= k {
		s.iterate()
	}
	return s.elements[s.at-k]
}

func (s *Series[T]) iterate() {
	next := s.generateNext(s.at+1, s.nOffset)
	for i := 0; i < len(s.elements)-1; i++ {
		s.elements[i+1] = s.elements[i]
	}
	s.elements[0] = next
	s.at++
}

// nOffset is only to be used by the iterate method
func (s *Series[T]) nOffset(k int) T {
	return s.elements[-(k + 1)]
}

// https://oeis.org/A055244
func A055244() *Series[*maths.Int] {
	return New(
		[]*maths.Int{maths.One(), maths.One()},
		func(n int, nOffset func(int) *maths.Int) *maths.Int {
			a := maths.NewInt(n - 4).TimesInt(n).MinusInt(6).Times(nOffset(-2))
			b := maths.NewInt(n - 5).TimesInt(n).MinusInt(11).Times(nOffset(-1))
			c := maths.NewInt(n - 6).TimesInt(n).MinusInt(1)
			return a.Plus(b).Div(c)
		},
	)
}
