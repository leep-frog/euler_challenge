package p207

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
)

func P207() *ecmodels.Problem {
	return ecmodels.NoInputNode(207, func(o command.Output) {
		// 4^t = 2^t + k
		// 4^t - 2^t = k
		// 2^t * 2^t - 2^t = k
		// 2^t * (2^t - 1) = k
		//
		// Iterate over x and plugin values to x * (x - 1) to get numbers
		// Perfect solutions are when x = 2^t and where t is an integer
		// (so when x is a power of 2)
		//
		// Given the above, we just need to find the number, z where (# of powers of 2 / z) < 1/12345
		// powersOfTwo/z < 1/12345
		// powersOfTwo * 12345 < z
		// floor(log_2(z)) * 12345 < z

		// o.Stdoutln(brute())
		o.Stdoutln(elegant())

	}, &ecmodels.Execution{
		Want: "44043947822",
	})
}

func brute() int {
	nextTwoPow := 2

	var powersOfTwo, count int
	for z := 2; ; z++ {
		count++
		if z == nextTwoPow {
			powersOfTwo++
			nextTwoPow *= 2
		}

		// powersOfTwo/count < 1/12345
		// powersOfTwo*12345 < count
		if powersOfTwo*12345 < count {
			return z * (z - 1)
		}
	}
}

func elegant() int {
	// First, find the power of two that drops below the threshold
	// We know the smallest numbers will always be at x=(2^k - 1) (since increments
	// happen at powers of 2). And since 1 isn't included (so x=2 -> count of 1),
	// then we also need to subtract an additional -1 (hence the -2 below)
	powerOfTwo := 1
	for v := 2; powerOfTwo*12345 > v*2-2; v, powerOfTwo = v*2, powerOfTwo+1 {

	}

	// Now that we know the value for powersOfTwo, we can simply solve for z:
	// powersOfTwo/z < 1/12345
	// powersOfTwo * 12345 < z
	// z > powersOfTwo * 12345
	// From the problem, we want the smallest value of z that satisfies the above equation (+1)
	// Additionally, since index counting starts at 2 we need to add an additional (+1)
	z := powerOfTwo*12345 + 2

	return z * (z - 1)
}
