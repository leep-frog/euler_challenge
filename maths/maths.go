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
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/functional"
	"golang.org/x/exp/maps"
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
	bread.Operable
}

type mathableOperator[T Mathable] struct {
	m T
}

func newMo[T Mathable](t T) *mathableOperator[T] {
	return &mathableOperator[T]{t}
}

func (mo *mathableOperator[T]) Plus(that *mathableOperator[T]) *mathableOperator[T] {
	return newMo(mo.m + that.m)
}

func (mo *mathableOperator[T]) Times(that *mathableOperator[T]) *mathableOperator[T] {
	return newMo(mo.m * that.m)
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

func IntSquareRoot(i int) int {
	return int(math.Sqrt(float64(i)))
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

var (
	zero = Zero()
)

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

func Chop(n, from, to int) int {
	s := fmt.Sprintf("%d", n)
	from = Max(0, from)
	to = Min(len(s), to)
	return parse.Atoi(s[from:to])
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

func (i *Int) Palindrome() bool {
	s := i.String()
	for idx := range s {
		if s[idx:idx+1] != s[len(s)-idx-1:len(s)-idx] {
			return false
		}
	}
	return true
}

// TODO: separate package for this and other common types/helpers
type Mappable interface {
	Code() string // TODO Change to Hash
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

// TODO: Invert f return value behavior
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

// Intersection returns a set (map[T]bool) that contains keys that are
// present (and have a true value) in all of the provided maps.
func Intersection[T comparable](ms ...map[T]bool) map[T]bool {
	if len(ms) == 0 {
		return map[T]bool{}
	}

	if len(ms) == 1 {
		return maps.Clone(ms[0])
	}

	m := map[T]bool{}
	for _, k := range maps.Keys(ms[0]) {
		if functional.All(ms, func(o map[T]bool) bool { return o[k] }) {
			m[k] = true
		}
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

func (s *Set[K]) Add(ks ...K) {
	for _, k := range ks {
		if s == nil {
			fmt.Println("NOOO1")
		}
		if s.m == nil {
			fmt.Println("NOOO2")
		}
		s.m.Set(k, true)
	}
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

func Reverse(i int) int {
	neg := i < 0
	if neg {
		i = -i
	}
	var rev int
	prod := 1
	for _, d := range Digits(i) {
		rev += prod * d
		prod *= 10
	}

	if neg {
		return -rev
	}
	return rev
}

func (i *Int) Reverse() *Int {
	r := IntFromDigits(bread.Reverse(i.Digits()))
	if i.Negative() {
		return r.Negation()
	}
	return r
}

func Sum(is ...*Int) *Int {
	if len(is) == 0 {
		return Zero()
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

func (i *Int) PP() {
	i.i = i.Plus(One()).i
}

func (i *Int) MM() {
	i.i = i.Minus(One()).i
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

var (
	powCache   = map[int][]*Int{}
	hexLetters = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
	intToHex   = func() map[string]int {
		r := map[string]int{}
		for i, v := range hexLetters {
			r[v] = i
		}
		return r
	}()
)

func ToHex(i int) string {
	return NewInt(i).Hex()
}

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

	ai := NewInt(a)
	for i := 1; i <= b; i++ {
		start = start.Times(ai)
		powCache[a] = append(powCache[a], start.Copy())
	}
	//fmt.Println(powCache[a], start)
	return start.Copy()
}

func Pow[T Mathable](a, b T) T {
	if b == 0 {
		return 1
	}
	ogA := a
	for i := T(1); i < b; i++ {
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

// TODO: map package
func Insert[K1, K2 comparable, V any](m map[K1]map[K2]V, k1 K1, k2 K2, v V) {
	if m[k1] == nil {
		m[k1] = map[K2]V{}
	}
	m[k1][k2] = v
}

func InsertAppend[K1, K2 comparable, V any](m map[K1]map[K2][]V, k1 K1, k2 K2, v V) {
	if m[k1] == nil {
		m[k1] = map[K2][]V{}
	}
	m[k1][k2] = append(m[k1][k2], v)
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
