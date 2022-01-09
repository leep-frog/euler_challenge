package maths

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/google/go-cmp/cmp"
	"github.com/leep-frog/euler_challenge/parse"
)

func Divisors(i int) []int {
	var r []int
	for j := 1; j*j <= i; j++ {
		if i%j == 0 {
			if j*j == i {
				r = append(r, j)
			} else {
				r = append(r, j, i/j)
			}
		}
	}
	sort.Ints(r)
	return r
}

func IsSquare(i int) bool {
	rt := int(math.Sqrt(float64(i)))
	return rt*rt ==
		i
}

func Sqrt(i int) int {
	return int(math.Sqrt(float64(i)))
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Min(as ...int) int {
	if len(as) == 0 {
		return 0
	}
	min := as[0]
	for _, a := range as {
		if a < min {
			min = a
		}
	}
	return min
}

func Max(as ...int) int {
	if len(as) == 0 {
		return 0
	}
	max := as[0]
	for _, a := range as {
		if a > max {
			max = a
		}
	}
	return max
}

type Int struct {
	parts    []uint64
	negative bool
}

var (
	biggestInt = NewInt(math.MaxInt).Plus(One())
	zero       = Zero()
)

func (i *Int) IsZero() bool {
	return i.EQ(zero)
}

func (i *Int) ToInt() int {
	d, m := i.Divide(biggestInt)
	if d.NEQ(zero) {
		log.Fatalf("Int is too big to convert to int")
	}
	return parse.Atoi(m.String())
}

var (
	// maxInt ~2^30
	// var so it can be modified for testing purposes.
	maxIntCached uint64 = 0
	maxDigits           = 9
)

func maxInt() uint64 {
	if maxIntCached == 0 {
		maxIntCached = 1
		for i := 0; i < maxDigits; i++ {
			maxIntCached *= 10
		}
	}
	return maxIntCached
}

var (
	intReg = regexp.MustCompile("^(-?)([0-9]*)$")
)

func IntFromString(s string) (*Int, error) {
	m := intReg.FindStringSubmatch(s)
	if len(m) == 0 {
		return nil, fmt.Errorf("Invalid string: %s", s)
	}

	r := &Int{
		negative: m[1] == "-",
		parts:    make([]uint64, (len(s) / maxDigits)),
	}

	numString := m[2]
	if len(numString)%maxDigits != 0 {
		r.parts = append(r.parts, 0)
	}

	for idx := range r.parts {
		end := len(numString) - maxDigits*idx
		start := Max(end-maxDigits, 0)
		// Shouldn't be an error because of earlier regex check
		n, _ := strconv.Atoi(numString[start:end])
		r.parts[idx] = uint64(n)
	}
	r.trim()
	return r, nil
}

func MustIntFromString(s string) *Int {
	r, err := IntFromString(s)
	if err != nil {
		log.Fatal(err)
	}
	return r
}

func Pandigital(v int) bool {
	m := map[int]bool{}
	for _, d := range Digits(v) {
		if m[d] {
			return false
		}
		m[d] = true
	}

	for i := 1; i <= len(m); i++ {
		if !m[i] {
			return false
		}
	}

	return true
}

func NewInt(i int64) *Int {
	r := &Int{}
	if i == 0 {
		return r
	}

	if i < 0 {
		r.negative = true
		i *= -1
	}

	ui := uint64(i)

	for ui >= maxInt() {
		r.append(ui % maxInt())
		ui /= maxInt()
	}
	r.append(ui)
	return r
}

func (i *Int) Copy() *Int {
	r := &Int{
		negative: i.negative,
	}
	for _, p := range i.parts {
		r.parts = append(r.parts, p)
	}
	return r
}

func (i *Int) String() string {
	if i.size() == 0 {
		return "0"
	}
	var s []string
	if i.negative {
		s = append(s, "-")
	}
	miLen := maxDigits + 1
	for idx := i.size() - 1; idx >= 0; idx-- {
		a := fmt.Sprintf("%d", i.get(idx))
		// Add leading zeros if not the last one
		if idx != i.size()-1 {
			a = fmt.Sprintf("%s%s", strings.Repeat("0", miLen-len(a)-1), a)
		}
		s = append(s, a)
	}
	return strings.Join(s, "")
}

func (i *Int) trim() {
	end := i.size()
	for idx := i.size() - 1; idx >= 0; idx-- {
		if i.get(idx) == 0 {
			end--
		} else {
			break
		}
	}
	i.parts = i.parts[:end]
}

func (i *Int) size() int {
	return len(i.parts)
}

func (i *Int) append(v uint64) {
	i.parts = append(i.parts, v)
}

func (i *Int) get(spot int) uint64 {
	return i.parts[spot]
}

func SumInts(is ...int) int {
	var s int
	for _, i := range is {
		s += i
	}
	return s
}

func Rotations(parts []string) []string {
	var r []string
	for i := 0; i < len(parts); i++ {
		r = append(r, strings.Join(append(parts[i:], parts[:i]...), ""))
	}
	return r
}

func Sets(parts []int) [][]int {
	m := map[string]bool{}
	r := [][]int{}
	sets(parts, m, []int{}, &r)
	return r
}

func sets(remaining []int, m map[string]bool, cur []int, r *[][]int) {
	if len(remaining) == 0 {
		if len(cur) == 0 {
			return
		}

		if s := fmt.Sprintf("%v", cur); m[s] {
			return
		} else {
			m[s] = true
		}

		k := make([]int, len(cur))
		copy(k, cur)
		*r = append(*r, k)
		return
	}

	sets(remaining[1:], m, cur, r)
	cur = append(cur, remaining[0])
	sets(remaining[1:], m, cur, r)
	cur = cur[:len(cur)-1]
}

func Permutations(parts []string) []string {
	m := map[string]bool{}
	remaining := map[string]int{}
	for _, part := range parts {
		remaining[part]++
	}
	permutations(parts, remaining, m, []string{})
	var r []string
	for perm := range m {
		r = append(r, perm)
	}
	return r
}

func (b *Binary) Len() int {
	return len(b.digits)
}

func (b *Binary) Concat(that *Binary) *Binary {
	var d []bool
	for _, v := range b.digits {
		d = append(d, v)
	}
	for _, v := range that.digits {
		d = append(d, v)
	}
	return &Binary{d}
}

func permutations(m []string, remaining map[string]int, r map[string]bool, cur []string) {
	if len(cur) == len(m) {
		r[strings.Join(cur, "")] = true
	}

	for _, p := range m {
		if remaining[p] == 0 {
			continue
		}
		cur = append(cur, p)
		remaining[p]--
		permutations(m, remaining, r, cur)
		remaining[p]++
		cur = (cur)[:len(cur)-1]
	}
}

func SumI(is ...int64) *Int {
	var ints []*Int
	for _, i := range is {
		ints = append(ints, NewInt(i))
	}
	return Sum(ints...)
}

func (i *Int) Palindrome() bool {
	s := i.String()
	for idx := range s {
		if s[idx:idx+1] != s[len(s)-idx-1:len(s)-idx] {
			return false
		}
	}
	return true
}

func (i *Int) Reverse() *Int {
	var r []string
	magOnlyFunc(i, func(pos *Int) {
		s := pos.String()
		for i := range s {
			r = append(r, s[len(s)-1-i:len(s)-i])
		}
	})
	return MustIntFromString(strings.Join(r, ""))
}

func Sum(is ...*Int) *Int {
	if len(is) == 0 {
		return &Int{}
	}

	r := is[0].Copy()
	for idx := 1; idx < len(is); idx++ {
		r = r.Plus(is[idx])
	}
	return r
}

func (i *Int) Plus(that *Int) *Int {
	if i.negative == that.negative {
		r := &Int{
			negative: i.negative,
		}
		var remainder uint64
		for idx := 0; idx < Max(i.size(), that.size()); idx++ {
			sum := remainder
			if idx < i.size() {
				sum += i.get(idx)
			}
			if idx < that.size() {
				sum += that.get(idx)
			}
			r.append(sum % maxInt())
			remainder = sum / maxInt()
		}
		if remainder != 0 {
			r.append(remainder)
		}
		r.trim()
		return r
	}

	// Otherwise we are subtracting

	// guarantee magnitude of "i" is always GTE "that".
	if i.LT(that.Times(NewInt(-1))) == !i.negative {
		return that.Plus(i)
	}

	r := &Int{
		negative: i.negative,
	}
	var borrowed bool
	for idx := 0; idx < i.size(); idx++ {
		// Remove one digit if the previous subtraction needed to split
		curRes := i.get(idx)
		if borrowed {
			// If zero, then we need to borrow again.
			if curRes == 0 {
				curRes = maxInt() - 1
			} else {
				curRes--
				borrowed = false
			}
		}
		if idx < that.size() {
			t := that.get(idx)
			if t > curRes {
				curRes += maxInt()
				borrowed = true
			}
			curRes -= t
		}

		r.append(curRes)
	}
	r.trim()
	return r
}

// Rule: Int.parts[i] is always the largest it can be
// TODO: trim

func (i *Int) LT(that *Int) bool {
	if i.negative != that.negative {
		return i.negative
	}

	if i.size() != that.size() {
		return i.size() < that.size() != i.negative
	}

	for idx := i.size() - 1; idx >= 0; idx-- {
		if i.get(idx) != that.get(idx) {
			return (i.get(idx) < that.get(idx)) != i.negative
		}
	}
	return false
}

func (i *Int) EQ(that *Int) bool {
	return !i.LT(that) && !that.LT(i)
}

func (i *Int) NEQ(that *Int) bool {
	return !i.EQ(that)
}

func (i *Int) GT(that *Int) bool {
	return that.LT(i)
}

func (i *Int) GTE(that *Int) bool {
	return !i.LT(that)
}

func (i *Int) LTE(that *Int) bool {
	return !i.GT(that)
}

// Magnitude less than.
func (i *Int) MagLT(that *Int) bool {
	var b bool
	magsOnlyFunc(i, that, func(i1, i2 *Int) {
		b = i1.LT(i2)
	})
	return b
}

func (i *Int) MagEQ(that *Int) bool {
	return !i.MagLT(that) && !that.MagLT(i)
}

func (i *Int) MagNEQ(that *Int) bool {
	return !i.MagEQ(that)
}

func (i *Int) MagGT(that *Int) bool {
	return that.MagLT(i)
}

func (i *Int) MagGTE(that *Int) bool {
	return !i.MagLT(that)
}

func (i *Int) MagLTE(that *Int) bool {
	return !i.MagGT(that)
}

func (i *Int) PP() {
	*i = *(i.Plus(NewInt(1)))
}

func (i *Int) MM() {
	*i = *(i.Plus(NewInt(-1)))
}

func (i *Int) Times(that *Int) *Int {
	var rs []*Int
	for idx := 0; idx < i.size(); idx++ {
		r := &Int{}
		for offset := 0; offset < idx; offset++ {
			r.append(0)
		}
		var remainder uint64
		for jdx := 0; jdx < that.size(); jdx++ {
			product := i.get(idx)*that.get(jdx) + remainder
			r.append(product % maxInt())
			remainder = product / maxInt()
		}
		if remainder != 0 {
			r.append(remainder)
		}
		rs = append(rs, r)
	}

	v := Sum(rs...)
	v.negative = i.negative != that.negative
	v.trim()
	return v
}

func (i *Int) MagMinus(that *Int) *Int {
	var r *Int
	magsOnlyFunc(i, that, func(i1, i2 *Int) {
		r = i1.Minus(i2)
	})
	return r
}

func magsOnlyFunc(this, that *Int, f func(*Int, *Int)) {
	thisNeg := this.negative
	thatNeg := that.negative
	this.negative = false
	that.negative = false
	f(this, that)
	this.negative = thisNeg
	that.negative = thatNeg
}

func magOnlyFunc(this *Int, f func(*Int)) {
	thisNeg := this.negative
	this.negative = false
	f(this)
	this.negative = thisNeg
}

func (i *Int) Minus(that *Int) *Int {
	that.Negate()
	r := i.Plus(that)
	that.Negate()
	return r
}

func (i *Int) Negate() {
	i.negative = !i.negative
}

func One() *Int {
	return NewInt(1)
}

func Zero() *Int {
	return NewInt(0)
}

func (i *Int) DivInt(by uint16) *Int {
	a, _ := i.divInt(by)
	return a
}

func (i *Int) ModInt(by uint16) uint16 {
	_, b := i.divInt(by)
	return b
}

func (i *Int) divInt(by16 uint16) (*Int, uint16) {
	if by16 == 0 {
		log.Fatal("Divide by zero exception")
	}
	by := uint64(by16)
	var rem uint16
	ret := &Int{
		negative: i.negative,
		parts:    make([]uint64, i.size()),
	}
	for idx := i.size() - 1; idx >= 0; idx-- {
		num := i.get(idx) + uint64(rem)*maxInt()
		ret.parts[idx] = num / by
		rem = uint16(num % by)
	}
	ret.trim()
	return ret, rem
}

func BigPow(a, b int) *Int {
	if b == 0 {
		return One()
	}
	ai := NewInt(int64(a))
	r := NewInt(int64(a))
	for i := 1; i < b; i++ {
		r = r.Times(ai)
	}
	return r
}

func Range(n int) []int {
	return make([]int, n)
}

func Pow(a, b int) int {
	if b == 0 {
		return 1
	}
	ogA := a
	for i := 1; i < b; i++ {
		ogA *= a
	}
	return ogA
}

type Binary struct {
	digits []bool
}

var (
	binaryRegex = regexp.MustCompile("^[01]*$")
)

func NewBinary(bs string) *Binary {
	if !binaryRegex.MatchString(bs) {
		log.Fatal("invalid binary string")
	}
	b := &Binary{}
	for i := len(bs) - 1; i >= 0; i-- {
		b.digits = append(b.digits, bs[i:i+1] == "1")
	}
	return b
}

// Return palindrome numbers that are n digits long.
func Palindromes(n int) []int {
	if n == 0 {
		return nil
	}
	var r, cur []string
	palindromeLeft(n, &cur, &r)

	var p []int
	for _, v := range r {
		p = append(p, parse.Atoi(v))
	}
	return p
}

func palindromeLeft(n int, cur, r *[]string) {
	start := 0
	if len(*cur) == 0 {
		start = 1
	}

	if n == 0 {
		full := *cur
		for i := len(full) - 1; i >= 0; i-- {
			full = append(full, full[i])
		}
		*r = append(*r, strings.Join(full, ""))
		return
	}
	if n == 1 {
		for s := start; s <= 9; s++ {
			full := append(*cur, strconv.Itoa(s))
			for i := len(full) - 2; i >= 0; i-- {
				full = append(full, full[i])
			}
			*r = append(*r, strings.Join(full, ""))
		}
		return
	}

	for i := start; i <= 9; i++ {
		*cur = append(*cur, strconv.Itoa(i))
		palindromeLeft(n-2, cur, r)
		*cur = (*cur)[:len(*cur)-1]
	}
}

type Bester struct {
	better func(int, int) bool
	best   int
	bestI  int

	set bool
}

func (b *Bester) Best() int {
	return b.best
}

func (b *Bester) BestIndex() int {
	return b.bestI
}

func (b *Bester) Check(v int) {
	if !b.set || b.better(v, b.best) {
		b.best = v
		b.set = true
	}
}

func (b *Bester) IndexCheck(idx, v int) {
	if !b.set || b.better(v, b.best) {
		b.best = v
		b.bestI = idx
		b.set = true
	}
}

func Largest() *Bester {
	return &Bester{
		better: func(i, j int) bool {
			return i > j
		},
	}
}

func Smallest() *Bester {
	return &Bester{
		better: func(i, j int) bool {
			return i < j
		},
	}
}

type IncrementalBester struct {
	b *Bester
	m map[int]int
}

func (ib *IncrementalBester) Best() int {
	return ib.b.best
}

func (ib *IncrementalBester) BestIndex() int {
	return ib.b.bestI
}

func (ib *IncrementalBester) Increment(v int) {
	ib.m[v]++
	ib.b.IndexCheck(v, ib.m[v])
}

func LargestIncremental() *IncrementalBester {
	return &IncrementalBester{
		b: Largest(),
		m: map[int]int{},
	}
}

func Digits(n int) []int {
	var r []int
	for v, i := strconv.Itoa(n), 0; i < len(v); i++ {
		r = append(r, parse.Atoi(v[i:i+1]))
	}
	return r
}

func (b *Binary) Palindrome() bool {
	for i := 0; i <= len(b.digits)/2; i++ {
		j := len(b.digits) - 1 - i
		if b.digits[i] != b.digits[j] {
			return false
		}
	}
	return true
}

func (b *Binary) Equals(that *Binary) bool {
	if len(b.digits) != len(that.digits) {
		return false
	}

	for i, v := range b.digits {
		if v != that.digits[i] {
			return false
		}
	}
	return true
}

type Intable interface {
	ToInt() int
}

func SumType[T Intable](ts []T) int{
	var sum int
	for _, t := range ts {
		sum += t.ToInt()
	}
	return sum
}

// TODO: map package
func Set[K1, K2 comparable, V any](m map[K1]map[K2]V, k1 K1, k2 K2, v V) {
	if m[k1] == nil {
		m[k1] = map[K2]V{}
	}
	m[k1][k2] = v
}

func XOR(a, b int) int {
	return ToBinary(a).XOR(ToBinary(b)).ToInt()
}

func (b *Binary) ToInt() int {
	start := 1
	total := 0
	for i := 0; i < len(b.digits); i++ {
		if b.digits[i] {
			total += start
		}
		start *= 2
	}
	return total
}

func (b *Binary) XOR(that *Binary) *Binary {
	end := Max(len(b.digits), len(that.digits))
	var d []bool
	for i := 0; i < end; i++ {
		f := i < len(b.digits) && b.digits[i]
		s := i < len(that.digits) && that.digits[i]
		d = append(d, f != s)
	}
	return &Binary{d}
}

func ToBinary(i int) *Binary {
	b := &Binary{}
	for ; i > 0; i /= 2 {
		b.digits = append(b.digits, i%2 == 1)
	}
	return b
}

func (b *Binary) String() string {
	var s []string
	for i := len(b.digits) - 1; i >= 0; i-- {
		if b.digits[i] {
			s = append(s, "1")
		} else {
			s = append(s, "0")
		}
	}
	return strings.Join(s, "")
}

func (i *Int) Div(that *Int) *Int {
	q, _ := i.Divide(that)
	return q
}

func (i *Int) Mod(that *Int) *Int {
	_, m := i.Divide(that)
	return m
}

func (i *Int) Divide(that *Int) (*Int, *Int) {
	var q, r *Int
	magsOnlyFunc(i, that, func(i, that *Int) {
		if that.EQ(zero) {
			log.Fatal("Divide by zero exception")
		}

		// Make "start" the biggest power of 2 such that (that * start) <= i
		start := One()
		for two := NewInt(2); start.Times(that).LTE(i); start = start.Times(two) {
		}
		start = start.DivInt(2)

		// Start subtracting
		ret := i.Copy()
		ret.negative = false
		quotient := NewInt(0)
		for ret.GTE(that) {
			if prod := start.Times(that); prod.LTE(ret) {
				quotient = quotient.Plus(start)
				ret = ret.Minus(prod)
			}
			start = start.DivInt(2)
			if start.EQ(zero) {
				quotient.trim()
				q, r = quotient, ret
			}
		}

		quotient.trim()
		q, r = quotient, ret
	})
	q.negative = i.negative != that.negative
	return q, r
}

func CmpOpts() []cmp.Option {
	return []cmp.Option{
		cmp.Comparer(func(this, that *Int) bool {
			if this == nil {
				return that == nil || that.EQ(zero)
			}
			if that == nil {
				return this == nil || this.EQ(zero)
			}
			return that != nil && this.EQ(that)
		}),
		cmp.Comparer(func(this, that *Binary) bool {
			return this.Equals(that)
		}),
	}
}

func (i *Int) Digits() []int {
	var r []int
	magOnlyFunc(i, func(i1 *Int) {
		for v, idx := i1.String(), 0; idx < len(v); idx++ {
			r = append(r, parse.Atoi(v[idx:idx+1]))
		}
	})
	return r
}

func (i *Int) DigitSum() int {
	var sum int
	for _, d := range i.Digits() {
		sum += d
	}
	return sum
}

// TODO: change Div function
func Choose(n, r int) *Int {
	return Factorial(n).Div(Factorial(r).Times(Factorial(n - r)))
}

func Factorial(n int) *Int {
	r := One()
	for i := 1; i <= n; i++ {
		r = r.Times(NewInt(int64(i)))
	}
	return r
}

func FactorialI(n int) int {
	r := 1
	for i := 2; i <= n; i++ {
		r *= i
	}
	return r
}
