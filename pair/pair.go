package pair

type Pair[T any] struct {
	A T
	B T
}

func New[T any](a, b T) *Pair[T] {
	return &Pair[T]{a, b}
}
