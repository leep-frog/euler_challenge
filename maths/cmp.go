package maths

func LT[T Comparable[T]](this, that T) bool {
	return this.LT(that)
}

func EQ[T Comparable[T]](this, that T) bool {
	return !that.LT(this) && !this.LT(that)
}

func NEQ[T Comparable[T]](this, that T) bool {
	return !EQ[T](that, this)
}

func GT[T Comparable[T]](this, that T) bool {
	return that.LT(this)
}

func LTE[T Comparable[T]](this, that T) bool {
	return !that.LT(this)
}

func GTE[T Comparable[T]](this, that T) bool {
	return LTE[T](that, this)
}