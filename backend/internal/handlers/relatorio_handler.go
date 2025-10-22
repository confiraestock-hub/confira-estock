package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Relatorio de vendas totais por dia (exemplo simples)
func RelatorioVendasDiarias(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	pipeline := mongo.Pipeline{
		bson.D{{"$group", bson.D{
			{"_id", bson.D{{"$dateToString", bson.D{{"format", "%Y-%m-%d"}, {"date", "$dataHora"}}}}},
			{"totalVendas", bson.D{{"$sum", "$valorTotal"}}},
			{"quantidadeVendas", bson.D{{"$sum", 1}}},
		}}},
		bson.D{{"$sort", bson.D{{"_id", 1}}}},
	}

	cursor, err := VendaCollection.Aggregate(ctx, pipeline)
	if err != nil {
		http.Error(w, "Erro na consulta de vendas: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var resultados []bson.M
	if err = cursor.All(ctx, &resultados); err != nil {
		http.Error(w, "Erro ao ler resultados: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultados)
}
