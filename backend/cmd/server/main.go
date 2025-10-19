package main

import (
	"log"
	"net/http"
	"os"

	"github.com/NordicManX/Confira-estock/internal/database"
)

func main() {
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("A variável MONGODB_URI não foi definida")
	}

	log.Println("Conectando ao MongoDB em", mongoURI)
	database.Connect(mongoURI)
	log.Println("Conectado ao MongoDB com sucesso!")

	// Rota básica para testar se a aplicação está online
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API está online e conectada ao MongoDB"))
	})

	log.Println("Servidor rodando na porta 8080")
	// Esta linha mantém a aplicação rodando
	log.Fatal(http.ListenAndServe(":8080", nil))
}
