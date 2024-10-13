package maths

// NthDigit returns the nth decimal digit of the number (1 / k)
//
// Below are terms to make this easy to ctrl+shift+f for:
// nth digit nth-digit
func NthDigit(k, n int) int {
	// d_n(k) = [ floor(10^n) / k ] mod 10

	// Not sure why (got equality online), but
	// d_n(k) = floor[ (10^n mod 10k) / k ]
	// return maths.PowMod(10, n, 10*k) / k

	// The above can also be simplified to
	// d_n(k) = floor[ 10 * (10^(n-1) mod k) / k ]
	return 10 * PowMod(10, n-1, k) / k
}
