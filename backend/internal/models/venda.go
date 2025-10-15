package models

import "time"

type Venda struct {
	ID             string
	CarrinhoID     string
	CarrinhoNome   string
	UsuarioID      string // Quem realizou a venda
	Produtos       []ItemVenda
	ValorTotal     float64
	FormaPagamento string // Dinheiro, cartão, pix, etc.
	DataHora       time.Time
	Observacao     string
	Status         string  // "finalizada", "cancelada", etc.
	SaldoInicial   float64 // Saldo em dinheiro no início do dia
	SaldoFinal     float64 // Saldo em dinheiro ao final do dia
}

type ItemVenda struct {
	ProdutoID     string
	ProdutoNome   string
	Quantidade    int
	PrecoUnitario float64
	PrecoTotal    float64
}
