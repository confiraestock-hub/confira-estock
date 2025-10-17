package models

import "time"

type Movimentacao struct {
	ID               string
	EstoqueOrigemID  string
	EstoqueDestinoID string
	ProdutoID        string
	Quantidade       int
	Tipo             string // "entrada", "saida", "transferencia"
	UsuarioID        string
	DataHora         time.Time
	Observacao       string
}
