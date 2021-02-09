package main

import "fmt"

func main() {

	fmt.Println("halo")

	months := [...]string{
		"Januari",
		"Februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"seprtember",
		"oktober",
		"november",
		"desember",
	}

	fmt.Println(months)

	slice1 := months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice1[0] = "jos"

	fmt.Println(slice1)
	fmt.Println(months)

	slice1 = append(slice1, "oaoe", "aaa")
	fmt.Println(slice1)
	fmt.Println(months)

	fmt.Println("======================")

	slice2 := months[2:4]
	fmt.Println(slice2)

	slice3 := append(slice2, "Fikri")
	fmt.Println(slice3)

	slice3[1] = "fuad"
	fmt.Println(slice3)
	fmt.Println(months)
}
