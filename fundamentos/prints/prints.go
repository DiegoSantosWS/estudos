package main

import (
	"fmt"
)

func main() {
	fmt.Print("Mesma")
	fmt.Print("Linha.")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 13.141516
	xs := fmt.Sprint(x)

	fmt.Println("o valor de x é", xs)
	fmt.Printf("O valor de x é %.2f. ", x)

	a := 1
	b := 1.2323
	c := false
	d := "opa"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
}
