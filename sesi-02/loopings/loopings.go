package main

import "fmt"

/*func main () {
	//praktek hal 2 (first way)
	for i := 0; i < 3; i++ {
		fmt.Println("Angka", i)
	}

fmt.Println("==================")

	//praktek hal 4 (second way)
	var a = 0

	for a < 3 {
		fmt.Println("Nilai", a)
		a++
	}

fmt.Println("==================")

	//praktek hal 5 (third way)
	var o = 0

	for  {
		fmt.Println("Bilangan", o)

		o++
		if o == 3 {
			break
		}
	}
fmt.Println("==================")
}*/

//Break and Continue
/*func main () {
	//praktek hal 7
	for z := 1; z <= 10; z++ {
		if z%2 == 1 {
			continue
		}

		if z > 8 {
			break
		}

		fmt.Println("Katakan", z)
	}
}*/

//Nested
/*func main() {
	//praktek hal 9
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}
}*/

//Label
func main() {
	//praktek hal 10
	outerLoop:
		for i := 0; i < 3; i++ {
			fmt.Println("Perulangan ke - ", i+1)
			for j := 0; j < 3; j++ {
				if i == 2 {
					break outerLoop
				}
				fmt.Print(j, " ")
			}
			fmt.Print("\n")
	}
}