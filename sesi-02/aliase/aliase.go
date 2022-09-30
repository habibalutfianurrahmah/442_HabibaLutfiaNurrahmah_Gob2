package main

import "fmt"

func main() {
	//praktek hal 3
	//deklarasi variabel dg tipe data uint8
	/*	var a uint8 = 10
		var b byte

		b = a //tidak eror, karena byte memiliki tipe data yang sama dg uint8
		_ = b
	*/

	//praktek hal 4 (mendeklarasikan tipe data alase > second)
	//tipe nama_alias = nama_tipe_data
	type second = uint

	var hour second = 3600
	fmt.Printf("hour type: %t\n", hour) // hour type itu 'uint'
}