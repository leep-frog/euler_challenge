package p743

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

const (
	mod = 1_000_000_007
)

func P743() *ecmodels.Problem {
	return ecmodels.NoInputWithExampleNode(743, func(o command.Output, ex bool) {
		// o.Stdoutln(n)
		if ex {
			o.Stdoutln(a2(3, 9), a2(4, 20))
		} else {
			o.Stdoutln(a2(maths.Pow(10, 8), maths.Pow(10, 16)))
		}
	}, []*ecmodels.Execution{
		{
			Args: []string{"-x"},
			Want: "560 1060870",
		},
		{
			Want:     "259158998",
			Estimate: 150,
		},
	})
}

func a(k, n int) int {
	var sum int

	for numOnes := k; numOnes >= 0; numOnes -= 2 {
		fmt.Println(k)
		numTwosAndZeros := k - numOnes

		// Number of ways they can be arranged
		// Choose positions of 1s, * choose position of 2s
		// (k choose numOnes) * (numOnes choose numTwos)
		coef := maths.Choose(k, numOnes).Times(maths.Choose(k-numOnes, numTwosAndZeros/2))

		v := coef.Times(maths.BigPow(2, numOnes*(n/k)))

		// numOnes can be in any order
		sum = (sum + v.ModInt(mod)) % mod
	}
	return sum
}

func a2(k, n int) int {
	var sum int

	kChooseNumOnes := 1
	remChooseNumTwos := 1

	for numOnes := k; numOnes >= 0; numOnes -= 2 {
		numTwosAndZeros := k - numOnes

		// Number of ways they can be arranged
		// Choose positions of 1s, * choose position of 2s (then zeros are fixed)
		// (k choose numOnes) * ((k - numOnes) choose (numTwosAndZeros/2))

		// coef := maths.Choose(k, numOnes).Times(maths.Choose(k-numOnes, numTwosAndZeros/2))
		coef := (remChooseNumTwos * kChooseNumOnes) % mod

		v := (coef * maths.PowMod(2, numOnes*(n/k), mod) % mod)

		// numOnes can be in any order
		sum = (sum + v) % mod

		// Update kChooseNumOnes
		// kChooseNumOnes = (k choose numOnes)       = (k!) / ( (numOnes!) * (k - numOnes)! )
		//                                           = (k * (k-1) * ... (k - numOnes + 1) ) / (numOnes * (numOnes - 1) * ... * 1 )
		// Want             (k choose (numOnes - 2)) = (k!) / ( (numOnes-2)! * (k - numOnes + 2)! )
		//                                           = (k * (k-1) * ... (k - numOnes + 3) ) / ((numOnes-2) * (numOnes - 3) * ... * 1 )
		// Divide by (k - numOnes + 2) and (k - numOnes + 3)
		// Multiply by numOnes and (numOnes - 1)
		kChooseNumOnes = (kChooseNumOnes * maths.PowMod(k-numOnes+2, -1, mod)) % mod
		kChooseNumOnes = (kChooseNumOnes * maths.PowMod(k-numOnes+1, -1, mod)) % mod
		kChooseNumOnes = (kChooseNumOnes * numOnes) % mod
		kChooseNumOnes = (kChooseNumOnes * (numOnes - 1)) % mod

		// x choose y = x! / y! (x-y)!
		// (x+2) choose (y+1) =(x+2)(x+1)x! / (x+2 - (y+1))! (y+1)!
		//                    = (x+2)(x+1)x! / (x-y+1)! * (y+1) * y!
		//                    = (x+2)(x+1)x! / (x-y+1) * (x-y)! * (y+1) * y!
		//                    = ((x+2)(x+1) / (x-y+1)*(y+1) * (x! / (x-y)! * y!)
		//                    = ((x+2)(x+1) / (x-y+1)*(y+1) * (x choose y)
		x := k - numOnes
		y := numTwosAndZeros / 2
		remChooseNumTwos = (remChooseNumTwos * (x + 1)) % mod
		remChooseNumTwos = (remChooseNumTwos * (x + 2)) % mod
		remChooseNumTwos = (remChooseNumTwos * maths.PowMod(x-y+1, -1, mod)) % mod
		remChooseNumTwos = (remChooseNumTwos * maths.PowMod(y+1, -1, mod)) % mod
	}
	return sum
}
