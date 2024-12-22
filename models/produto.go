package models

import "github.com/deyweson/learning-go-api-mvc/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscarProdutos() []Produto {
	// Abre conexão com banco
	db := db.ConectarDB()

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
	// defer -> Ele executa depois de ter executado tudo da aplicação
	// Encerra conexão com banco
	defer db.Close()
	return produtos
}
