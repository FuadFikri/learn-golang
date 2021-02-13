package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	// pass by value

	address1 := Address{"Sleman", "DIY", "Indonesia"}
	address2 := address1

	address2.City = "Jogja"

	fmt.Println(address1) //{Sleman DIY Indonesia}
	fmt.Println(address2) //{Jogja DIY Indonesia}

	// pass by reference

	// Pointer
	// kemampuan membuat reference ke lokasi data di memory yang sama,
	// tanpa menduplikasi data yang sudah ada

	var address3 Address = Address{"Manchester", "Inggris barat", "UK"}
	var address4 *Address = &address3

	address2.City = "Jogja"

	fmt.Println(address3) // {Manchester Inggris barat UK}
	fmt.Println(address4) // &{Manchester Inggris barat UK}

	address4 = &Address{"Malang", "Jawa timur", "Indonesia"}
	fmt.Println(address3) // {Manchester Inggris barat UK}     -> TIDAK BERUBAH
	fmt.Println(address4) //&{Malang Jawa timur Indonesia}     -> mereference object baru di memory

	var address5 *Address = &address3

	// mengubah seluruh variabel yg mereference address 3
	*address5 = Address{"Jakarta", "DKI J", "Indonesia"}
	fmt.Println(address3) // {Jakarta DKI J Indonesia}
	fmt.Println(address5) // &{Jakarta DKI J Indonesia}

	// KEYWORD NEW
	// membuat pointer yg merefer ke data kosong

	var alamat1 *Address = new(Address)
	fmt.Println(alamat1) //&{  }

	alamat1.City = "London"
	fmt.Println(alamat1) //&{London  }

}
