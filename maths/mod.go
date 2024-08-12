package maths

// SolveMod solves the following equation for x: a*x == target (mod mod)
// This logic was determined from page 4 of https://www.math.utah.edu/~fguevara/ACCESS2013/Euclid.pdf
//
// NOTE: This method assumes that `a` and `mod` are co-prime!!
func SolveMod(a, mod, target int) int {
	if target == 0 {
		return mod
	}

	if a == 1 {
		return target
	}

	r := recurSolveMod(a, mod)[1]
	if r < 0 {
		r += mod
	}
	return r * target % mod
}

func recurSolveMod(a, mod int) []int {
	factor := mod / a
	offset := mod % a

	if offset == 1 {
		return []int{1, -factor}
	}

	p := recurSolveMod(offset, a)
	return []int{
		p[1],
		p[0] - factor*p[1],
	}
}

// Coprime returns whether or not a and b are coprime.
func Coprime(a, b int) bool {
	return Gcd(a, b) == 1
}

// Gcd uses the Euclidean Algorithm to determine GCD(a, b).
// This logic was determined from page 4 of https://www.math.utah.edu/~fguevara/ACCESS2013/Euclid.pdf
func Gcd(a, b int) int {
	offset := b % a

	if offset == 0 {
		return a
	}

	return Gcd(offset, a)
}
