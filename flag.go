package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "defaultnya", "put your host")
	var user *string = flag.String("user", "defaultnya", "put your user")
	var password *string = flag.String("password", "defaultnya", "put your password")

	flag.Parse()

	fmt.Println("Host : ", *host)
	fmt.Println("User : ", *user)
	fmt.Println("Password : ", *password)

	// kita bisa run di terminal dengan mengirim value
	// go run flag.go -host=localhost -user=root -password=rahasia
}
