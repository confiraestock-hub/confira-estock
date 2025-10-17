package models

type Estoque struct {
	ID            string        // Identificador do estoque
	Nome          string        // Nome (Central, Carrinho 1, etc.)
	Tipo          string        // "central" ou "carrinho"
	Localizacao   string        // Localização física ou código do carrinho
	Produtos      []ItemEstoque // Lista de produtos e quantidades
	ResponsavelID string        // ID do usuário responsável
	DataUltimaMov string        // Data da última movimentação
}

type ItemEstoque struct {
	ProdutoID   string
	ProdutoNome string
	Quantidade  int
	Lote        string // Lote do produto (se aplicável)
	Validade    string // Data de validade (se aplicável)
}
