package main

import "fmt"

func main3() {
	var alas, atas, tinggi float64

	fmt.Printf("Program ini akan menghitung luas dari trapesium. \n")
	fmt.Print("Masukkan panjang alas : \n")
	fmt.Scan(&alas)

	fmt.Print("Masukkan panjang sisi atas : \n")
	fmt.Scan(&atas)

	fmt.Print("Masukkan tinggi trapesium: \n")
	fmt.Scan(&tinggi)

	luas := 0.5 * (alas + atas) * tinggi

	fmt.Printf("Luas trapesium adalah %g\n", luas)

}
