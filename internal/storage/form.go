package storage

import "fmt"

var lastId = uint(0)

type Form struct {
	id   uint
	name string
}

func NewForm(idForm string) (*Form, error) {
	u := Form{}
	if err := u.SetName(idForm); err != nil {
		return nil, err
	}
	lastId++
	u.id = lastId
	return &u, nil
}

func (f *Form) SetName(idForm string) error {
	if len(idForm) == 0 || len(idForm) > 10 {
		return fmt.Errorf("bad name <%v>", idForm)
	}
	f.name = idForm
	return nil
}

func (f Form) String() string {
	return fmt.Sprintf("%d: %s", f.id, f.name)
}
func (f Form) GetName() string {
	return f.name
}

func (f Form) GetId() uint {
	return f.id
}
