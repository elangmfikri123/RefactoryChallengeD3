package main

import "fmt"

//palindrom
func main() {
	var angka, remainder, temp int
	var reverse int = 0

	fmt.Print("Masukkan Bilangan Positif : ")
	fmt.Scan(&angka)

	temp = angka

	for {
		remainder = angka % 10
		reverse = reverse*10 + remainder
		angka /= 10

		if angka == 0 {
			break
		}
	}

	if temp == reverse {
		fmt.Printf("%d adalah Palindrom", temp)
	} else {
		fmt.Printf("%d Bukan Palindrom", temp)
	}

}
