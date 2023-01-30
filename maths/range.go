package maths

import (
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

type Range struct {
	// Alternating integers, where each integer flips
	// the presence of range. For example: []int{3, 7, 12, 13}
	// includes the numbers [3, 4, 5, 6, 7, 12, 13]
	inflectionPoints []int
}

func NewRange(vs ...int) *Range {
	r := &Range{vs}
	r.verifyAndSimplify()
	return r
}

func (r *Range) String() string {
	var s []string
	for i := 0; i < len(r.inflectionPoints); i += 2 {
		s = append(s, fmt.Sprintf("%d-%d", r.inflectionPoints[i], r.inflectionPoints[i+1]))
	}

	return fmt.Sprintf("[%v]", strings.Join(s, ";"))
	// return fmt.Sprintf("[%d])
}

func (r *Range) Contains(k int) bool {
	// idx, _ := slices.BinarySearch(r.inflectionPoints, k)
	idx, _ := slices.BinarySearch(r.inflectionPoints, k)

	if idx >= len(r.inflectionPoints) {
		return false
	}

	// k is at an inflection point, so valid
	if k == r.inflectionPoints[idx] {
		return true
	}

	// k is between an inflection point
	return idx%2 == 1
}

func (r *Range) verifyAndSimplify() {
	if !slices.IsSorted(r.inflectionPoints) {
		panic("Range.inflectioPoints must be sorted")
	}
	if len(r.inflectionPoints)%2 != 0 {
		panic("Range.inflectioPoints requires an even number of things")
	}

	// If there are any adjacent entries, then we have a redundant range.
	var newPoints []int
	for i := 0; i < len(r.inflectionPoints); i++ {
		v := r.inflectionPoints[i]
		if i%2 == 0 {
			newPoints = append(newPoints, v)
		} else {
			if i+1 < len(r.inflectionPoints) && v >= r.inflectionPoints[i+1]-1 {
				i++
			} else {
				newPoints = append(newPoints, v)
			}
		}
	}
	r.inflectionPoints = newPoints
}

func (r *Range) Merge(that *Range) *Range {
	var newPoints []int
	var inRangeCount int
	for i, j := 0, 0; i < len(r.inflectionPoints) || j < len(that.inflectionPoints); {
		var value int
		var enteringRange bool
		if i >= len(r.inflectionPoints) {
			value = that.inflectionPoints[j]
			enteringRange = j%2 == 0
			j++
		} else if j >= len(that.inflectionPoints) {
			value = r.inflectionPoints[i]
			enteringRange = i%2 == 0
			i++
		} else if r.inflectionPoints[i] <= that.inflectionPoints[j] {
			value = r.inflectionPoints[i]
			enteringRange = i%2 == 0
			i++
		} else {
			value = that.inflectionPoints[j]
			enteringRange = j%2 == 0
			j++
		}

		if enteringRange {
			inRangeCount++
			if inRangeCount == 1 {
				newPoints = append(newPoints, value)
			}
		} else {
			inRangeCount--
			if inRangeCount == 0 {
				newPoints = append(newPoints, value)
			}
		}
	}
	merged := &Range{newPoints}
	merged.verifyAndSimplify()
	return merged
}
