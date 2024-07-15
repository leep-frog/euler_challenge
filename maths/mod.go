package maths

func SolveMod(a, mod, target int) int {

	if target == 0 {
		return mod
	}

	if a == 1 {
		return target
	}

	r := recurSolveMod(a, mod, "")[1]
	if r < 0 {
		r += mod
	}
	return r * target % mod
}

func recurSolveMod(a, mod int, indent string) []int {

	factor := mod / a
	offset := mod % a

	if offset == 1 {
		return []int{1, -factor}
	}

	c := []int{1, -factor}
	p := recurSolveMod(offset, a, indent+"  ")

	r := []int{
		c[0] * p[1],
		p[0] + c[1]*p[1],
	}

	return r
}
