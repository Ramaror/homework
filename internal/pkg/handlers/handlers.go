package handlers

//
//import (
//	"fmt"
//	"github.com/pkg/errors"
//	"homework/internal/commander"
//	storage2 "homework/internal/pkg/storage"
//	"log"
//	"strconv"
//	"strings"
//)
//
//const (
//	helpCmd   = "help"
//	listCmd   = "list"
//	addCmd    = "add"
//	updateCmd = "update"
//	deleteCmd = "delete"
//)
//
//var BadArgument = errors.New("bad argument")
//
//func listFunc(s string) string {
//	data := storage2.List()
//	res := make([]string, 0, len(data))
//	for _, v := range data {
//		res = append(res, v.String())
//	}
//	return strings.Join(res, "\n")
//}
//
//func helpFunc(s string) string {
//	return "/help - list commands\n" +
//		"/list - list data\n" +
//		"/add <id> - add new id with form\n" +
//		"/delete <id> - delete id with form\n" +
//		"/update <id> <new id> - updated"
//}
//
//func addFunc(data string) string {
//	log.Printf("add command param: <data>")
//	params := strings.Split(data, " ")
//	if len(params) != 1 {
//		return errors.Wrapf(BadArgument, "%d items: <%v>", len(params), params).Error()
//	}
//	u, err := storage2.NewForm(params[0])
//	if err != nil {
//		return err.Error()
//	}
//	err = storage2.Add(u)
//	if err != nil {
//		return err.Error()
//	}
//	return fmt.Sprintf("form %v added", u)
//}
//func deleteFunc(data string) string {
//	log.Printf("upp command param: <data>")
//	params := strings.Split(data, " ")
//	if len(params) != 1 {
//		return errors.Wrapf(BadArgument, "%d items: <%v>", len(params), params).Error()
//	}
//	u, err := strconv.Atoi(params[0])
//	if err != nil {
//		return err.Error()
//	}
//	err = storage2.Delete(uint(u))
//	if err != nil {
//		return err.Error()
//	}
//	return fmt.Sprintf("form %v deleted", u)
//}
//
//func updateFunc(data string) string {
//	log.Printf("add command param: <data>")
//	params := strings.Split(data, " ")
//	fmt.Println("params", params)
//	if len(params) != 2 {
//		return errors.Wrapf(BadArgument, "%d items: <%v>", len(params), params).Error()
//	}
//	id, err := strconv.Atoi(params[0])
//	if err != nil {
//		return err.Error()
//	}
//	form, err := storage2.Get(uint(id))
//	if err != nil {
//		return err.Error()
//	}
//	err = form.SetName(params[1])
//	if err != nil {
//		return err.Error()
//	}
//	err = storage2.Update(form)
//	fmt.Println("update", form)
//	return fmt.Sprintf("form %v updated", form)
//}
//
//func AddHandlers(c *commander.commander) {
//	c.RegisterHandler(helpCmd, helpFunc)
//	c.RegisterHandler(listCmd, listFunc)
//	c.RegisterHandler(addCmd, addFunc)
//	c.RegisterHandler(updateCmd, updateFunc)
//	c.RegisterHandler(deleteCmd, deleteFunc)
//}
