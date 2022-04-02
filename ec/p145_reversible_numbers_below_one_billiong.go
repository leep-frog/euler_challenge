package eulerchallenge

import (
	"github.com/leep-frog/command"
)

/*
Consider each spot as just a number from 0 to 18
(since putting 7 and 2 in paired spots is basically equivalent to
putting 4 and 5)

Then, we can just consider a sequence of
[k_1, k_2, ..., k_n/2, k_n/2, ..., k_2, k_1]
where k_i is between 0 and 18 and the sum of
k_1 + k_2*10 + k_3*100 + ... has all odd digits.

So, consider the boolean tuple for each k_i:
(odd, two_digits (ie >= 10)):

1. If k_i=PREV is (true, false), then digits are:
[..., PREV, NEXT, EMPTY, ..., EMPTY, NEXT, PREV]
(PREV is odd)
So k_i+1=NEXT must be odd since the right NEXT doesn't carry over a 1
and it can't be two digits, otherwise the left PREV would become even)

2. If If k_i=PREV is (true, true), then digits are:
[..., PREV, NEXT, EMPTY, ..., EMPTY, NEXT + 1, PREV]
(PREV is odd)
Now we need NEXT to be even so NEXT+1 is odd. Additionally,
we can't have NEXT be two digits, otherwise, PREV would become even.

3. If If k_i=PREV is (false, false), then digits are:
[..., PREV, NEXT, EMPTY, ..., EMPTY, NEXT, PREV+1]
(PREV is EVEN)
We need NEXT to be ODD, since the NEXT on the right doesn't have a carried
over 1. And we need NEXT to be two digits to make the PREV on the left odd.

4. If If k_i=PREV is (false, true), then digits are:
[..., PREV, NEXT, EMPTY, ..., EMPTY, NEXT+1, PREV+1]
(PREV is EVEN)
We need NEXT to be even so NEXT+1 is odd
We need NEXT to be two digits so PREV on the left becomes odd

In short, we have the following state machine:
1. (true, false) -> (true, false) (1)
2. (true, true) -> (false, false) (3)
3. (false, false) -> (true, true) (2)
4. (false, true) -> (false, true) (4)

Even length digits are in loop 1 -> 1 -> 1 -> ...
Odd length digits are in loop 2 -> 3 -> 2 -> 3 -> ...
*/

func P145() *problem {
	return intInputNode(145, func(o command.Output, n int) {
		var sum int
		for length := 2; length <= n; length++ {
			var count int
			if length%2 == 0 {
				// Loop is case 1 -> case 1 -> case 1 -> ...
				// Outer one is two digits and odd, but not using a zero
				// one digit and odd (no zeroes)
				// numbers:      1,  3,  5,  7,  9
				// ways to make: 0 + 2 + 4 + 6 + 8 = 20
				count = 20
				for j := length - 2; j > 0; j -= 2 {
					// one digit and odd (with zeroes)
					// numbers:      1,  3,  5,  7,  9
					// ways to make: 2 + 4 + 6 + 8 + 10 = 20
					count *= 30
				}
			} else {
				// Loop is case 2 -> case 3 -> case 2 -> ...
				// There will be one number exactly in the middle, and that
				// can't be odd. So if the last case is case 3, then this doesn't work
				if (length-1)%4 == 0 {
					continue
				}

				// middle point will need to be even and single digit, and single number
				// 5 ways: 1+1, 2+2, 3+3, 4+4, 0+0
				count = 5
				for j := length - 1; j > 0; j -= 2 {
					if j%4 == 2 {
						// two digits and odd:
						// numbers:     11, 13, 15, 17
						// ways to make: 8 + 6 + 4 + 2 = 20
						count *= 20
					} else {
						//
						// one digit and even:
						// numbers:      0,  2,  4,  6,  8
						// ways to make: 1 + 3 + 5 + 7 + 9 = 25
						count *= 25
					}
				}
			}
			sum += count
		}
		o.Stdoutln(sum)
	})
}
