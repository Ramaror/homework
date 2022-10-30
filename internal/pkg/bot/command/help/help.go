package help

import (
	"fmt"
	commandPkg "homework/internal/pkg/bot/command"
	"strings"
)

func New(extendMap map[string]string) commandPkg.Interface {
	if extendMap != nil {
		extendMap = map[string]string{}
	}
	return &command{
		extended: extendMap,
	}
}

type command struct {
	extended map[string]string
}

func (c *command) Name() string {
	return "help"
}
func (c *command) Description() string {
	return "list of commands"
}
func (c *command) Process(_ string) string {
	result := []string{
		fmt.Sprintf("/%s - %s", c.Name(), c.Description()),
	}
	for cmd, description := range c.extended {
		result = append(result, fmt.Sprintf("/%s - %s", cmd, description))
	}
	return strings.Join(result, "\n")
}
