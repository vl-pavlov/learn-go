package main

type number interface {
	~int | ~float64
}

type elm[T number] struct {
	val T
}

func NewNumber[T number](value T) elm[T] {
	return elm[T]{
		val: value,
	}
}

func (e elm[T]) toFloat() float64 {
	return float64(e.val)
}
