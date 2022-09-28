package main

import (
	"fmt"
)

func main() {
	studentLists := print("Ichi", "Nini", "Yoyo", "Ucrit", "Moi")

	fmt.Printf("%v", studentLists)
}

func print(names ...string) []map[string]string {
	var result []map[string]string

	for i, v := range names {
		key := fmt.Sprintf("student%d", i+1)
		tempt := map[string]string{
			key: v,
		}
		result = append(result, tempt)
	}

	return result
}
