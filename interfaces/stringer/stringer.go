package stringer

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	return fmt.Sprintf("My name is %s. I'm %d years old. ", p.Name, p.Age)
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age: age,
	}
}


