package main

import (
	"fmt"
)

//looping over string (byte-by-byte)
func main() {
//praktek hal 3
/*	kota := "Jakarta"

	for i := 0; i < len(kota); i++ {
		fmt.Printf("index: %d, bye: %d\n", i, kota[i])
	}
*/

//praktek hal 5
/*	var kota []byte = []byte{74, 97, 107, 97, 114, 116, 97}

	fmt.Println(string(kota))
*/

//praktek hal 7
/*	str1	:= "makan"
	str2	:= "mânca"

	fmt.Printf("str1 byte length => %d\n", len(str1))
	fmt.Printf("str2 byte length => %d\n", len(str2))
*/

//praktek hal 9
/*	str1	:= "makan"
	str2	:= "mânca"

	fmt.Printf("str1 character length => %d\n", utf8.RuneCountInString(str1))
	fmt.Printf("str2 character length => %d\n", utf8.RuneCountInString(str2))
*/

//praktek hal 11
	str := "mânca"

	for i, s := range str{
		fmt.Printf("index => %d, string => %s\n", i, string(s))
	}
}