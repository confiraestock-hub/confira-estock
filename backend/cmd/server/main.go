package main

import (
	"log"
	"net/http"
	"os"
	"strings"

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
	// Lê e remove espaços e quebras de linha (\n, \r) do valor da variável
	mongoURI := strings.TrimSpace(os.Getenv("MONGODB_URI"))
	if mongoURI == "" {
		log.Fatal("A variável MONGODB_URI não foi definida")
	}

	// Loga com aspas para identificar espaços invisíveis
	log.Printf("URI recebida: %q\n", mongoURI)

	log.Println("Conectando ao MongoDB em", mongoURI)
	database.Connect(mongoURI)
	log.Println("Conexão ao MongoDB feita com sucesso!")

	r := mux.NewRouter()
	r.HandleFunc("/health", healthHandler).Methods("GET")

	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
