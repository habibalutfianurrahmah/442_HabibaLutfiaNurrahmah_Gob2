package main

/*import "fmt"*/

/*func main() {
	//praktek hal 2
	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}

	var strings = [3]string{"Shiro", "Yuki", "Kiiro"}

	fmt.Printf("%#v\n", numbers)
	fmt.Printf("%#v\n", strings)
}*/

//modify element through index
/*func main() {
	//praktek hal 4
	var warna = [3]string{"krem", "putih", "hitam"}
	warna[0]	= "krem"
	warna[1]	= "putih"
	warna[2]	= "hitam"

	fmt.Printf("%#v\n", warna)
}*/

//loop through elements
//praktek hal 5
/*import (
	"fmt"
	"strings"
)

func main() {
	var kucing	= [3]string{"Dusty", "Leah", "Loli"}

	//1
	for c, d := range kucing {
		fmt.Printf("Index: %d, Value: %s\n", c, d)
	}

	fmt.Println(strings.Repeat("#", 25))

	//2
	for p := 0; p < len(kucing); p++ {
		fmt.Printf("Index : %d, Value: %s\n", p, kucing[p])
	}
}*/

//multidimensional
//praktek hal 6
import (
	"fmt"
)

func main() {
	
	balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}

}