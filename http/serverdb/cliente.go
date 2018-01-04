package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Clientes :)
type Clientes struct {
	ID         int    `json:"id"`
	Nome       string `json:"nome"`
	Cpf        string `json:"cpf"`
	Cnpj       string `json:"cnpj"`
	Nascimento string `json:"nascimento"`
	Sexo       string `json:"sexo"`
	Status     string `json:"status"`
	Tipo       string `json:"tipo"`
	Login      string `json:"login"`
}

// ClientesHandler analisa o resquest e delega para função adequada
func ClientesHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/clientes/")
	id, _ := strconv.Atoi(sid)
	switch {
	case r.Method == "GET" && id > 0:
		clientePorID(w, r, id)
	case r.Method == "GET":
		clientesTodos(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Desculpa... :(")
	}
}

func clientePorID(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("mysql", "user:pass@tcp(serverip:3306)/namebanco")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var cli Clientes
	db.QueryRow("SELECT id, nome, cpf, cnpj, nascimento, sexo, status, tipo, login FROM clientes WHERE id = ?", id).Scan(&cli.ID, &cli.Nome, &cli.Cpf, &cli.Cnpj, &cli.Nascimento, &cli.Sexo, &cli.Status, &cli.Tipo, &cli.Login)

	json, _ := json.Marshal(cli)
	w.Header().Set("Conteont-Type", "application/json")
	fmt.Fprint(w, string(json))
}

func clientesTodos(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "user:pass@tcp(serverip:3306)/namebanco")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("select  id, nome, cpf, cnpj, nascimento, sexo, status, tipo, login  from clientes")
	defer rows.Close()

	var clientes []Clientes
	for rows.Next() {
		var cliente Clientes
		rows.Scan(&cliente.ID, &cliente.Nome, &cliente.Cpf, &cliente.Cnpj,
			&cliente.Nascimento,
			&cliente.Sexo, &cliente.Status, &cliente.Tipo, &cliente.Login)
		clientes = append(clientes, cliente)
	}

	json, _ := json.Marshal(clientes)

	w.Header().Set("Conteont-Type", "application/json")
	fmt.Fprint(w, string(json))
}
