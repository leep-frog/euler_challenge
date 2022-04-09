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
	"golang.org/x/exp/constraints"
)

type Operable[T any] interface {
	Plus(T) T
	Times(T) T
	//Div(T) (T, T)
	Comparable[T]
}

type Comparable[T any] interface {
	LT(T) bool
}

type Mathable interface {
	~int | ~float64 | ~int64 | ~uint64
}

type mathableOperator[T Mathable] struct {
	m T
}

func newMo[T Mathable](t T) *mathableOperator[T] {
	return &mathableOperator[T]{t}
}

func (mo *mathableOperator[T]) Plus(that *mathableOperator[T]) *mathableOperator[T] {
	return newMo[T](mo.m + that.m)
}

func (mo *mathableOperator[T]) Times(that *mathableOperator[T]) *mathableOperator[T] {
	return newMo[T](mo.m * that.m)
}

/*func (mo *mathableOperator[T]) Div(that *mathableOperator[T]) (*mathableOperator[T], *mathableOperator[T]) {
	return newMo[T](mo.m / that.m), newMo[T](mo.m % that.m)
}*/

func (mo *mathableOperator[T]) LT(that *mathableOperator[T]) bool {
	return mo.m <= that.m
}

var (
	cachedDivisors = map[int][]int{}
)

func Divisors(i int) []int {
	v, ok := cachedDivisors[i]
	if !ok {
		for j := 1; j*j <= i; j++ {
			if i%j == 0 {
				if j*j == i {
					v = append(v, j)
				} else {
					v = append(v, j, i/j)
				}
			}
		}
		sort.Ints(v)
		cachedDivisors[i] = v
	}
	r := make([]int, len(v))
	copy(r, v)
	return r
}

func IsSquare(i int) bool {
	rt := int(math.Sqrt(float64(i)))
	return rt*rt == i
}

func IsUSquare(i uint64) bool {
	rt := uint64(math.Sqrt(float64(i)))
	return rt*rt == i
}

func Sqrt(i int) int {
	return int(math.Sqrt(float64(i)))
}

func Abs[T Mathable](a T) T {
	if a < 0 {
		return -a
	}
	return a
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

func (i *Int) Negative() bool {
	return i.negative
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

func (i *Int) TrimDigits(n int) *Int {
	iStr := i.String()
	n = Min(n, len(iStr))
	return MustIntFromString(iStr[len(iStr)-n:])
}

func Chop(n, from, to int) int {
	s := fmt.Sprintf("%d", n)
	from = Max(0, from)
	to = Min(len(s), to)
	return parse.Atoi(s[from:to])
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

func SumSys[T Mathable](ts ...T) T {
	var s T
	for _, t := range ts {
		s += t
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

func ChooseSets[T constraints.Ordered](parts []T, n int) [][]T {
	cur := []T{}
	var r [][]T
	chooseSets(parts, map[T]bool{}, n, &cur, &r)
	return r
}

func chooseSets[T constraints.Ordered](parts []T, ignore map[T]bool, n int, cur *[]T, all *[][]T) {
	if n == 0 && len(*cur) > 0 {
		new := make([]T, len(*cur))
		copy(new, *cur)
		*all = append(*all, new)
		return
	}

	if len(parts) == 0 {
		return
	}

	for idx, p := range parts {
		/*if ignore[p] {
			continue
		}
		ignore[p] = true*/
		*cur = append(*cur, p)
		chooseSets(parts[idx+1:], ignore, n-1, cur, all)
		*cur = (*cur)[:len(*cur)-1]
		//chooseSets(parts[1:], ignore, n, cur, all)
		//delete(ignore, p)
	}
}

type Trie[T comparable] struct {
	subTries      map[T]*Trie[T]
	endOfSequence bool
}

func (t *Trie[T]) Insert(ts []T) {
	if len(ts) == 0 {
		t.endOfSequence = true
		return
	}

	sub, ok := t.subTries[ts[0]]
	if !ok {
		t.subTries[ts[0]] = NewTrie[T]()
		sub = t.subTries[ts[0]]
	}
	sub.Insert(ts[1:])
}

func (t *Trie[T]) values(cur *[]T, cum *[][]T) {
	if t.endOfSequence {
		k := make([]T, len(*cur))
		copy(k, *cur)
		*cum = append(*cum, k)
	}

	for v, sub := range t.subTries {
		*cur = append(*cur, v)
		sub.values(cur, cum)
		*cur = (*cur)[:len(*cur)-1]
	}
}

func NewTrie[T comparable]() *Trie[T] {
	return &Trie[T]{map[T]*Trie[T]{}, false}
}

func StringPermutations(parts []string) []string {
	var r []string
	for _, perm := range Permutations(parts, len(parts), false) {
		r = append(r, strings.Join(perm, ""))
	}
	return r
}

func Permutations[T comparable](parts []T, length int, allowReplacements bool) [][]T {
	root := NewTrie[T]()

	remaining := map[T]int{}
	for _, part := range parts {
		remaining[part]++
	}
	permutations(parts, remaining, []T{}, root, length, allowReplacements)

	var cur []T
	var r [][]T
	root.values(&cur, &r)
	return r
}

func permutations[T comparable](m []T, remaining map[T]int, cur []T, root *Trie[T], length int, allowReplacements bool) {
	if len(cur) == length {
		root.Insert(cur)
		return
	}

	for _, p := range m {
		if remaining[p] == 0 {
			continue
		}
		cur = append(cur, p)
		if !allowReplacements {
			remaining[p]--
		}
		permutations(m, remaining, cur, root, length, allowReplacements)
		if !allowReplacements {
			remaining[p]++
		}
		cur = (cur)[:len(cur)-1]
	}
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

func Join[T any](ts []T, s string) string {
	var r []string
	for _, t := range ts {
		r = append(r, fmt.Sprintf("%v", t))
	}
	return strings.Join(r, s)
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

// TODO: separate package for this
type Mappable interface {
	Code() string
}

type Map[K Mappable, V any] struct {
	m  map[string]V
	km map[string]K
}

func NewMap[K Mappable, V any]() *Map[K, V] {
	return &Map[K, V]{
		m:  map[string]V{},
		km: map[string]K{},
	}
}

func (m *Map[K, V]) ForKs(f func(K) bool) {
	for _, k := range m.km {
		if f(k) {
			break
		}
	}
}

func (m *Map[K, V]) ForKVs(f func(K, V) bool) {
	for c, k := range m.km {
		v := m.m[c]
		if f(k, v) {
			break
		}
	}
}

func (m *Map[K, V]) ForVs(f func(V) bool) {
	for c := range m.km {
		v := m.m[c]
		if f(v) {
			break
		}
	}
}

func (m *Map[K, V]) Delete(k K) {
	c := k.Code()
	delete(m.m, c)
	delete(m.km, c)
}

func (m *Map[K, V]) Set(k K, v V) {
	c := k.Code()
	m.m[c] = v
	m.km[c] = k
}

func (m *Map[K, V]) Get(k K) V {
	return m.m[k.Code()]
}

func (m *Map[K, V]) GetB(k K) (V, bool) {
	v, ok := m.m[k.Code()]
	return v, ok
}

func (m *Map[K, V]) Contains(k K) bool {
	_, ok := m.m[k.Code()]
	return ok
}

func (m *Map[K, V]) Len() int {
	return len(m.m)
}

type Set[K Mappable] struct {
	m *Map[K, bool]
}

func (s *Set[K]) String() string {
	var r []string
	s.For(func(k K) bool {
		r = append(r, fmt.Sprintf("%v", k))
		return false
	})
	return fmt.Sprintf("{%v}", strings.Join(r, ", "))
}

func NewSimpleSet[T comparable](ts ...T) map[T]bool {
	m := map[T]bool{}
	for _, t := range ts {
		m[t] = true
	}
	return m
}

func NewSet[K Mappable](ks ...K) *Set[K] {
	s := &Set[K]{m: NewMap[K, bool]()}
	for _, k := range ks {
		s.Add(k)
	}
	return s
}

func (s *Set[K]) For(f func(K) bool) {
	s.m.ForKs(f)
}

func (s *Set[K]) Add(k K) {
	s.m.Set(k, true)
}

func (s *Set[K]) Delete(k K) {
	s.m.Delete(k)
}

func (s *Set[K]) Contains(k K) bool {
	return s.m.Contains(k)
}

func (s *Set[K]) Len() int {
	return s.m.Len()
}

func Palindrome(n int) bool {
	s := strconv.Itoa(n)
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

/* Use these when https://github.com/golang/go/issues/47619 is done
func SortedMap[K comparable, V any](m map[K]V) []V {
	var ks []K
	for k := range m {
		ks = append(ks, k)
	}
	sort.SliceOf(ks)
	var vs []V
	for _, k := range ks {
		vs = append(vs, m[k])
	}
	return vs
}

func SortedKeys[K comparable, V any](m map[K]V) []K {
	var ks []K
	for k := range m {
		ks = append(ks, k)
	}
	sort.SliceOf(ks)
	return ks
}*/

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
	return NEQ[*Int](i, that)
}

func (i *Int) GT(that *Int) bool {
	return GT[*Int](i, that)
}

func (i *Int) GTE(that *Int) bool {
	return GTE[*Int](i, that)
}

func (i *Int) LTE(that *Int) bool {
	return LTE[*Int](i, that)
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

func BigMin(is []*Int) *Int {
	if len(is) == 0 {
		return Zero()
	}
	min := is[0]
	for _, i := range is {
		if i.LT(min) {
			min = i
		}
	}
	return min
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

var (
	powCache = map[int][]*Int{}
)

func BigPow(a, b int) *Int {
	if b == 0 {
		return One()
	}

	// check cache
	//fmt.Println(powCache[a])
	var start *Int
	if r, ok := powCache[a]; ok {
		if b < len(r) {
			//if !r[b].EQ(OldBigPow(a, b)) {
			//				panic(fmt.Sprintf("%d %d: %v", a, b, r))
			//}
			return r[b].Copy()
		}
		start = r[len(r)-1]
		//fmt.Println("thanks cache", b, len(r))
		b = b + 1 - len(r)
	} else {
		start = One()
		powCache[a] = []*Int{start}
	}

	ai := NewInt(int64(a))
	for i := 1; i <= b; i++ {
		start = start.Times(ai)
		powCache[a] = append(powCache[a], start.Copy())
	}
	//fmt.Println(powCache[a], start)
	return start.Copy()
}

func Sort(is []*Int) {
	sort.SliceStable(is, func(i, j int) bool {
		return is[i].LTE(is[j])
	})
}

func OldBigPow(a, b int) *Int {
	if b == 0 {
		return One()
	}

	start := One()
	ai := NewInt(int64(a))
	for i := 1; i <= b; i++ {
		start = start.Times(ai)
		//powCache[a] = append(powCache[a], start)
	}
	return start.Copy()
}

// Range returns an empty int slice of length n
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

func Digits(n int) []int {
	var r []int
	for v, i := strconv.Itoa(n), 0; i < len(v); i++ {
		r = append(r, parse.Atoi(v[i:i+1]))
	}
	return r
}

func DigitMap(n int) map[int]int {
	m := map[int]int{}
	for v, i := strconv.Itoa(n), 0; i < len(v); i++ {
		m[parse.Atoi(v[i:i+1])]++
	}
	return m
}

func QuadraticRoots(a, b, c float64) []float64 {
	root := b*b - 4*a*c
	if root < 0 {
		return nil
	}
	return []float64{
		(-b - math.Sqrt(root)) / (2 * a),
		(-b + math.Sqrt(root)) / (2 * a),
	}
}

func FromDigits(digits []int) int {
	n := 0
	coef := 1
	for i := len(digits) - 1; i >= 0; i-- {
		n += coef * digits[i]
		coef *= 10
	}
	return n
}

// Anagrams returns all anagram integers of n, not including numbers with leading zeroes.
func Anagrams(n int) map[int]bool {
	r := map[int]bool{}
	digits := Digits(n)
	for _, p := range Permutations(digits, len(digits), false) {
		if p[0] != 0 {
			r[FromDigits(p)] = true
		}
	}
	return r
}

func Anagram(j, k int) bool {
	jm := DigitMap(j)
	km := DigitMap(k)
	if len(jm) != len(km) {
		return false
	}
	for k, v := range jm {
		if v != km[k] {
			return false
		}
	}
	return true
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

func Reverse[T any](ts []T) []T {
	st := make([]T, len(ts))
	for i, v := range ts {
		st[len(ts)-1-i] = v
	}
	return st
}

func CopySlice[T any](ts []T) []T {
	var r []T
	for _, t := range ts {
		r = append(r, t)
	}
	return r
}

func CopyMap[K comparable, V any](m map[K]V, except ...K) map[K]V {
	ignore := map[K]bool{}
	for _, k := range except {
		ignore[k] = true
	}
	r := map[K]V{}
	for k, v := range m {
		if !ignore[k] {
			r[k] = v
		}
	}
	return r
}

type Intable interface {
	ToInt() int
}

func SumType[T Intable](ts []T) int {
	var sum int
	for _, t := range ts {
		sum += t.ToInt()
	}
	return sum
}

// TODO: map package
func Insert[K1, K2 comparable, V any](m map[K1]map[K2]V, k1 K1, k2 K2, v V) {
	if m[k1] == nil {
		m[k1] = map[K2]V{}
	}
	m[k1][k2] = v
}

func DeepInsert[K1, K2, K3 comparable, V any](m map[K1]map[K2]map[K3]V, k1 K1, k2 K2, k3 K3, v V) {
	if m[k1] == nil {
		m[k1] = map[K2]map[K3]V{}
	}
	if m[k1][k2] == nil {
		m[k1][k2] = map[K3]V{}
	}
	m[k1][k2][k3] = v
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

func (i *Int) Part(idx int) int {
	return int(i.parts[idx])
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

// DividingPeriod returns the integer and decimal part of num/den.
// The boolean return value is whether or not the decimal is repeating.
func DividingPeriod(num, den int) (int, []int, bool) {
	checked := map[int]bool{}
	quo := num / den
	var r []int
	rem := num % den
	for rem != 0 && !checked[rem] {
		checked[rem] = true
		r = append(r, (rem*10)/den)
		rem = (rem * 10) % den
	}

	return quo, r, rem != 0
}

/*type DeepMap[K constraints.Ordered, V any] struct {
	m     map[K]V
	next  map[K]DeepMap[K, V]
	depth int
}

type DeepSet

*/

// SquareRootPeriod returns the whole integer and then repeating decimal
// of the SquareRootPeriod
func SquareRootPeriod(n int) (int, []int) {
	if IsSquare(n) {
		return Sqrt(n), nil
	}

	remainder := map[int]map[int]bool{}
	start := 1
	for i := 1; i*i < n; i++ {
		start = i
	}
	num := 1
	den := start
	var as []int
	for !remainder[num][den] && num != 0 {
		Insert(remainder, num, den, true)
		tmpDen := (n - den*den) / num
		newNum := den
		count := 0
		for ; (start + newNum) >= tmpDen; newNum -= tmpDen {
			count++
		}
		as = append(as, count)
		num, den = tmpDen, -newNum
	}
	return start, as
}

func Biggify(is []int) []*Int {
	var r []*Int
	for _, i := range is {
		r = append(r, NewInt(int64(i)))
	}
	return r
}

func Cumulative(is []int) []int {
	if len(is) == 0 {
		return nil
	}
	rs := []int{is[0]}
	for i := 1; i < len(is); i++ {
		rs = append(rs, is[i]+rs[i-1])
	}
	return rs
}
