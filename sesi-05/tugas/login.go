package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/howeyc/gopass"
)

func main() {
	var username string
	var username2 string

	//masukkan data baru
	fmt.Println("\nBuat Akun Baru\n")
	fmt.Print("Username : ")
	fmt.Scanln(&username)
	fmt.Print("Password : ")
	pass1, _ := gopass.GetPasswdMasked()
	fmt.Print("-Akun Berhasil Dibuat-\n")
	
fmt.Println(strings.Repeat("_", 50))

	//login
	fmt.Println("\nLogin\n")
	fmt.Print("Username : ")
	fmt.Scanln(&username2)
	fmt.Print("Password : ")
	//password berbentuk bintang
	pass2, _ := gopass.GetPasswdMasked() 

	//pengecekan password
	if valid, err := validPass(username, username2, pass1, pass2); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(valid)
	}

}

func validPass(username string, username2 string, pass1 []byte, pass2 []byte) (string, error) {
	p1 := string(pass1[:])
	p2 := string(pass2[:])
	if username != username2 || p1 != p2 {
		return "", errors.New("Username atau Password Salah")
	}

	return "Login Berhasil", nil
}