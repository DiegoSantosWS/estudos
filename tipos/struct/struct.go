package main

import (
	"fmt"
)

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

//Metodo: função com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	fmt.Println("===========================================Exemplo 1========================================")
	var produto1 produto
	produto1 = produto{
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}

	fmt.Println(produto1, produto1.precoComDesconto())
	fmt.Println("===========================================Exemplo 2========================================")
	produto2 := produto{"Notebook", 1989.90, 0.10}
	fmt.Println(produto2, produto2.precoComDesconto())
}
