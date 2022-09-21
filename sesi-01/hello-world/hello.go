package main

import "fmt"

func main() {
	fmt.Println("Hello")
	fmt.Println("Pia", "unyu", "unyu")
	fmt.Println("==============")

	/* Praktek Hal 28 */

	fmt.Println("Hello")
	fmt.Println("Hello", "Pia")

	fmt.Print("Pia ayo mandi\n")
	fmt.Print("Pia", " ", "ayo", " mandi\n")
	fmt.Print("Pia", " ", "mandi\n")
	fmt.Println("==============")

	/* Praktek Variabel */

	/*dengan tipe data string-init*/
	var namalengkap string = "Habiba Lutfia"
	var umur int = 21

	fmt.Println("Dia bernama", namalengkap)
	fmt.Println("Berusia", umur)
	fmt.Println("==============")

	/*tanpa tipe data*/
	var namasaya = "Pia Unyu"
	var umursaya = 21

	fmt.Printf("%T, %T\n", namasaya, umursaya)
	fmt.Println("==============")

	/*short declaration*/
	namaku := "Pia"
	umurku := 21

	fmt.Printf("%T, %T\n", namaku, umurku)
	fmt.Println("==============")

	/*multiple var declaration*/
	var one, two, three string = "satu", "dua", "tiga"
	var pertama, kedua, ketiga int = 1, 2, 3

	fmt.Println("ayo berhitung", one, two, three)
	fmt.Println("sebutkan", pertama, kedua, ketiga)
	fmt.Println("==============")

	/*string*/
	var sapaan string = "Selamat Malam"
	fmt.Println(sapaan)

	var pesan = "Hai aku Pia, salam kenal"
	fmt.Println(pesan)
}
