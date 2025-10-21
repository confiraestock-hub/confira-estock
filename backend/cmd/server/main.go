package main

import (
    "log"
    "net/http"
    "os"

    "github.com/NordicManX/Confira-estock/backend/internal/database"
)

func main() {
    mongoURI := os.Getenv("MONGODB_URI")
    if mongoURI == "" {
        log.Fatal("A variável MONGODB_URI não foi definida")
    }

    log.Println("Conectando ao MongoDB em", mongoURI)
    database.Connect(mongoURI)
    log.Println("Conexão ao MongoDB feita com sucesso!")

    // Rota de teste para ver se está no ar
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("API está online e conectada ao MongoDB"))
    })

    log.Println("Servidor rodando na porta 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
