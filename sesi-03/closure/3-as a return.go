package main

import (
	"fmt"
	"strings"
)

func main() {
	var petLists = []string{"Ichi", "Nini", "Yoyo", "Ucrit", "Moi"}

	var find = findpet(petLists)

	fmt.Println(find("Yoyo"))

}

func findpet(pets []string) func(string) string {

	return func(s string) string {
		var pet string
		var position int

		for i, v := range pets {
			if strings.ToLower(v) == strings.ToLower(s) {
				pet = v
				position = i
				break
			}
		}
		if pet == "" {
			return fmt.Sprintf("%s tidak ada", s)
		}
		return fmt.Sprintf("ada %s di %d", s, position+1)
	}

}
