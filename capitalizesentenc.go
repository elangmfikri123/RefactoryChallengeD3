package main

import (
	"fmt"
	"strings"
)

//capitalizesentenc
func main() {

	kalimat := "elang muhammad Fikhri"
	kapital := strings.ToLower(kalimat)

	fmt.Println("Sebelum:", kalimat)
	fmt.Println("Setelah: ", kapital)
}
