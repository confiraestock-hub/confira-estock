package main

import (
	"log"
	"net/http"
	"os"

	"github.com/confiraestock-hub/confira-estock/internal/database"
	"github.com/gorilla/mux"
	// importe seus handlers aqui, por exemplo:
	// "github.com/NordicManX/Confira-estock/backend/internal/handlers"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API está online e conectada ao MongoDB"))
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 - Rota não encontrada"))
}

func main() {
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("A variável MONGODB_URI não foi definida")
	}

	// Imprime a URI com aspas para revelar espaços ou caracteres invisíveis
	log.Printf("URI recebida: %q\n", mongoURI)

	log.Println("Conectando ao MongoDB em", mongoURI)
	database.Connect(mongoURI)
	log.Println("Conexão ao MongoDB feita com sucesso!")

	r := mux.NewRouter()
	r.HandleFunc("/health", healthHandler).Methods("GET")

	// Exemplos de rotas futuras:
	// r.HandleFunc("/estoques", handlers.ListarEstoques).Methods("GET")
	// r.HandleFunc("/produtos", handlers.ListarProdutos).Methods("GET")
	// r.HandleFunc("/vendas", handlers.ListarVendas).Methods("GET")

	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
