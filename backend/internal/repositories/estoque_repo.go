package repositories

import (
	"errors"

	"github.com/NordicManX/Confira-estock/internal/models"
)

var estoques = []models.Estoque{}

// aqui cria o estoque
func CriarEstoque(e models.Estoque) error {
	estoques = append(estoques, e)
	return nil
}

// aqui lista os estoques
func ListarEstoques() ([]models.Estoque, error) {
	return estoques, nil
}

// aqui busca o estoque pelo id
func BuscarEstoquePorID(id string) (*models.Estoque, error) {
	for _, e := range estoques {
		if e.ID == id {
			return &e, nil
		}
	}
	return nil, errors.New("estoque não encontrado")
}

// aqui atualiza o estoque pelo id
func AtualizarEstoque(id string, novo models.Estoque) error {
	for i, e := range estoques {
		if e.ID == id {
			estoques[i] = novo
			return nil
		}
	}
	return errors.New("estoque não encontrado")
}
