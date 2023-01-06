package maths

// TwoSum returns two elements from `values` that add up to `k`.
// Repeat values are not allowed.
func TwoSum(target int, values []int) (int, int, bool) {
	m := map[int]bool{}
	for _, a := range values {
		if m[a] && 2*a != target {
			return a, target - a, true
		}
		m[target-a] = true
	}
	return 0, 0, false
}

// ThreeSum returns three elements from `values` that add up to `target`.
// Repeat values are not allowed.
func ThreeSum(target int, values []int) (int, int, int, bool) {
	if len(values) < 3 {
		return 0, 0, 0, false
	}
	for i, a := range values[:len(values)-2] {
		if b, c, ok := TwoSum(target-a, values[i+1:]); ok {
			return a, b, c, true
		}
	}
	return 0, 0, 0, false
}
