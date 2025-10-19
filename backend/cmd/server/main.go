package main

import (
	"log"
	"os"
)

func main() {
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("A variável MONGODB_URI não foi definida")
	}

	connectToDB(mongoURI)
	// chama inicialização de rotas HTTP
}

func connectToDB(uri string) {
	// Stub de conexão: para conectar de verdade ao MongoDB,
	// adicione a dependência "go.mongodb.org/mongo-driver/mongo"
	// ao go.mod e substitua este stub pela lógica de conexão.
	log.Printf("Conectando ao MongoDB em %s (stub)", uri)
}
