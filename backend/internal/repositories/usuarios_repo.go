package repositories

import (
	"context"
	"time"

	"github.com/confiraestock-hub/Confira-estock/backend/internal/database"
	"github.com/confiraestock-hub/Confira-estock/backend/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var usuarioCollection *mongo.Collection = database.GetCollection("confiraestock", "usuarios")

func CriarUsuario(u models.Usuario) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := usuarioCollection.InsertOne(ctx, u)
	return err
}

func ListarUsuarios() ([]models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := usuarioCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var usuarios []models.Usuario
	if err := cursor.All(ctx, &usuarios); err != nil {
		return nil, err
	}
	return usuarios, nil
}

func BuscarUsuarioPorEmail(email string) (*models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var u models.Usuario
	err := usuarioCollection.FindOne(ctx, bson.M{"email": email}).Decode(&u)
	return &u, err
}

func AtualizarUsuario(email string, novo models.Usuario) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filtro := bson.M{"email": email}
	update := bson.M{"$set": novo}
	_, err := usuarioCollection.UpdateOne(ctx, filtro, update)
	return err
}

func DeletarUsuario(email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := usuarioCollection.DeleteOne(ctx, bson.M{"email": email})
	return err
}
