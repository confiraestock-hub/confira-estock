package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/confiraestock-hub/confira-estock/internal/database"
	"github.com/gorilla/mux"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API está online e conectada ao MongoDB"))
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 - Rota não encontrada"))
}

func main() {
	// Lê e limpa a variável de ambiente
	mongoURI := strings.TrimSpace(os.Getenv("MONGODB_URI"))
	if mongoURI == "" {
		log.Fatal("A variável MONGODB_URI não foi definida")
	}

	log.Printf("URI recebida: %q\n", mongoURI)

	// Cria contexto com timeout de 20 segundos para conexão
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Conecta ao MongoDB usando contexto
	client, err := database.ConnectWithContext(ctx, mongoURI)
	if err != nil {
		log.Fatalf("Falha ao conectar ao MongoDB: %v", err)
	}

	// Confirma conexão com Ping
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Não foi possível pingar o MongoDB: %v", err)
	}

	log.Println("Conexão ao MongoDB feita com sucesso!")

	// Configura router
	r := mux.NewRouter()
	r.HandleFunc("/health", healthHandler).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
