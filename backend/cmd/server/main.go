package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/confiraestock-hub/confira-estock/internal/database"
	"github.com/confiraestock-hub/confira-estock/internal/handlers"
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
	mongoURI := strings.TrimSpace(os.Getenv("MONGODB_URI"))
	if mongoURI == "" {
		log.Fatal("A variável MONGODB_URI não foi definida")
	}

	log.Printf("URI recebida: %q\n", mongoURI)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := database.ConnectWithContext(ctx, mongoURI)
	if err != nil {
		log.Fatalf("Falha ao conectar ao MongoDB: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Não foi possível pingar o MongoDB: %v", err)
	}

	log.Println("Conexão ao MongoDB feita com sucesso!")

	// Inicializando as coleções
	handlers.CategoriaCollection = client.Database("confiraestock").Collection("categorias")
	handlers.ProdutoCollection = client.Database("confiraestock").Collection("produtos")
	handlers.EstoqueCollection = client.Database("confiraestock").Collection("estoques")
	handlers.MovimentacaoCollection = client.Database("confiraestock").Collection("movimentacoes")
	handlers.UsuarioCollection = client.Database("confiraestock").Collection("usuarios")
	handlers.VendaCollection = client.Database("confiraestock").Collection("vendas")

	r := mux.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	r.HandleFunc("/health", healthHandler).Methods("GET")

	// Rotas Categorias
	r.HandleFunc("/categorias", handlers.CriarCategoria).Methods("POST")
	r.HandleFunc("/categorias", handlers.ListarCategorias).Methods("GET")
	r.HandleFunc("/categorias/{id}", handlers.AtualizarCategoria).Methods("PUT")
	r.HandleFunc("/categorias/{id}", handlers.DeletarCategoria).Methods("DELETE")

	// Rotas Produtos
	r.HandleFunc("/produtos", handlers.CriarProduto).Methods("POST")
	r.HandleFunc("/produtos", handlers.ListarProdutos).Methods("GET")
	r.HandleFunc("/produtos/{id}", handlers.AtualizarProduto).Methods("PUT")
	r.HandleFunc("/produtos/{id}", handlers.DeletarProduto).Methods("DELETE")

	// Rotas Estoques
	r.HandleFunc("/estoques", handlers.CriarEstoque).Methods("POST")
	r.HandleFunc("/estoques", handlers.ListarEstoques).Methods("GET")
	r.HandleFunc("/estoques/{id}", handlers.AtualizarEstoque).Methods("PUT")
	r.HandleFunc("/estoques/{id}", handlers.DeletarEstoque).Methods("DELETE")

	// Rotas Movimentações
	r.HandleFunc("/movimentacoes", handlers.CriarMovimentacao).Methods("POST")
	r.HandleFunc("/movimentacoes", handlers.ListarMovimentacoes).Methods("GET")
	r.HandleFunc("/movimentacoes/{id}", handlers.AtualizarMovimentacao).Methods("PUT")
	r.HandleFunc("/movimentacoes/{id}", handlers.DeletarMovimentacao).Methods("DELETE")

	// Rotas Usuários
	r.HandleFunc("/usuarios", handlers.CriarUsuario).Methods("POST")
	r.HandleFunc("/usuarios", handlers.ListarUsuarios).Methods("GET")
	r.HandleFunc("/usuarios/{id}", handlers.AtualizarUsuario).Methods("PUT")
	r.HandleFunc("/usuarios/{id}", handlers.DeletarUsuario).Methods("DELETE")

	// Rotas Vendas
	r.HandleFunc("/vendas", handlers.CriarVenda).Methods("POST")
	r.HandleFunc("/vendas", handlers.ListarVendas).Methods("GET")
	r.HandleFunc("/vendas/{id}", handlers.AtualizarVenda).Methods("PUT")
	r.HandleFunc("/vendas/{id}", handlers.DeletarVenda).Methods("DELETE")

	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
