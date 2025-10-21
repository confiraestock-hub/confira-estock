package services

import (
	"errors"

	"github.com/confiraestock-hub/confira-estock/internal/models"
	"github.com/confiraestock-hub/confira-estock/internal/repositories"
)

func RegistrarVenda(v models.Venda) error {
	estoque, err := repositories.BuscarEstoquePorID(v.CarrinhoID)
	if err != nil {
		return err
	}

	//aqui verifica se tem estoque suficiente para cada item no carrinho
	for _, itemVenda := range v.Produtos {
		encontrado := false
		for i, itemEstoque := range estoque.Produtos {
			if itemEstoque.ProdutoID == itemVenda.ProdutoID {
				encontrado = true
				if itemEstoque.Quantidade < itemVenda.Quantidade {
					return errors.New("quantidade insulficiente em estoque")
				}
				estoque.Produtos[i].Quantidade -= itemVenda.Quantidade
			}
		}
		if !encontrado {
			return errors.New("produto não encontrado no estoque")
		}
	}

	//aqui salva a venda no repositorio de vendas
	return repositories.RegistrarVenda(v)
}
