package main

import "fmt"

func main2() {
	var angka int

	fmt.Print("program ini berfungsi untuk menginfokan suatu bilangan kelipatan 7 atau bukan. \n ")
	fmt.Print("Silakan masukkan angka yang ingin di cek: ")
	_, err := fmt.Scanf("%d", &angka)
	//pengecekan kalau ini angka
	if err != nil {
		fmt.Println("Input tidak valid")
		return
	}

	if angka%7 == 0 {
		fmt.Printf("Angka %d adalah kelipatan 7.\n", angka)
	} else {
		fmt.Printf("Angka %d bukan kelipatan 7.\n", angka)
	}
}
