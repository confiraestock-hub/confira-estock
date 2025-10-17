package repositories

import (
	"errors"

	"github.com/NordicManX/Confira-estock/models"
)

var usuarios = []models.Usuario{}

func CriarUsuario(u models.Usuario) (models.Usuario, error) {
	usuarios = append(usuarios, u)
	return u, nil
}

func BuscarUsuarioPorEmail(email string) (*models.Usuario, error) {
	for _, u := range usuarios {
		if u.Email == email {
			return &u, nil
		}
	}
	return nil, errors.New("usuário não encontrado")
}
func ListarUsuarios() ([]models.Usuario, error) {
	return usuarios, nil
}
