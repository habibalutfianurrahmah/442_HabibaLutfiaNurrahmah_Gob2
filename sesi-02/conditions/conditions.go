package main

import "fmt"

//IF
/*func main() {
	//praktek hal 4
	var tahunIni	= 2022

	if umur := tahunIni - 2012; umur < 17 {
		fmt.Println("Anda belum diperkenankan membuat KTP")
	} else {
		fmt.Println("Anda diperbolehkan membuat KTP")
	}
}*/

//Switch
func main() {
	//praktek hal 6
	var nilai	= 7

	switch nilai {
	case 8:
		fmt.Println("Luar Biasa")
	case 7:
		fmt.Println("Mengesankan")
	default:
		fmt.Println("Bagus")
	}

	//praktek hal 7 (with relational operators)
	var score	= 6

	switch {
	case score == 8:
		fmt.Println("Perfect")
	case (score < 8) && (score > 3):
		fmt.Println("Okay")
	default: {
		fmt.Println("Terus tingkatkan prestasimu")
		fmt.Println("Lebih banyak belajar!")
		}
	}

	//praktek hal 8 (fallthrough)
	var hasil	= 6

	switch {
	case hasil == 8:
		fmt.Println("Perfect")
	case (score < 8) && (score > 3):
		fmt.Println("Kerja bagus")
	fallthrough
	case hasil < 5:
		fmt.Println("It's okay, kamu sudah cukup baik")
	default: {
		fmt.Println("Terus tingkatkan prestasimu")
		fmt.Println("Lebih banyak belajar!")
		}
	}

	//praktek hal 9 (nested)
	var bijih	= 10

	if bijih > 7 {
		switch bijih {
		case 10 :
			fmt.Println("Sangat baik!")
		default:
			fmt.Println("Kerja bagus!")
		}
	} else {
		if bijih == 5 {
			fmt.Println("Bagus!")
		} else if bijih == 3 {
			fmt.Println("Tingkatkan prestasimu!")
		} else {
			fmt.Println("Belajar lagi!")
			if bijih == 0 {
				fmt.Println("Kamu mendapat kelas tambahan!")
			}
		}
	}
}