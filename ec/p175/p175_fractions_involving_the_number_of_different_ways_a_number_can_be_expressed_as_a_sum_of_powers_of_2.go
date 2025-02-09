package p175

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/binary"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/fraction"
	"github.com/leep-frog/euler_challenge/maths"
)

// Related to problem 169

// Number order:
//  1  2
//  1  4  3  2 (swap 2-4)
//  1  8  7  2  5  4  3  6

// 1/5     4/5     3/7       4/7       2/7     5/7       3/8       5/8
//   1       8       7         2         5       4         3         6
// 1/6 5/6 4/9 5/9 3/10 7/10 4/11 7/11 2/9 7/9 5/12 7/12 3/11 8/11 5/13 8/13
//   1  16  15   2   13    4    3   14   9   8    7   10    5   12   11    6

//  1  2
//  1  4  3  2
//  1  4  7  6  5  8  3  2
//  1 16 15  2  13  4    3   14   9   8    7   10    5   12   11    6

//                                1
//                1                               2
//        1               4               3               2
//    1       8       7       2       5       4       3       6
//  1  16  15   2  13   4   3  14   9   8   7  10   5  12  11   6

//                                                 1
//                       1                                                   10
//            1                      100                         11                     10
//    1             1000        111         10            101          100         11         110
//  1   10000    1111   10  1101   100   1011  1110   1001   1000   111  1010   101  1100  1011   110
// If at node with value N:
// Go left -> 2*N - 1
// Go right -> maths.Pow(2, k) - (2*n - 1)

//                                1
//                2                               3
//        4               7               6               5
//    8       15       14       9       12       11       10       13
//  16  31  30   17  28   19   18  29   24   23   22  25   20  27  26   21

//                                1
//                2                                   3
//        4               7                  6               5
//    8       15       14       9         12       11       10       13
//  16  31  30   17  28   19   18  29   24   23   22  25   20  27  26   21

//                                                            1
//                             10                                                           11
//             100                            111                           110                           101
//      1000          1111            1110           1001           1100            1011           1010          1101
//  10000  11111  11110   10001  11100   10011   10010  11101   11000   10111   10110  11001   10100  11011  11010   10101

// Shortened Binary Expansion (1 index)
//                                      1
//                  1,1                                         2
//          1,2               3                    2,1                      1,1,1
//    1,3       4       3,1       1,2,1       2,2       1,1,2       1,1,1,1       2,1,1

// Shortened Binary Expansion (0 index)
//                                0
//                1                               2
//        3               6               5               4
//    7       14       13       8       11       10        9       12
//  15  30  29   16  27   18   17  28   23   22   21  24   19  26  25   20

// Shortened Binary Expansion (0 index)
//                                                           0
//                             1                                                           10
//             11                           110                            101                            100
//     111            1110           1101           1000            1011           1010           1001          1100
//  1111  11110  11101   10000  11011   10010   10001  11100   10111   10110   10101  11000   10011  11010  11001   10100

// Shortened Binary Expansion (0 index)
//                                _
//                1                               1,1
//        2               2,1               1,1,1               1,2
//    3         3,1         2,1,1       1,3       1,1,2       1,1,1,1        1,2,1       2,2
//  4  4,1  3,1,1  1,4  2,1,2   1,2,1,1   1,3,1  3,2   1,1,3   1,1,2,1   1,1,1,1,1  2,3   1,2,2  2,1,1,1  2,2,1   1,1,1,2

// Pattern:
// Do the first quarter the same as prev first half,
// the second quarter the same as prev second half + k/2,
// the third quarter the same as prev first half + k/2,
// the second quarter the same as prev second half

func P175() *ecmodels.Problem {
	return ecmodels.IntsInputNode(175, 2, 0, func(o command.Output, ns []int) {
		f := fraction.New(ns[0], ns[1])

		// Use Kepler tree of fractions (https://oeis.org/A294442)
		// with the following observation:
		// A node on the Kepler tree is a
		// - left child if it is less than one-half
		// - right child otherwise
		// However, the ordering of fractions is not as straight-forward:
		// Here is the Kepler Tree
		//                             1/1
		//                             1/2
		//             1/3                             2/3
		//     1/4             3/4             2/5             3/5
		// 1/5     4/5     3/7     4/7     2/7     5/7     3/8     5/8

		// And here is the ordering of fractions we get from the actual sequence.
		//                             1/1
		//                             1/2
		//             2/3                             1/3
		//     3/4             2/5             3/5             1/4
		// 4/5     3/7     5/8     2/7     5/7     3/8     4/7     1/5
		// 5/6 ...

		// If we label the Kepler tree fractions in their order relative to this sequence, we get:
		//                             1/1 (0)
		//                             1/2 (1)
		//             1/3 (3)                         2/3 (2)
		//     1/4 (7)         3/4 (4)         2/5 (5)         3/5 (6)
		// 1/5(15) 4/5 (8) 3/7 (9) 4/7(14) 2/7(11) 5/7(12) 3/8(13) 5/8 (2)
		// Removing the fractions, we see
		//                      0
		//                      1
		//          3                       2
		//    7           4           5           6
		// 15    8     9    14    11    12    13     2

		// Notice the following traits:
		// * f(n->left) = 2*f(n) + 1
		// * Each pair of siblings in the same row has the same sum
		//   with the sequence (5, 11, 23, ...) -> double and add one at each step
		//   Let s_k be the sequence for a given row, then
		//   s_(k+1) = 2*s_k + 1
		// * f(n->right) = s_k - (2*f(n) + 1) = s_k - f(n->left)

		// Create the Kepler path, where (path[i] == true) means we go left at row i
		var path []bool
		half := fraction.New(1, 2)
		for maths.NEQ(f, half) {
			path = append(path, f.LT(half))
			if f.LT(half) {
				f = fraction.New(f.N, f.D-f.N)
			} else {
				f = fraction.New(f.D-f.N, f.N)
			}
		}
		path = bread.Reverse(path)

		// Keep track of the sequence and the index value
		fractionIndex := binary.BinaryFromInt(1)
		sequenceValue := binary.BinaryFromInt(5)
		for _, left := range path {
			// f(n->left) = 2*f(n) + 1
			fractionIndex.DoublePlusOne()
			if !left {
				// * f(n->right) = s_k - (2*f(n) + 1) = s_k - f(n->left)
				fractionIndex = sequenceValue.Minus(fractionIndex)
			}

			// s_(k+1) = 2*s_k + 1
			sequenceValue.DoublePlusOne()
		}

		// The number itself is actually the nth number, where the numerator and
		// denominator are each a separate value, so double and increment (to undo 0 index).
		fractionIndex.DoublePlusOne()
		o.Stdoutln(fractionIndex.ShortenedBinaryExpansion())
	}, []*ecmodels.Execution{
		{
			Args: []string{"4", "7"},
			Want: "3,1,1",
		},
		{
			Args: []string{"13", "17"},
			Want: "4,3,1",
		},
		{
			Args: []string{"123456789", "987654321"},
			Want: "1,13717420,8",
		},
	})
}
