package main

import "fmt"

type Student struct {
	Name, Faculty string
	age           int
}

func (student Student) sayHello(name string) {
	fmt.Println("Hello ", student.Name, " my name is ", name)
}

func main() {

	var student1 Student
	student1.Name = "Fikri"
	student1.age = 20
	student1.Faculty = "Engineering"

	fmt.Println(student1)
	fmt.Println(student1.Name)

	kamu := Student{
		Name:    "fik",
		Faculty: "Engineering",
		age:     20,
	}

	fmt.Println(kamu)

	kamu.sayHello("isna")

}
