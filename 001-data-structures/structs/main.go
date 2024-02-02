package main

import (
	"fmt"
	"strconv"
)

type person struct {
	name string
	age  int

	// You can do that too (Nested Structure)
	// location struct {
	//	country, city string
	// }
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

// Method
func (p *person) String() string {
	return p.name + " (" + strconv.Itoa(p.age) + " y.o.)"
}

func main() {
	// This syntax creates a new struct.
	fmt.Println(person{"Bob", 20})

	// You can name the fields when initializing a struct.
	fmt.Println(person{name: "Alice", age: 30})

	// Omitted fields will be zero-valued.

	fmt.Println(person{name: "Fred"})

	// An & prefix yields a pointer to the struct.
	fmt.Println(&person{name: "Ann", age: 40})

	// It’s idiomatic to encapsulate new struct creation in constructor functions
	fmt.Println(newPerson("Jon"))

	// Access struct fields with a dot.
	person := person{name: "Sean", age: 50}
	fmt.Println(person.name)

	// You can also use dots with struct pointers - the pointers are automatically dereferenced.
	personPointer := &person
	fmt.Println(personPointer.age)

	// Structs are mutable.
	personPointer.age = 51
	fmt.Println(personPointer.age)

	// If a struct type is only used for a single value, we don’t have to give it a name. The value can have an anonymous struct type. This technique is commonly used for table-driven tests.
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}

	fmt.Println(dog)
}
