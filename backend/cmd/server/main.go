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

func main() {
	mongoURI := strings.TrimSpace(os.Getenv("MONGODB_URI"))
	if mongoURI == "" {
		log.Fatal("A variável MONGODB_URI não foi definida")
	}

	log.Printf("URI recebida: %q\n", mongoURI)

	// Cria um contexto com timeout de 20 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Conecta ao MongoDB usando o contexto
	client, err := database.ConnectWithContext(ctx, mongoURI) // função ajustada no database
	if err != nil {
		log.Fatalf("Falha ao conectar ao MongoDB: %v", err)
	}

	// Confirma conexão com Ping usando o mesmo contexto
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Não foi possível pingar o MongoDB: %v", err)
	}

	log.Println("Conexão ao MongoDB feita com sucesso!")

	r := mux.NewRouter()
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API está online e conectada ao MongoDB"))
	}).Methods("GET")

	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
