package main

import "fmt"

type person struct {
	first string
}

type secreteAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("from a person - this is my name", p.first)
}

func (sa secreteAgent) speak() {
	fmt.Println("I am a secrete agent - this is my name", sa.first)
}

// any TYPE that has the methods specified by an interface
// is also the interface type
// an interface says
// "Hey baby, if you have these method, then you're my type"

type human interface {
	speak()
}

func foo(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Miss Moneypenny",
	}

	sa1 := secreteAgent{
		person: person{
			first: "James",
		},
		ltk: true,
	}
	fmt.Printf("%T\n", p1)

	// in Go, a VALUE can be of more than one TYPE
	// in this example, p1 is both TYPE person and human
	var x, y human
	x = p1
	y = sa1
	x.speak()
	y.speak()
	fmt.Printf("--------")
	foo(x)
	foo(y)
	foo(p1)
	foo(sa1)
}
