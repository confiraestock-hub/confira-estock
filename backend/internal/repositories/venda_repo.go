package repositories

import (
	"github.com/NordicManX/Confira-estock/internal/models"
)

var vendas = []models.Venda{}

func RegistrarVenda(v models.Venda) error {
	vendas = append(vendas, v)
	return nil
}

func ListarVendas() ([]models.Venda, error) {
	return vendas, nil
}
