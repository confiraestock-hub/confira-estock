package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Produto struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nome          string             `bson:"nome" json:"nome"`
	Descricao     string             `bson:"descricao" json:"descricao"`
	CategoriaID   string             `bson:"categoria_id" json:"categoriaId"`
	CategoriaNome string             `bson:"categoria_nome" json:"categoriaNome"`
	CodigoBarras  string             `bson:"codigo_barras" json:"codigoBarras"`
	Unidade       string             `bson:"unidade" json:"unidade"`
	Marca         string             `bson:"marca" json:"marca"`
	Fornecedor    string             `bson:"fornecedor" json:"fornecedor"`
	PrecoCusto    float64            `bson:"preco_custo" json:"precoCusto"`
	PrecoVenda    float64            `bson:"preco_venda" json:"precoVenda"`
	EstoqueMinimo int                `bson:"estoque_minimo" json:"estoqueMinimo"`
	EstoqueMaximo int                `bson:"estoque_maximo" json:"estoqueMaximo"`
	ImagemURL     string             `bson:"imagem_url" json:"imagemUrl"`
	Ativo         bool               `bson:"ativo" json:"ativo"`
}
