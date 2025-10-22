package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// ConnectWithContext conecta ao MongoDB Atlas usando contexto externo
func ConnectWithContext(ctx context.Context, uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao MongoDB: %w", err)
	}
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
