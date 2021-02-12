package main

import "fmt"

func Ups(i int) interface{} {
	// bisa return int
	// bisa return boolean
	// bisa return apapun

	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "halo"
	}
}

// bisa menerima parameter dalam tipedata apapun
func terserah(apapun interface{}) string {
	return "halo halo"
}

func main() {

	var data interface{} = Ups(4)
	fmt.Println(data)

	var data2 interface{} = terserah(true)
	fmt.Println(data2)

}
