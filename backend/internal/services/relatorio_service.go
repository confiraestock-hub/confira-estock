package services

import (
	"time"

	"github.com/NordicManX/Confira-estock/internal/repositories"
	"github.com/cconfiraestock-hub/Confira-estock/internal/models"
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
