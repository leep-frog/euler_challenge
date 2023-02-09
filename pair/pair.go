package pair

type Pair[A, B any] struct {
	A A
	B B
}

func New[T any](a, b T) *Pair[T, T] {
	return &Pair[T, T]{a, b}
}

func NewDiff[A, B any](a A, b B) *Pair[A, B] {
	return &Pair[A, B]{a, b}
}

func Zip[A comparable, B any](m map[A]B) []*Pair[A, B] {
	var r []*Pair[A, B]
	for k, v := range m {
		r = append(r, NewDiff(k, v))
	}
	return r
}
