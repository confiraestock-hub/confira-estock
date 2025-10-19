package repositories

import (
	"context"
	"time"

	"github.com/NordicManX/Confira-estock/backend/internal/database"
	"github.com/NordicManX/Confira-estock/backend/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var estoqueCollection *mongo.Collection = database.GetCollection("confiraestock", "estoques")

func CriarEstoque(e models.Estoque) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := estoqueCollection.InsertOne(ctx, e)
	return err
}

func ListarEstoques() ([]models.Estoque, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := estoqueCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var estoques []models.Estoque
	if err := cursor.All(ctx, &estoques); err != nil {
		return nil, err
	}
	return estoques, nil
}

func BuscarEstoquePorID(id string) (*models.Estoque, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var e models.Estoque
	err := estoqueCollection.FindOne(ctx, bson.M{"id": id}).Decode(&e)
	return &e, err
}

func AtualizarEstoque(id string, novo models.Estoque) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filtro := bson.M{"id": id}
	update := bson.M{"$set": novo}
	_, err := estoqueCollection.UpdateOne(ctx, filtro, update)
	return err
}

func DeletarEstoque(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := estoqueCollection.DeleteOne(ctx, bson.M{"id": id})
	return err
}
