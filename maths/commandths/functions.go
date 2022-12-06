package commandths

type Function interface {
	NumArguments() int
	Execute([]int)
}

type SingleArgFunction struct {
	f func(int) int
}

type DoubleArgFunction struct {
	f func(int, int) int
}
