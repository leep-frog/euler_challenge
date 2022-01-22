package maths

func Min[T Mathable](as ...T) T {
	var min T
	if len(as) == 0 {
		return min
	}
	min = as[0]
	for _, a := range as {
		if a < min {
			min = a
		}
	}
	return min
}

func Max[T Mathable](as ...T) T {
	var max T
	if len(as) == 0 {
		return max
	}
	max = as[0]
	for _, a := range as {
		if a > max {
			max = a
		}
	}
	return max
}