package main

import "fmt"

func main() {
	var angka int
	fmt.Printf("Program ini akan menginformasikan apakah suatu angka merupakan bilangan prima\n")
	fmt.Print("Masukkan angka: ")
	_, err := fmt.Scanf("%d", &angka)

	if err != nil {
		fmt.Println("Input tidak valid")
		return
	}
	
	if angka < 2 {
		fmt.Printf("Angka %d bukanlah bilangan prima.", angka)
		return
	}

	for i := 2; i <= angka/2; i++ {
		if angka%i == 0 {
			fmt.Printf("%d bukan bilangan prima\n", angka)
			return
		}
	}
	fmt.Printf("%d bilangan prima\n", angka)
}
