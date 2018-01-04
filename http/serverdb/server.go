package main

import (
	"log"
	"net/http"
)

func main() {
	// aqui eu passo as rodas que eu quero
	http.HandleFunc("/clientes/", ClientesHandler) // pegando dados /usuarios/ lista todos /usuarios/id pega so id
	log.Println("Executando...")                   // falando que server está executando
	log.Fatal(http.ListenAndServe(":3000", nil))   // passo qual a porta de acesso
}
