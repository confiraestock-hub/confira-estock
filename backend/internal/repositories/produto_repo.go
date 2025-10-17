package repositories

import (
	"errors"

	"github.com/NordicManX/Confira-estock/internal/models"
)

var produtos = []models.Produto{}

// aqui cria o produto
func CriarProduto(p models.Produto) error {
	produtos = append(produtos, p)
	return nil
}

// aqui lista os produtos
func ListarProdutos() ([]models.Produto, error) {
	return produtos, nil
}

// aqui busca o produto pelo id
func BuscarProdutoPorID(id string) (models.Produto, error) {
	for _, p := range produtos {
		if p.ID == id {
			return p, nil
		}
	}
	return models.Produto{}, errors.New("produto não encontrado")
}

// aqui atualiza o produto pelo id
func AtualizarProduto(id string, p models.Produto) error {
	for i, prod := range produtos {
		if prod.ID == id {
			produtos[i] = p
			return nil
		}
	}
	return errors.New("produto não encontrado")
}

// aqui deleta o produto pelo id
func DeletarProduto(id string) error {
	for i, p := range produtos {
		if p.ID == id {
			produtos = append(produtos[:i], produtos[i+1:]...)
			return nil
		}
	}
	return errors.New("produto não encontrado")
}
