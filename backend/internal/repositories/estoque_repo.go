package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/confiraestock-hub/confira-estock/internal/database"
	"github.com/confiraestock-hub/confira-estock/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var estoqueCollection *mongo.Collection = database.GetCollection("confiraestock", "estoques")

// aqui estão as funções de CRUD para o estoque
func CriarEstoque(e models.Estoque) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := estoqueCollection.InsertOne(ctx, e)
	return err
}

// aqui retorna todos os estoques do banco de dados
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

// aqui busc Por ID e retorna um estoque com base no ID
func BuscarEstoquePorID(id string) (*models.Estoque, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var e models.Estoque
	err := estoqueCollection.FindOne(ctx, bson.M{"id": id}).Decode(&e)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("estoque não encontrado")
	}
	return &e, err
}

// aqui atualiza estoque e modifica um estoque existente com base no ID
func AtualizarEstoque(id string, novo models.Estoque) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filtro := bson.M{"id": id}
	update := bson.M{
		"$set": bson.M{
			"nome":          novo.Nome,
			"tipo":          novo.Tipo,
			"localizacao":   novo.Localizacao,
			"produtos":      novo.Produtos,
			"responsavelid": novo.ResponsavelID,
			"dataultimamov": novo.DataUltimaMov,
		},
	}

	result, err := estoqueCollection.UpdateOne(ctx, filtro, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("nenhum estoque encontrado para atualização")
	}

	return nil
}

// aqui deleta o estoque do banco de dados com base no ID
func DeletarEstoque(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := estoqueCollection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("nenhum estoque encontrado para exclusão")
	}

	return nil
}
