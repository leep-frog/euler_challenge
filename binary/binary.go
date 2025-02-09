package binary

import (
	"fmt"
	"strings"

	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/functional"
)

// Binary represents a number in binary
type Binary struct {
	// digits is ordered from left to right in order of decreasing significance.
	// For example, [1, 0, 0, 1, 1] = 16 + 2 + 1 = 19
	digits []bool
}

func (b *Binary) IsZero() bool {
	return len(b.digits) == 0 || functional.All(b.digits, func(b bool) bool { return !b })
}

func (b *Binary) Copy() *Binary {
	var c []bool
	for _, d := range b.digits {
		c = append(c, d)
	}
	return &Binary{c}
}

// Double the value of the binary bit
func (b *Binary) Double() {
	b.digits = append(b.digits, false)
}

func (b *Binary) DoublePlusOne() {
	b.digits = append(b.digits, true)
}

func (b *Binary) Half() {
	if len(b.digits) > 0 {
		b.digits = b.digits[:len(b.digits)-1]
	}
}

func (b *Binary) Size() int {
	if b.IsZero() {
		return 1
	}
	return len(b.digits)
}

func (b *Binary) String() string {
	if b.IsZero() {
		return "0"
	}

	var r []string

	for _, v := range b.digits {
		if v {
			r = append(r, "1")
		} else {
			r = append(r, "0")
		}
	}

	return strings.Join(r, "")
}

func (b *Binary) ToInt() int {
	start := 1
	sum := 0
	for i := 0; i < len(b.digits); i++ {
		if b.At(i) {
			sum += start
		}
		start *= 2
	}
	return sum
}

func BinaryFromInt(k int) *Binary {
	var digits []bool
	for i := k; i > 0; i /= 2 {
		digits = append(digits, i%2 == 1)
	}
	return &Binary{bread.Reverse(digits)}
}

func (b *Binary) At(i int) bool {
	return b.digits[len(b.digits)-1-i]
}

func (b *Binary) ShortenedBinaryExpansion() string {
	var counts []string
	cur := true
	startIndex := 0
	for i, v := range b.digits {
		if v != cur {
			counts = append(counts, fmt.Sprintf("%d", i-startIndex))
			startIndex = i
			cur = !cur
		}
	}
	counts = append(counts, fmt.Sprintf("%d", len(b.digits)-startIndex))
	return strings.Join(counts, ",")
}

func (b *Binary) Minus(that *Binary) *Binary {
	var digits []bool
	var rem bool
	for i := 0; i < len(b.digits) && i < len(that.digits); i++ {
		count := 0
		if rem {
			count++
			rem = false
		}
		if that.At(i) {
			count++
		}
		if count == 0 {
			digits = append(digits, b.At(i))
		}
		if count == 1 {
			if b.At(i) {
				digits = append(digits, false)
			} else {
				digits = append(digits, true)
				rem = true
			}
		}
		if count == 2 {
			digits = append(digits, b.At(i))
			rem = true
		}
	}

	nb := &Binary{bread.Reverse(digits)}
	nb.trim()
	return nb
}

func (b *Binary) Reverse() *Binary {
	return &Binary{bread.Reverse(b.digits)}
}

func (b *Binary) XOR(that *Binary) *Binary {
	var digits []bool
	for i := 0; i < len(b.digits) || i < len(that.digits); i++ {

		left := i < len(b.digits) && b.digits[len(b.digits)-1-i]
		right := i < len(that.digits) && that.digits[len(that.digits)-1-i]
		digits = append(digits, left != right)
	}

	nb := &Binary{bread.Reverse(digits)}
	nb.trim()
	return nb
}

// Trim the leading zeroes
func (b *Binary) trim() {
	for len(b.digits) > 0 && !b.digits[0] {
		b.digits = b.digits[1:]
	}
}
