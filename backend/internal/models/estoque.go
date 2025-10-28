package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Estoque struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nome          string             `bson:"nome" json:"nome"`
	Tipo          string             `bson:"tipo" json:"tipo"`
	Localizacao   string             `bson:"localizacao" json:"localizacao"`
	Produtos      []ItemEstoque      `bson:"produtos" json:"produtos"`
	ResponsavelID string             `bson:"responsavel_id" json:"responsavelId"`
	DataUltimaMov string             `bson:"data_ultima_mov" json:"dataUltimaMov"`
}
type ItemEstoque struct {
	ProdutoID   string
	ProdutoNome string
	Quantidade  int
	Lote        string // Lote do produto (se aplicável)
	Validade    string // Data de validade (se aplicável)
}
