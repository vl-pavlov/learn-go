package main

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
