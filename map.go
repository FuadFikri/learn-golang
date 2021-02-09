package main

import "fmt"

func main() {

	person := map[string]string{
		"name": "fikri",
		"age":  "20",
	}

	fmt.Println(person["name"])
	fmt.Println(person["age"])

	book := make(map[string]string)
	book["title"] = "Harry Potter"
	book["author"] = "rowling"
	book["language"] = "english"

	name := make(map[int8]string)
	name[0] = "fikri"

	fmt.Println(name)

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "language")

	fmt.Println(len(book))
	fmt.Println(book)
}
