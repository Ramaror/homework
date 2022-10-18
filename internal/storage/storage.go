package storage

import (
	"github.com/pkg/errors"
	"log"
	"strconv"
)

var data map[uint]*Form

var UserNotExists = errors.New("user does not exist")
var UserExists = errors.New("user exists")

func init() {
	log.Println("init storage")
	data = make(map[uint]*Form)
	u, _ := NewForm("0")
	if err := Add(u); err != nil {
		log.Panic(err)
	}
}

func List() []*Form {
	res := make([]*Form, 0, len(data))
	for _, v := range data {
		res = append(res, v)
	}
	return res
}
func Add(f *Form) error {
	if _, ok := data[f.GetId()]; ok {
		return errors.Wrap(UserExists, strconv.FormatUint(uint64(f.GetId()), 10))
	}
	data[f.GetId()] = f
	return nil
}

func Update(f *Form) error {
	if _, ok := data[f.GetId()]; !ok {
		return errors.Wrap(UserNotExists, strconv.FormatUint(uint64(f.GetId()), 10))
	}
	data[f.GetId()] = f
	return nil
}

func Delete(id uint) error {
	if _, ok := data[id]; ok {
		delete(data, id)
		return nil
	}
	return errors.Wrap(UserNotExists, strconv.FormatUint(uint64(id), 10))
}
