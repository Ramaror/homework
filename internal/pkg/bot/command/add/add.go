package add

import (
	"github.com/pkg/errors"
	commandPkg "homework/internal/pkg/bot/command"
	userPkg "homework/internal/pkg/core/user"
	"homework/internal/pkg/core/user/models"
	"strings"
)

func New(user userPkg.Interface) commandPkg.Interface {
	return &command{
		user: user,
	}
}

type command struct {
	user userPkg.Interface
}

func (c *command) Name() string {
	return "add"
}
func (c *command) Description() string {
	return "create user"
}

func (c *command) Process(args string) string {
	params := strings.Split(args, " ")
	if len(params) != 2 {
		return "invalid args"
	}
	if err := c.user.Create(models.User{
		Name:     params[0],
		Password: params[1],
	}); err != nil {
		if errors.Is(err, userPkg.ErrValidation) {
			return "invalid args"
		}
		return "internal error"
	}
	return "succes"
}
