package main

import (
	"fmt"
	"os"
	"strconv"
)

type Participant struct {
	Nomor	int
	Nama	string
	Alamat	string
	Hobi	string
}

func main() {
	var participants = []Participant {
		{
			Nomor:  1,
			Nama:   "Alfian",
			Alamat: "Metro",
			Hobi:   "Bermain basket",
		},
		{
			Nomor:  2,
			Nama:   "Pia",
			Alamat: "Lampung Timur",
			Hobi:   "Membaca novel",
		},
		{
			Nomor:  3,
			Nama:   "Teza",
			Alamat: "Lampung Timur",
			Hobi:   "Menonton film",
		},
	}
	var argsRaw = os.Args
	var args = argsRaw[1]
	nomor, err := strconv.Atoi(args)
	_ = err

	getParticipant(participants, nomor)
}

func getParticipant(p []Participant, nomor int) {
	if nomor <= 3 {
		fmt.Println("=========================")
		fmt.Println("Anggota Keluarga Gen Suti")
		fmt.Println("Nomor:", p[nomor-1].Nomor)
		fmt.Println("Nama:", p[nomor-1].Nama)
		fmt.Println("Hobi:", p[nomor-1].Hobi)
		fmt.Println("=========================")
	} else {
		fmt.Println("Tidak Terdaftar!")
	}
}