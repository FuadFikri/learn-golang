package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {

	if pembagi == 0 {
		return 0, errors.New("pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	hasil, err := Pembagian(10, 0)
	if err == nil {
		fmt.Println("hasilnya : ", hasil)

	} else {
		fmt.Println("Error", err.Error())
	}
}
