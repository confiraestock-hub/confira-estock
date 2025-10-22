package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/confiraestock-hub/confira-estock/internal/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var MovimentacaoCollection *mongo.Collection

func CriarMovimentacao(w http.ResponseWriter, r *http.Request) {
	var mov models.Movimentacao
	if err := json.NewDecoder(r.Body).Decode(&mov); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mov.ID = primitive.NewObjectID().Hex()
	mov.DataHora = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := MovimentacaoCollection.InsertOne(ctx, mov)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if insertedID, ok := res.InsertedID.(primitive.ObjectID); ok {
		mov.ID = insertedID.Hex()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mov)
}

func ListarMovimentacoes(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := MovimentacaoCollection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var movimentacoes []models.Movimentacao
	if err = cursor.All(ctx, &movimentacoes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movimentacoes)
}

func AtualizarMovimentacao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]

	var mov models.Movimentacao
	if err := json.NewDecoder(r.Body).Decode(&mov); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": idParam}
	update := bson.M{"$set": mov}
	_, err := MovimentacaoCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mov)
}

func DeletarMovimentacao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := MovimentacaoCollection.DeleteOne(ctx, bson.M{"_id": idParam})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
