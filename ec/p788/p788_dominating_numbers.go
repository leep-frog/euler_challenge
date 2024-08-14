package p788

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

const (
	mod = 1_000_000_007
)

func P788() *ecmodels.Problem {
	return ecmodels.IntInputNode(788, func(o command.Output, n int) {
		var sum int
		for i := 1; i <= n; i++ {
			sum = (sum + dominatingNumbers(i)) % mod
		}
		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args: []string{"4"},
			Want: "603",
		},
		{
			Args: []string{"10"},
			Want: "21893256",
		},
		{
			Args:     []string{"2022"},
			Want:     "471745499",
			Estimate: 5,
		},
	})
}

// dominatingNumbers returns the number of dominating numbers with *exacly* n digits.
func dominatingNumbers(n int) int {
	var sum int

	nChooseI := 1 // n choose 0 is 1
	for i := 1; i < (n/2)+1; i++ {
		//
		nChooseI = (nChooseI * (n - i + 1)) % mod
		inv := maths.PowMod(i, -1, mod)
		nChooseI = (nChooseI * inv) % mod
	}

	for i := (n / 2) + 1; i < n; i++ {
		// number of non-dominating digits (s for subordinate digits)
		s := n - i

		// Update nChooseI
		nChooseI = (nChooseI * (n - i + 1)) % mod
		inv := maths.PowMod(i, -1, mod)
		nChooseI = (nChooseI * inv) % mod

		// Dominating digit is non-zero and in the front
		// 9 -> non-zero digits
		// (n-1 choose i-1) -> places aside the front where the dominating digit is
		// 9^s -> options for non-dominating digits
		//
		// v1 = 9 * (n-1 choose i-1) * 9^s
		//    = (n-1 choose i-1) * 9^(s+1)

		// Dominating digit is non-zero and not in the front
		// 9 -> non-zero digit in the front
		// (n-1 choose i) -> places aside the front where the dominating digit is
		// 9^(s-1) -> options for non-dominating digits
		//
		// v2 = 9 * (n-1 choose i) * 8 * 9^(s-1)
		//    = (n-1 choose i) * 8 * 9^s

		// Dominating digit is zero and is not in the front
		// (n-1 choose i) -> places aside the front where the zeros are
		// 9^s -> options for non-dominating digits
		// v3 = (n-1 choose i) * 9^s

		// v2+v3 = (n-1 choose i) * 8 * 9^s + (n-1 choose i) * 9^s
		//       = (n-1 choose i) * (8*9^s + 9^s)
		//       = (n-1 choose i) * 9^s (8 + 1)
		//       = (n-1 choose i) * 9^(s+1)

		// v1+v23 = (n-1 choose i-1) * 9^(s+1) + (n-1 choose i) * 9^(s+1)
		//        = 9^(s+1) * ((n-1 choose i-1) + (n-1 choose i))

		//                                      a!              a!
		// (a choose b) + (a choose b-1) = -----------  +  -----------------
		//                                 (a-b)! * b!     (a-b+1)! * (b-1)!
		//
		//                                      a!                       a!          * b
		//                               = -----------  +  -----------------------------
		//                                 (a-b)! * b!     (a-b+1) * (a-b)! * (b-1)! * b
		//
		//                                      a!            b           a!
		//                               = -----------  +  ------- * -----------
		//                                 (a-b)! * b!     (a-b+1) * (a-b)! * b!
		//
		//                                      a!       (       b   )
		//                               = ----------- * ( 1 + ----- )
		//                                 (a-b)! * b!   (     a-b+1 )
		//
		//                                      a!       ( a-b+1+b )
		//                               = ----------- * ( ------- )
		//                                 (a-b)! * b!   (  a-b+1  )
		//
		//                                      a!       (  a+1  )
		//                               = ----------- * ( ----- )
		//                                 (a-b)! * b!   ( a-b+1 )
		//
		//                                                (  a+1  )
		//                               = (a choose b) * ( ----- )
		//                                                ( a-b+1 )
		//
		// Let: a = n-1, b = i
		// (n-1 choose i-1) + (n-1 choose i) = (n-1 choose i) * ((n-2)/(n-i))
		// v123 = 9^(s+1) * (n-1 choose i) * n / (n-i)
		//      = 9^(s+1) * { (n-1)! / (i! * (n-1-i)!) } * { n / (n-i) }
		//      = 9^(s+1) * { (n-1)! * n } / { (i! * (n-1-i)! * (n-i) }
		//      = 9^(s+1) * { n! } / { (i! * (n-i)! }
		//      = 9^(s+1) * (n choose i)
		// v123 := maths.BigPow(9, s+1).Times(maths.Choose(n, i))
		v123 := (maths.PowMod(9, s+1, mod) * nChooseI) % mod

		sum = (sum + v123) % mod

	}
	return (sum + 9%mod)
}

func brute(n int) int {
	var cnt int
	for i := 1; i < maths.Pow(10, n); i++ {
		ds := maths.DigitMap(i)
		for _, v := range ds {
			if v > len(maths.Digits(i))/2 {
				cnt++
				break
			}
		}
	}
	return cnt
}
