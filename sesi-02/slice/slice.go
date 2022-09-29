package main

import (
	"fmt"
)

/*func main() {
	//praktek hal 2
	var buah = []string{"pisang", "jambu", "apel"}

	_ = buah
}*/

//make function
//praktek hal 3
/*func main() {
	var buah = make([]string, 3)

	_ = buah

	fmt.Printf("%#v", buah)
}*/

//apend function
//praktek hal 4
/*func main() {
	var kendaraan = make([]string, 3)
	_ = kendaraan

	kendaraan[0] = "Honda"
	kendaraan[1] = "Suzuki"
	kendaraan[2] = "Toyota"

	fmt.Printf("%#v", kendaraan)

fmt.Println(" ")
fmt.Println("==========================")

	//cara append
	var pakde = make([]string, 3)

	pakde = append(pakde, "Edi", "Wahono", "Nanang")

	fmt.Printf("%#v", pakde)
}*/

//append function with ellipsis
//praktek hal 5
/*func main() {
	var kakak1	= []string{"Asih", "Isti", "Mirda"}
	var kakak2	= []string{"Edward", "Erik", "Edo"}

	kakak1 = append(kakak1, kakak2...)

	fmt.Printf("%#v", kakak1)
}*/

//copy
//praktek hal 6
/*func main() {
	var adik1	= []string{"Teza", "Arimbi", "Leticia"}
	var adik2	= []string{"Bima", "Putra", "Attala"}

	nn := copy(adik1, adik2)

	fmt.Println("adik1 =>", adik1)
	fmt.Println("adik2 =>", adik2)
	fmt.Println("Copied elements =>", nn)
}*/

//slicing
//praktek hal 8
/*func main() {
	var bunga1 = []string{"mawar", "melati", "kamboja", "kenanga", "anggrek"}

	var bunga2 = bunga1[1:4]
	fmt.Printf("%#v\n", bunga2)

	var bunga3 = bunga1[0:]
	fmt.Printf("%#v\n", bunga3)

	var bunga4 = bunga1[:3]
	fmt.Printf("%#v\n", bunga4)

	var bunga5 = bunga1[:]
	fmt.Printf("%#v\n", bunga5)
}*/

//combine slicing dan append
//praktek hal 11
/*func main() {
	var hewan1 = []string{"kucing", "bebek", "kelinci", "ayam", "burung"}

	hewan1 = append(hewan1[:3], "anjing")

	fmt.Printf("%#v\n", hewan1)
}*/

//backing array
//praktek hal 13
/*func main() {
	var sayur1 = []string{"tomat", "timun", "terong", "sawi", "bayam"}
	var sayur2 = sayur1[2:4]

	sayur2[0] = "cabai"

	fmt.Println("sayur1 => ", sayur1)
	fmt.Println("sayur2 => ", sayur2)
}*/

//cap function
//praktek hal 15
/*func main() {
	var benda1 = []string{"buku", "bola", "bandana", "bom"}

	fmt.Println("benda1 cap:", cap(benda1))
	fmt.Println("benda1 len:", len(benda1))

	fmt.Println(strings.Repeat("#", 20))

	var benda2 = benda1[0:3]

	fmt.Println("benda2 cap:", cap(benda2))
	fmt.Println("benda2 len:", len(benda2))

	fmt.Println(strings.Repeat("#", 20))

	var benda3 = benda1[1:]

	fmt.Println("benda3 cap:", cap(benda3))
	fmt.Println("benda3 len:", len(benda3))
}*/

//creating a new backing array
//praktek hal 18
func main() {
	mobil := []string{"pajero", "agya", "fortuner", "mobilio"}
	mobilBaru := []string{}

	mobilBaru = append(mobilBaru, mobil[0:2]...)

	mobil[0] = "truk"
	fmt.Println("mobil:", mobil)
	fmt.Println("mobil baru:", mobilBaru)
}