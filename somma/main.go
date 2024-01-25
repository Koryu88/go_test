package main

import "fmt"

func main() {
	fmt.Println("inserisci primo numero")
	var a, b int
	fmt.Scanln(&a)
	fmt.Println("inserisci secondo numero")
	fmt.Scanln(&b)

	var c int
	c = a + b

	fmt.Printf("La somma è %d", c)

}
