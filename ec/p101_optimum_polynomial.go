package eulerchallenge

import (
	"math/big"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P101() *problem {
	return intInputNode(101, func(o command.Output, k int) {
		/* Logic for solution
		  Let N = |--------------------|
			        | 1 1  1  1 ...      |
						  | 1 2  4  8 ...      |
						  | 3 8 27 81 ...      |
			  			| ...                |
			        |--------------------|

			u_i is i-th term in actual series
			Let U = |--------------------|
			        | u_1 u_1 u_1 ...    |
							| u_2 u_2 u_2 ...    |
							| u_3 u_3 u_3 ...    |
			  			| ...                |
			        |--------------------|

			Coefficients of actual solution
			Let A = |--------------------|
			        | a_1 a_2 a_3 ...    |
							| a_1 a_2 a_3 ...    |
							| a_1 a_2 a_3 ...    |
			  			| ...                |
			        |--------------------|

			(we need A and U to be square so the dimensions work out with the inverse matrix)
			A*N = U
			(A*N)*(N^-1) = U*(N^-1)
			A = U*(N^-1)
		*/

		// 1 − n + n^2 − n^3 + n^4 − n^5 + n^6 − n^7 + n^8 − n^9 + n^10
		uf := []int64{1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1}
		if k == 1 {
			// n^3
			uf = []int64{0, 0, 0, 1}
		}

		var us []*big.Rat
		for i := 1; i <= len(uf); i++ {
			sum := big.NewRat(0, 1)
			nTerm := big.NewRat(1, 1)
			for _, iCoef := range uf {
				coef := big.NewRat(iCoef, 1)
				sum.Add(sum, big.NewRat(0, 1).Mul(coef, nTerm))
				nTerm.Mul(nTerm, big.NewRat(int64(i), 1))
			}
			us = append(us, sum)
		}

		// Estimates
		var estimators [][]*big.Rat

		for size := 1; size <= len(us); size++ {
			var N, U [][]*big.Rat
			for i := 1; i <= size; i++ {
				v := big.NewRat(1, 1)
				var nRow []*big.Rat
				for j := 1; j <= size; j++ {
					nRow = append(nRow, v)
					v = big.NewRat(0, 1).Mul(v, big.NewRat(int64(i), 1))
				}
				N = append(N, nRow)
				U = append(U, us[:size])
			}
			N = maths.Transpose(N)
			I := maths.Inverse(N)
			A := maths.MultiplyMatrices(U, I)
			estimators = append(estimators, A[0])
		}

		fitSum := big.NewRat(0, 1)
		for i := 0; i < len(estimators)-1; i++ {
			e := estimators[i]
			width := big.NewRat(int64(len(e)+1), 1)
			pow := big.NewRat(1, 1)
			for _, coef := range e {
				fitSum.Add(fitSum, big.NewRat(0, 1).Mul(coef, pow))
				pow.Mul(pow, width)
			}
		}
		o.Stdoutln(fitSum.Num().Int64())
	})
}
