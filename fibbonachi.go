package main

import (
	"fmt"
) //fibbonachi
func fibbonachi(i int) (ret int) {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibbonachi(i-1) + fibbonachi(i-2)
}
func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d ", fibbonachi(i))
	}
}
