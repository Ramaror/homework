package cache

import (
	"homework/internal/pkg/core/user/models"
)

type Interface interface {
	Add(user models.User) error
	Delete(name string) error
	List() []models.User
	Update(user models.User) error
}
