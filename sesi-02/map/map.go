package main

import "fmt"

/*func main() {
	//praktek hal 3
	var person map [string]string	//deklarasi

	person = map[string]string{}	//inisialisasi

	person["nama"]	= "Isdet"
	person["usia"]	= "28"
	person["sifat"]	= "pemarah dan jutek"

	fmt.Println("namanya adalah", person["nama"])
	fmt.Println("dia berumur", person["usia"], "tahun")
	fmt.Println("sifatnya sangat", person["sifat"])
}*/

//looping with map
//praktek hal 5
/*func main() {
	var person = map[string]string{
		"nama panggilan":	"mbak Isdet",
		"hobi"			:	"membaca buku",
		"sifat"			:	"sangat pemarah dan jutek",
	}

	for key, value := range person {
		fmt.Println(key, ":", value)
	}
}*/

//deleting value
//praktek hal 6
/*func main() {
	var person	= map[string]string{
		"nama lengkap":	"Habiba Nurul Istiqomah",
		"pekerjaan" :	"digital marketer",
		"pendidikan" : "S1 Agroteknologi Unila 2015",
	}

	fmt.Println("Before deleting:", person)

	delete(person, "pekerjaan")

	fmt.Println("After deleting:", person)
}*/

//detecting a value
//praktek hal 7
/*func main() {
	var person = map[string]string{
		"panggilan": "Isdet",
		"usia": "28",
		"pekerjaan": "digital marketing",
	}

	value, exist := person["usia"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value does'nt exist")
	}

	delete(person, "usia")

	value, exist = person["usia"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value has been deleted")
	}
}*/

//combining slice with map
//praktek hal 8
func main() {
	var people = []map[string]string{
		{"nama": "Piya", "umur": "21"},
		{"nama": "Isti", "umur": "28"},
		{"nama": "Weli", "umur": "50"},
	}

	for i, people := range people{
		fmt.Printf("Index: %d, nama: %s, umur: %s\n", i, people["nama"], people["umur"])
	}
}