package combinatorics

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/functional"
	"golang.org/x/exp/slices"
)

var (
	permutationCountCache = map[string]generator.PrimeFactoredNumber{}
	p                     = generator.Primes()
)

func PermutationCount[T any](parts []T) *maths.Int {
	counts, _ := createCounts(parts)
	return PermutationFromCount(counts)
}

func PermutationFromCount(ogCounts []int) *maths.Int {
	return PermutationFromCountNew(ogCounts).ToBigInt(p)
}

func PermutationFromCountNew(ogCounts []int) generator.PrimeFactoredNumber {
	counts := bread.Copy(ogCounts)
	slices.Sort(counts)

	code := strings.Join(functional.Map(counts, strconv.Itoa), "_")
	if v, ok := permutationCountCache[code]; ok {
		return v
	}

	// Total count if all elements were different
	v := p.PrimeFactoredNumberFactorial(bread.Sum(counts))

	for _, c := range counts {
		if c == 1 {
			continue
		}
		v = v.Div(p.PrimeFactoredNumberFactorial(c))
	}

	permutationCountCache[code] = v
	return v
}

func StringPermutations(parts []string) []string {
	var r []string
	for _, perm := range Permutations(parts) {
		r = append(r, strings.Join(perm, ""))
	}
	return r
}

func Permutations[T any](parts []T) [][]T {
	return PermutationsOfLength(parts, len(parts))
}

func PermutationsOfLength[T any](parts []T, length int) [][]T {
	return GenerateCombos(&Combinatorics[T]{
		Parts:            parts,
		MinLength:        length,
		MaxLength:        length,
		AllowReplacement: false,
		OrderMatters:     true,
	})
}

func ChooseAllSets[T any](parts []T) [][]T {
	return GenerateCombos(&Combinatorics[T]{
		Parts:            parts,
		MinLength:        1,
		MaxLength:        len(parts),
		AllowReplacement: false,
		OrderMatters:     false,
	})
}

func ChooseSets[T any](parts []T, minLength, maxLength int) [][]T {
	return GenerateCombos(&Combinatorics[T]{
		Parts:            parts,
		MinLength:        minLength,
		MaxLength:        maxLength,
		AllowReplacement: false,
		OrderMatters:     false,
	})
}

func ChooseSetsOfLength[T any](parts []T, length int) [][]T {
	return ChooseSets(parts, length, length)
}

// Anagrams returns all anagram integers of n, not including numbers with leading zeroes.
func Anagrams(n int) map[int]bool {
	r := map[int]bool{}
	digits := maths.Digits(n)
	for _, p := range Permutations(digits) {
		if p[0] != 0 {
			r[maths.FromDigits(p)] = true
		}
	}
	return r
}

func Anagram(j, k int) bool {
	jm := maths.DigitMap(j)
	km := maths.DigitMap(k)
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

type Combinatorics[T any] struct {
	Parts            []T
	MinLength        int
	MaxLength        int
	AllowReplacement bool
	OrderMatters     bool
}

func createCounts[T any](parts []T) ([]int, []T) {
	var realParts []T
	var counts []int
	indexMap := map[string]int{}
	for _, part := range parts {
		if i, ok := indexMap[fmt.Sprintf("%v", part)]; ok {
			counts[i]++
		} else {
			indexMap[fmt.Sprintf("%v", part)] = len(counts)
			counts = append(counts, 1)
			realParts = append(realParts, part)
		}
	}
	return counts, realParts
}

func GenerateCombos[T any](c *Combinatorics[T]) [][]T {
	var all [][]T
	EvaluateCombos(c, func(t []T) { all = append(all, bread.Copy(t)) })
	return all
}

func EvaluateCombos[T any](c *Combinatorics[T], applyFunc func([]T)) {
	counts, realParts := createCounts(c.Parts)
	generateCombos(c, counts, realParts, 0, nil, applyFunc)
}

func generateCombos[T any](c *Combinatorics[T], counts []int, realParts []T, minIndex int, cur []T, applyFunc func([]T)) {
	if c.MinLength <= len(cur) && len(cur) <= c.MaxLength && len(cur) > 0 {
		applyFunc(cur)
	}

	if len(cur) >= c.MaxLength {
		return
	}

	start := 0
	if !c.OrderMatters {
		start = minIndex
	}

	for k := start; k < len(counts); k++ {
		if counts[k] > 0 || c.AllowReplacement {
			counts[k]--
			generateCombos(c, counts, realParts, k, append(cur, realParts[k]), applyFunc)
			counts[k]++
		}
	}
}

// Rotations returns all rotations of the elements of parts.
// For example, if parts is `["ab", "cd", "ef"]`, then this will return
// `["abcdef", "cdefab", "efabcd"]`.
func Rotations(parts []string) []string {
	var r []string
	for i := 0; i < len(parts); i++ {
		r = append(r, strings.Join(append(parts[i:], parts[:i]...), ""))
	}
	return r
}
