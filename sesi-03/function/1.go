package main

import "fmt"

func main() {
	greet("Pia", 21)
}

func greet(name string, age int8) {
	fmt.Printf("Salam kenal semua! Saya %s dan umur saya %d tahun", name, age)
}
