package handlers

import (
	"fmt"
	"github.com/pkg/errors"
	"homework/internal/commander"
	"homework/internal/storage"
	"log"
	"strings"
)

const (
	helpCmd   = "help"
	listCmd   = "list"
	addCmd    = "add"
	updateCmd = "update"
	deleteCmd = "delete"
)

var BadArgument = errors.New("bad argument")

func listFunc(s string) string {
	data := storage.List()
	res := make([]string, 0, len(data))
	for _, v := range data {
		res = append(res, v.String())
	}
	return strings.Join(res, "\n")
}

func helpFunc(s string) string {
	return "/help - list commands\n" +
		"/list - list data\n" +
		"/add <id> - add new id with form"
}

func addFunc(data string) string {
	log.Printf("add command param: <data>")
	params := strings.Split(data, " ")
	if len(params) != 1 {
		return errors.Wrapf(BadArgument, "%d items: <%v>", len(params), params).Error()
	}
	u, err := storage.NewForm(params[0])
	if err != nil {
		return err.Error()
	}
	err = storage.Add(u)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("user %v added", u)
}
func deleteFunc(data string) string {
	log.Printf("upp command param: <data>")
	params := strings.Split(data, " ")
	if len(params) != 1 {
		return errors.Wrapf(BadArgument, "%d items: <%v>", len(params), params).Error()
	}
	u, err := storage.NewForm(params[0])
	if err != nil {
		return err.Error()
	}
	err = storage.Delete(uint(0))
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("user %v added", u)
}
func updateFunc(data string) string {
	log.Printf("upp command param: <data>")
	params := strings.Split(data, " ")
	if len(params) != 1 {
		return errors.Wrapf(BadArgument, "%d items: <%v>", len(params), params).Error()
	}
	u, err := storage.NewForm(params[0])
	if err != nil {
		return err.Error()
	}
	err = storage.Update(u)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("user %v added", u)
}

func AddHandlers(c *commander.Commander) {
	c.RegisterHandler(helpCmd, helpFunc)
	c.RegisterHandler(listCmd, listFunc)
	c.RegisterHandler(addCmd, addFunc)
	c.RegisterHandler(updateCmd, updateFunc)
	c.RegisterHandler(deleteCmd, deleteFunc)
}
