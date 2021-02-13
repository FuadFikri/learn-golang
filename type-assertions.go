package main

import "fmt"

func random() interface{} {
	return "ups"
}

func main() {
	var result interface{} = random()
	var resultString string = result.(string)

	fmt.Println(resultString)

	// var resultBoolean bool = result.(bool) panic

	switch value := result.(type) {
	case string:
		fmt.Println("value ", value, " is string")
	case int:
		fmt.Println("value ", value, "is int")
	default:
		fmt.Println("unknown type")
	}
}
