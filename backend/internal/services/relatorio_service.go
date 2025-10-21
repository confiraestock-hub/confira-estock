package services

import (
	"time"

	"github.com/confiraestock-hub/confira-estock/internal/models"
	"github.com/confiraestock-hub/confira-estock/internal/repositories"
)

type Relatorio struct {
	Data     time.Time
	Vendas   []models.Venda
	TotalDia float64
	Estoques []models.Estoque
}

func GerarRelatorio(data time.Time) Relatorio {
	vendas, _ := repositories.ListarVendas()
	estoques, _ := repositories.ListarEstoques()

	total := 0.0
	for _, v := range vendas {
		if v.DataHora.Format("2006-01-02") == data.Format("2006-01-02") {
			total += v.ValorTotal
		}
	}

	return Relatorio{
		Data:     data,
		Vendas:   vendas,
		TotalDia: total,
		Estoques: estoques,
	}
}
