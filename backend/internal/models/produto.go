package models

type Produto struct {
	ID            string  // Identificador único
	Nome          string  // Nome do produto
	Descricao     string  // Descrição detalhada
	CategoriaID   string  // Relacionamento com categoria
	CategoriaNome string  // Nome da categoria (opcional para facilitar consultas)
	CodigoBarras  string  // Código de barras
	Unidade       string  // Unidade de medida (un, kg, l, etc.)
	Marca         string  // Marca do produto
	Fornecedor    string  // Nome do fornecedor
	PrecoCusto    float64 // Preço de custo
	PrecoVenda    float64 // Preço de venda
	EstoqueMinimo int     // Quantidade mínima para alerta
	EstoqueMaximo int     // Quantidade máxima permitida
	ImagemURL     string  // URL da imagem do produto
	Ativo         bool    // Produto ativo/inativo
}
