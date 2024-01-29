package main

import "fmt"

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

type summarizable interface {
	toFloat() float64
}

type list struct {
	elms []summarizable
}

func NewList() list {
	return list{
		elms: []summarizable{},
	}
}

func (l *list) Sum() float64 {
	result := 0.0

	for _, elm := range l.elms {
		result += elm.toFloat()
	}

	return result
}

func (l *list) Add(elm summarizable) {
	l.elms = append(l.elms, elm)
}

func main() {
	list := NewList()

	list.Add(NewNumber(1))
	list.Add(NewNumber(1.2))

	// Error
	// list.Add(NewNumber("1.2"))

	fmt.Println("Sum: ", list.Sum())
}
