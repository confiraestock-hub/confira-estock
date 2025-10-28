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

var EstoqueCollection *mongo.Collection // Inicialize no main

// Criar Estoque
func CriarEstoque(w http.ResponseWriter, r *http.Request) {
	var est models.Estoque
	if err := json.NewDecoder(r.Body).Decode(&est); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := EstoqueCollection.InsertOne(ctx, est)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	est.ID = res.InsertedID.(primitive.ObjectID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(est)
}

// Listar todos os Estoques
func ListarEstoques(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := EstoqueCollection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var estoques []models.Estoque
	if err = cursor.All(ctx, &estoques); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(estoques)
}

// Atualizar Estoque pelo ID
func AtualizarEstoque(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	objId, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	var est models.Estoque
	if err := json.NewDecoder(r.Body).Decode(&est); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"_id": objId}
	update := bson.M{"$set": est}
	result, err := EstoqueCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if result.ModifiedCount == 0 {
		http.Error(w, "Estoque não encontrado ou não alterado", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(est)
}

// Deletar Estoque pelo ID
func DeletarEstoque(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := EstoqueCollection.DeleteOne(ctx, bson.M{"_id": idParam})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
