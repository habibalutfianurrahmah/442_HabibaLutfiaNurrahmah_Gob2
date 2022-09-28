package main

import (
	"fmt"
	"strings"
)

func main() {
	profile("Pia", "Ichi", "Nini", "Yoyo", "Ucrit", "Moi")
}

func profile(name string, myPets ...string) {
	mergemyPets := strings.Join(myPets, ",")

	fmt.Println("Holaa guys! Aku", name)
	fmt.Println("Aku memiliki beberapa kucing peliharaan, namanya", mergemyPets)
}
