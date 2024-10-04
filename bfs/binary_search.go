package bfs

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// BinarySearch searches for the target for an unbounded set of indices (mapped by evaluateFunc). If it finds an exact match
// it returns (matchingIdx, true). Otherwise, it returns the first element that has a value larger than the target.
//
// If it could not find a match (either due to invalid arguments or because the solution is above the maximum integer value),
// then it panics.
//
// (We could return -1 in those cases, but I prefer failures of this to be really noisy).
func BinarySearch[T constraints.Integer | constraints.Float](start int, target T, evaluteFunc func(int) T) (int, bool) {
	if start < 0 {
		panic(fmt.Sprintf("invalid start=%d", start))
	}

	left, leftValue := start, evaluteFunc(start)
	if leftValue > target {
		panic(fmt.Sprintf("invalid start=%d; startValue=%v; target=%v", start, leftValue, target))
	} else if leftValue == target {
		return left, true
	}

	right, rightValue := 2*left, evaluteFunc(2*left)
	if left == 0 {
		right, rightValue = 1, evaluteFunc(1)
	}
	for ; rightValue < target; left, right, rightValue = right, right*2, evaluteFunc(right*2) {
	}
	if rightValue == target {
		return right, true
	}

	// Now we know that leftValue < target < rightValue
	// Continue to narrow in
	for middle := (left + right) / 2; left+1 < right; middle = (left + right) / 2 {
		middleValue := evaluteFunc(middle)
		if middleValue < target {
			left = middle
		} else if middleValue > target {
			right = middle
		} else {
			return middle, true
		}
	}

	// If here, then we know that [right = left + 1] and that [leftValue < target < rightValue]
	return right, false
}
