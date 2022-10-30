package local

import (
	"github.com/pkg/errors"
	cachePkg "homework/internal/pkg/core/user/cache"
	"homework/internal/pkg/core/user/models"
	"sync"
)

var (
	ErrUserNotExists = errors.New("user does not exist")
	ErrUserExists    = errors.New("user exists")
)

func New() cachePkg.Interface {
	return &cache{
		mu:   sync.RWMutex{},
		data: map[string]models.User{},
	}
}

type cache struct {
	mu   sync.RWMutex
	data map[string]models.User
}

func (c *cache) List() []models.User {
	c.mu.RLock()
	defer c.mu.RUnlock()

	result := make([]models.User, 0, len(c.data))
	for _, value := range c.data {
		result = append(result, value)
	}
	return result
}
func (c *cache) Add(user models.User) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.data[user.Name]; ok {
		return errors.Wrapf(ErrUserExists, "user-name: [%s]", user.Name)
	}
	c.data[user.Name] = user
	return nil
}

func (c *cache) Update(user models.User) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.data[user.Name]; !ok {
		return errors.Wrapf(ErrUserNotExists, "user-name: [%s]", user.Name)
	}
	c.data[user.Name] = user
	return nil
}

func (c *cache) Delete(name string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.data[name]; ok {
		delete(c.data, name)
		return nil
	}
	return errors.Wrapf(ErrUserNotExists, "user-name: [%s]", name)
}
