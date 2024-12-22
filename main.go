package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func conectarDB() *sql.DB {
	// user-database-password-host-ssl
	conexao := "user=postgres dbname=alura_loja password=0512 host=localhost sslmode=disable"

	db, err := sql.Open("postgres", conexao)

	if err != nil {
		// Encerra a aplicação
		panic(err.Error())
	}

	return db
}

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	// Abre conexão com banco
	db := conectarDB()
	// defer -> Ele executa depois de ter executado tudo da aplicação
	// Encerra conexão com banco
	defer db.Close()

	http.HandleFunc("/", Index)

	http.ListenAndServe(":8000", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {

	// Abre conexão com banco
	db := conectarDB()

	selectProds, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProds.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProds.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Id = id
		p.Quantidade = quantidade
		p.Preco = preco

		produtos = append(produtos, p)
	}

	temp.ExecuteTemplate(w, "Index", produtos)

	// defer -> Ele executa depois de ter executado tudo da aplicação
	// Encerra conexão com banco
	defer db.Close()
}
