package maths

// TODO: Separate package

// This file matches things
type Matchable[CTX, R any] interface {
	Matches(CTX, R) bool
}

type Match[L, R any] struct {
	Left       L
	LeftIndex  int
	Right      R
	RightIndex int
}

func MatchItems[CTX any, L Matchable[CTX, R], R any](ctx CTX, left []L, right []R) []*Match[L, R] {
	if len(left) != len(right) {
		panic("Items to be matched must have the same length")
	}

	size := len(left)

	leftIndices, rightIndices := map[int]bool{}, map[int]bool{}

	// matches[left][right] returns whether left.Matches(right) is true
	var isMatch [][]bool
	for i := 0; i < size; i++ {
		isMatch = append(isMatch, make([]bool, size, size))
		leftIndices[i] = true
		rightIndices[i] = true
	}

	for li, l := range left {
		for ri, r := range right {
			isMatch[li][ri] = l.Matches(ctx, r)
		}
	}

	var matches []*Match[L, R]
	for removed := true; len(leftIndices) > 0 && removed; {
		removed = false

		// leftMatches[i] returns the right matches that work
		leftMatches := map[int][]int{}
		rightMatches := map[int][]int{}
		for li := range leftIndices {
			for ri := range rightIndices {
				if isMatch[li][ri] {
					leftMatches[li] = append(leftMatches[li], ri)
					rightMatches[ri] = append(rightMatches[ri], li)
				}
			}
		}

		var match *Match[L, R]
		for li, lm := range leftMatches {
			if len(lm) == 1 {
				match = &Match[L, R]{left[li], li, right[lm[0]], lm[0]}
				goto NEXT_LOOP
			}
		}

		for ri, rm := range rightMatches {
			if len(rm) == 1 {
				match = &Match[L, R]{left[rm[0]], rm[0], right[ri], ri}
				goto NEXT_LOOP
			}
		}

		// First, see if
	NEXT_LOOP:
		if match != nil {
			removed = true
			delete(leftIndices, match.LeftIndex)
			delete(rightIndices, match.RightIndex)
			matches = append(matches, match)
		}
	}

	if len(leftIndices) > 0 {
		panic("NO")
	}

	return matches
}
