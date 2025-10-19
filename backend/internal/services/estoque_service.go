package services

import (
	"errors"

	"github.com/NordicManX/Confira-estock/internal/models"
	"github.com/NordicManX/Confira-estock/internal/repositories"
)

func TransferirProduto(origemID, destinoID, produtoID string, quantidade int) error {
	origem, err := repositories.BuscarEstoquePorID(origemID)
	if err != nil {
		return err
	}

	destino, err := repositories.BuscarEstoquePorID(destinoID)
	if err != nil {
		return err
	}
	//aqui busca o produto no estoque de origem
	var produto *models.ItemEstoque
	for i, item := range origem.Produtos {
		if item.ProdutoID == produtoID {
			produto = &origem.Produtos[i]
			break
		}
	}

	if produto == nil || produto.Quantidade < quantidade {
		return errors.New("estoque insulficiente para transferência")
	}

	//Aqui atualiza a origem e destino
	produto.Quantidade -= quantidade
	destino.Produtos = append(destino.Produtos, models.ItemEstoque{
		ProdutoID:   produtoID,
		ProdutoNome: produto.ProdutoNome,
		Quantidade:  quantidade,
	})

	return nil
}

func CalcularNivelEstoque(quantidade, estoqueMaximo int) string {
	percentual := float64(quantidade) / float64(estoqueMaximo) * 100
	switch {
	case percentual > 50:
		return "verde"
	case percentual > 20:
		return "amarelo"
	default:
		return "vermelho"
	}
}
