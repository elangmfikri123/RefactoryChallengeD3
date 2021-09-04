package main

import "fmt"

func main() {
	var f = do
	f()
	f = func() { fmt.Println("Berhenti lakukan!") }
}
func do() {
	fmt.Println()
	fmt.Println("Function Expression")
	fmt.Println()
	fmt.Println("Lakukan!")
}
