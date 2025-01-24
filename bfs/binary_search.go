package bfs

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type BinarySearchable interface {
	constraints.Integer | constraints.Float
}

// BinarySearch searches for the target value between the provided start and end
// indicies (mapped by evalFunc).
// If it finds an exact match, it returns (matchingIdx, true). Otherwise, it
// returns the first element that has a larger value than the target.
//
// If it could not find a match (either due to invalid arguments or because the
// solution is located outside of the provided indices), then it panics.
func BinarySearch[T BinarySearchable](start, end int, target T, evalFunc func(int) T) (int, bool) {
	left, right := start, end
	if left >= right {
		panic(fmt.Sprintf("start [%d] >= end [%d]", left, right))
	}

	rightValue := evalFunc(right)
	if rightValue == target {
		return right, true
	}

	if rightValue < target {
		panic(fmt.Sprintf("invalid end=%d; endValue=%v; target=%v", end, rightValue, target))
	}

	return binarySearch[T](left, right, target, evalFunc)
}

// UnboundedBinarySearch searches for the target value starting at the provided
// start index and increasing indefinitely (mapped by evalFunc).
// If it finds an exact match, it returns (matchingIdx, true). Otherwise, it
// returns the first element that has a larger value than the target.
//
// If it could not find a match (either due to invalid arguments or because the
// solution is located outside of the provided indices), then it panics.
func UnboundedBinarySearch[T BinarySearchable](start int, target T, evalFunc func(int) T) (int, bool) {
	left := start

	if left < 0 {
		panic(fmt.Sprintf("invalid start=%d", left))
	}

	right, rightValue := 2*left, evalFunc(2*left)
	if left == 0 {
		right, rightValue = 1, evalFunc(1)
	}
	for ; rightValue < target; left, right, rightValue = right, right*2, evalFunc(right*2) {
	}
	if rightValue == target {
		return right, true
	}

	return binarySearch[T](left, right, target, evalFunc)
}

// binarySearch assumes that target < rightValue
func binarySearch[T BinarySearchable](left, right int, target T, evalFunc func(int) T) (int, bool) {

	leftValue := evalFunc(left)
	if leftValue > target {
		panic(fmt.Sprintf("invalid start=%d; startValue=%v; target=%v", left, leftValue, target))
	} else if leftValue == target {
		return left, true
	}

	// After the above logic, we know that leftValue < target < rightValue
	// Continue to narrow in
	for middle := (left + right) / 2; left+1 < right; middle = (left + right) / 2 {
		middleValue := evalFunc(middle)
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
