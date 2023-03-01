package main

import (
	"fmt"
	"strings"
)

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""
	plainText := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	cipherText := "copqrdxyzefabgjuhiklmnstvwCOPQRDXYZEFABGJUHIKLMNSTVW"
	for _, alphabet := range s.name {
		text := strings.Index(plainText, string(alphabet))
		nameEncode += string(cipherText[text])
	}
	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""
	plainText := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	cipherText := "copqrdxyzefabgjuhiklmnstvwCOPQRDXYZEFABGJUHIKLMNSTVW"
	for _, alphabet := range s.name {
		text := strings.Index(cipherText, string(alphabet))
		nameDecode += string(plainText[text])
	}
	return nameDecode
}

func main() {
	var menu int
	var a student = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Println("\nEncode of student’s name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Println("\nDecode of student’s name " + a.name + " is : " + c.Decode())
	} else {
		fmt.Println("Masukkan input dengan benar!")
	}
}
