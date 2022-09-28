package main

import "fmt"

func main() {
	greet("Pia", "Jl. Merdeka Barat, kecamatan Purbolinggo, Lampung Timur")
}

func greet(name, address string) {
	fmt.Println("Halo semuanya! Saya", name)
	fmt.Println("Saat ini saya tinggal di", address)
}
