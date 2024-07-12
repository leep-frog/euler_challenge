// bread is a package with slice utilities
package bread

import (
	"math/rand"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

type Operable interface {
	~int | ~float64 | ~int64 | ~uint64
}

type Copyable[T any] interface {
	Copy() T
}

func Sum[T Operable](elements []T) T {
	var sum T
	for _, e := range elements {
		sum += e
	}
	return sum
}

func Product[T Operable](elements []T) T {
	sum := T(1)
	for _, e := range elements {
		sum *= e
	}
	return sum
}

func Reverse[T any](ts []T) []T {
	st := make([]T, len(ts))
	for i, v := range ts {
		st[len(ts)-1-i] = v
	}
	return st
}

func DeepCopy[T Copyable[T]](ts []T) []T {
	var r []T
	for _, t := range ts {
		r = append(r, t.Copy())
	}
	return r
}

func Copy[T any](ts []T) []T {
	return slices.Clone(ts)
}

func Zip[T any](slc ...[]T) [][]T {
	var zipped [][]T
	for i := 0; i < len(slc[0]); i++ {
		var r []T
		for _, s := range slc {
			r = append(r, s[i])
		}
		zipped = append(zipped, r)
	}
	return zipped
}

// MergeSort merge sorts the provided arrays. It assumes the
// arrays are sorted.
func MergeSort[T constraints.Ordered](a, b []T, removeDuplicates bool) []T {
	var merged []T
	for ai, bi := 0, 0; ai < len(a) || bi < len(b); {
		var contender T
		if ai == len(a) {
			contender = b[bi]
			bi++
		} else if bi == len(b) {
			contender = a[ai]
			ai++
		} else if a[ai] <= b[bi] {
			contender = a[ai]
			ai++
		} else {
			contender = b[bi]
			bi++
		}
		if !removeDuplicates || len(merged) == 0 || contender != merged[len(merged)-1] {
			merged = append(merged, contender)
		}
	}
	return merged
}

func Shuffle[T any](items []T) {
	for i := 0; i < len(items)-1; i++ {
		// Pick a random number between
		swapIdx := i + (rand.Int() % (len(items) - i))
		if swapIdx == i {
			continue
		}
		items[i], items[swapIdx] = items[swapIdx], items[i]
	}
}
