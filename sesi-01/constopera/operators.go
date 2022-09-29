package main

import "fmt"

func main() {
	var nilai = 2 + 2*3
	fmt.Println(nilai)

//praktek hal 6 (relational)
	var situasiSatu bool = 2 < 3
	var situasiDua bool = "Pia" == "Mooi"
	var situasiTiga bool = 10 != 2.3
	var situasiEmpat bool = 11 <= 11

	fmt.Println("Keadaan satu:", situasiSatu)
	fmt.Println("Keadaan dua:", situasiDua)
	fmt.Println("Keadaan tiga:", situasiTiga)
	fmt.Println("Keadaan empat:", situasiEmpat)

//praktek hal 10 (logical)
	var right	= true
	var wrong	= false

	var benardansalah	= wrong && right
	fmt.Printf("wrong && right \t(%t) \n", benardansalah)

	var benaratausalah	= wrong || right
	fmt.Printf("wrong || right \t(%t) \n", benaratausalah)

	var terbaliksalah	= !wrong
	fmt.Printf("!wrong \t\t(%t) \n", terbaliksalah)
}

