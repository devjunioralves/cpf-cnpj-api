package repositories

import "cpf-cnpj-api/internal/domain/models"

type UserRepository interface {
	Save(user *models.User) error
	FindByUsername(username string) (*models.User, error)
}
