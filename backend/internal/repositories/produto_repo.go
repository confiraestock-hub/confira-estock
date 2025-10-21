package repositories

import (
	"context"
	"time"

	"github.com/confiraestock-hub/confira-estock/internal/database"
	"github.com/confiraestock-hub/confira-estock/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var produtoCollection *mongo.Collection = database.GetCollection("confiraestock", "produtos")

func CriarProduto(prod models.Produto) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := produtoCollection.InsertOne(ctx, prod)
	return err
}

func ListarProdutos() ([]models.Produto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := produtoCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var produtos []models.Produto
	if err := cursor.All(ctx, &produtos); err != nil {
		return nil, err
	}
	return produtos, nil
}

func BuscarProdutoPorID(id string) (*models.Produto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var prod models.Produto
	err := produtoCollection.FindOne(ctx, bson.M{"id": id}).Decode(&prod)
	return &prod, err
}

func AtualizarProduto(id string, novo models.Produto) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filtro := bson.M{"id": id}
	update := bson.M{"$set": novo}
	_, err := produtoCollection.UpdateOne(ctx, filtro, update)
	return err
}

func DeletarProduto(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := produtoCollection.DeleteOne(ctx, bson.M{"id": id})
	return err
}
