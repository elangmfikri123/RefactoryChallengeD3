package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(randomString(10))
}

func randomString(n int) string {
	fmt.Println()
	fmt.Println("Kalimat random sebanyak", n, ":")
	var kata = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = kata[rand.Intn(len(kata))]
	}
	return string(s)
}
