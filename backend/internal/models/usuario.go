package models

type Usuario struct {
	ID          string
	Nome        string
	Email       string
	SenhaHash   string // Senha criptografada
	Perfil      string // "admin" ou "carrinho"
	Ativo       bool
	UltimoLogin string
}
