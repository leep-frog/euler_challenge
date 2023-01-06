// functional implements functional programming utilities.
package functional

// Any returns true if f(t) == true for any t in ts.
func Any[T any](ts []T, f func(t T) bool) bool {
	for _, t := range ts {
		if f(t) {
			return true
		}
	}
	return false
}

// All Returns true if f(t) == true for all t in ts.
func All[T any](ts []T, f func(t T) bool) bool {
	for _, t := range ts {
		if !f(t) {
			return false
		}
	}
	return true
}

// None returns true if f(t) == false for all t in ts.
func None[T any](ts []T, f func(t T) bool) bool {
	for _, t := range ts {
		if f(t) {
			return false
		}
	}
	return true
}

// MapWithIndex maps all of the elements in `items` with function f.
func MapWithIndex[I, O any](items []I, f func(int, I) O) []O {
	var r []O
	for idx, item := range items {
		r = append(r, f(idx, item))
	}
	return r
}

// Map maps all of the elements in `items` with the function f.
func Map[I, O any](items []I, f func(I) O) []O {
	return MapWithIndex(items, func(idx int, i I) O {
		return f(i)
	})
}

// Reduce reduces `base` across `items` with function f.
func Reduce[B, T any](base B, items []T, f func(B, T) B) B {
	b := base
	for _, t := range items {
		b = f(b, t)
	}
	return b
}
