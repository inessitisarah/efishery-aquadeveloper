package main

import "fmt"

func rectriangle() {
	fmt.Print("Masukkan angka genap atau ganjil: ")

	// var then variable name then variable type
	var angka int
	fmt.Scan(&angka)
	// Taking input from user

	//fmt.Println(angka)
	if angka%2 == 0 { //angka genap
		fmt.Println("Ini yang Genap:")
		for i := 1; i <= angka; i++ {
			for j := 1; j <= angka; j++ {
				fmt.Print("*")
			}
			fmt.Print("\n")

		}
	} else { //ganjil
		fmt.Println("Ini yang Ganjil:")
		for i := 1; i <= angka; i++ {
			for j := 1; j <= i; j++ {
				fmt.Print("*")
			}
			fmt.Print("\n")

		}
	}
}

func main() {
	rectriangle()
}
