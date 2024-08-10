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
	return ecmodels.IntInputNode(743, func(o command.Output, n int) {
		// o.Stdoutln(n)
		fmt.Println(a2(3, 9))
		fmt.Println(a2(4, 20))
		fmt.Println(a2(maths.Pow(10, 8), maths.Pow(10, 16)))
	}, []*ecmodels.Execution{
		{
			Args: []string{"1"},
			Want: "",
		},
		{
			Args: []string{"2"},
			Want: "259158998",
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
		// fmt.Println(k, numOnes, numTwosAndZeros)

		// big, small := numOnes, numTwosAndZeros
		// if big < small {
		// 	big, small = small, big
		//
		coef := maths.Choose(k, numOnes).Times(maths.Choose(k-numOnes, numTwosAndZeros/2))

		v := coef.Times(maths.BigPow(2, numOnes*(n/k)))
		fmt.Println(coef, v)

		// numOnes can be in any order
		sum = (sum + v.ModInt(mod)) % mod
	}
	return sum
}

func a2(k, n int) int {
	var sum int

	kChooseNumOnes := 1
	remChooseNumTwos := 1

	// (k-e) choose ()
	// (k-e)!

	for numOnes := k; numOnes >= 0; numOnes -= 2 {
		if numOnes%10_0000 == 0 {
			fmt.Println(numOnes)
		}
		// fmt.Println(k)
		numTwosAndZeros := k - numOnes

		// Number of ways they can be arranged
		// Choose positions of 1s, * choose position of 2s (then zeros are fixed)
		// (k choose numOnes) * ((k - numOnes) choose (numTwosAndZeros/2))

		// coef := maths.Choose(k, numOnes).Times(maths.Choose(k-numOnes, numTwosAndZeros/2))
		coef := (remChooseNumTwos * kChooseNumOnes) % mod
		// fmt.Println(maths.Choose(k, numOnes), kChooseNumOnes)
		// fmt.Println("NN", maths.Choose(k-numOnes, numTwosAndZeros/2), remChooseNumTwos)

		// v := coef.Times(maths.BigPow(2, numOnes*(n/k)))
		// v :=
		v := (coef * maths.PowMod(2, numOnes*(n/k), mod) % mod)
		// fmt.Println(coef, v)

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
