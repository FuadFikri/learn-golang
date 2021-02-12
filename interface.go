package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func sayHello(hasName HasName) {
	fmt.Println("hello ", hasName.GetName())
}
func main() {
	var person1 Person
	person1.Name = "Fikri"
	sayHello(person1)

	cat := Animal{
		Name: "Push",
	}

	sayHello(cat)
}
