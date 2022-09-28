package main

import "fmt"

func main() {
	var fristnumber int = 4
	var secondnumber int = &fristnumber
	
	fmt.Println("fristnumber (value) :", fristnumber)
	fmt.Println("fristnumber (memori address) :", &fristnumber)

	fmt.Println("secondnumber (value) :", secondnumber)
	fmt.Println("secondnumber (memori address) :", secondnumber)
}