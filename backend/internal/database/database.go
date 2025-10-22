package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// Connect conecta ao MongoDB e retorna o client e o erro (se houver)
func Connect(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao MongoDB: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("não foi possível pingar o MongoDB: %w", err)
	}

	// Armazena o client globalmente se necessário
	Client = client
	return client, nil
}

// GetCollection retorna uma coleção do MongoDB
func GetCollection(databaseName, collectionName string) *mongo.Collection {
	if Client == nil {
		panic("MongoDB client não inicializado. Chame Connect primeiro.")
	}
	return Client.Database(databaseName).Collection(collectionName)
}
