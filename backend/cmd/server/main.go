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

	log.Println("Tentando conectar ao MongoDB...")
	database.Connect(mongoURI)
	log.Println("Tentando conectar ao MongoDB...")
	connectDatabase(mongoURI)
	log.Println("Conectado ao MongoDB com sucesso!")

	// Rota de teste HTTP para conferir se a API está rodando
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API está online e conectada ao MongoDB"))
	})

	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// connectDatabase is a lightweight local stub to avoid the missing external package.
// Replace this with the real MongoDB connection logic when the internal/database package is available.
func connectDatabase(mongoURI string) {
	log.Printf("Mock connect to MongoDB at %s (no-op)", mongoURI)
}
