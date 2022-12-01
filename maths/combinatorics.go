package maths

import "fmt"

func Permutations[T any](parts []T) [][]T {
	return GenerateCombos(&Combinatorics[T]{
		Parts:            parts,
		MinLength:        len(parts),
		MaxLength:        len(parts),
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
	digits := Digits(n)
	for _, p := range Permutations(digits) {
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

type Combinatorics[T any] struct {
	Parts            []T
	MinLength        int
	MaxLength        int
	AllowReplacement bool
	OrderMatters     bool
}

func GenerateCombos[T any](c *Combinatorics[T]) [][]T {
	var counts []int
	var realParts []T
	indexMap := map[string]int{}
	for _, part := range c.Parts {
		if i, ok := indexMap[fmt.Sprintf("%v", part)]; ok {
			counts[i]++
		} else {
			indexMap[fmt.Sprintf("%v", part)] = len(counts)
			counts = append(counts, 1)
			realParts = append(realParts, part)
		}
	}

	var all [][]T
	generateCombos(c, counts, 0, nil, &all)
	return all
}

func generateCombos[T any](c *Combinatorics[T], counts []int, minIndex int, cur []T, all *[][]T) {
	if c.MinLength <= len(cur) && len(cur) <= c.MaxLength && len(cur) > 0 {
		*all = append(*all, CopySlice(cur))
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
			generateCombos(c, counts, k, append(cur, c.Parts[k]), all)
			counts[k]++
		}
	}
}
