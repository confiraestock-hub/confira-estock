package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Categoria struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nome  string             `bson:"nome" json:"nome"`
	Ativo bool               `bson:"ativo" json:"ativo"`
}
