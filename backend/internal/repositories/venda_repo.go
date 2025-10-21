package repositories

import (
	"context"
	"time"

	"github.com/confiraestock-hub/Confira-estock/backend/internal/database"
	"github.com/confiraestock-hub/Confira-estock/backend/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var vendaCollection *mongo.Collection = database.GetCollection("confiraestock", "vendas")

func RegistrarVenda(v models.Venda) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := vendaCollection.InsertOne(ctx, v)
	return err
}

func ListarVendas() ([]models.Venda, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := vendaCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var vendas []models.Venda
	if err := cursor.All(ctx, &vendas); err != nil {
		return nil, err
	}
	return vendas, nil
}

func BuscarVendaPorID(id string) (*models.Venda, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var v models.Venda
	err := vendaCollection.FindOne(ctx, bson.M{"id": id}).Decode(&v)
	return &v, err
}

func DeletarVenda(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := vendaCollection.DeleteOne(ctx, bson.M{"id": id})
	return err
}
