package note

import (
	models "notes-app/data"
)

type Tag interface {
	GetID() uint
	GetName() string

	FindTag(tag string) (Tag, error)
	GetTags() ([]Tag, error)
	Create() error
}

type TagImp struct {
	models.Tag
}

func (t TagImp) GetID() uint {
	return t.ID
}

func (t TagImp) GetName() string {
	return t.Name
}

func (t TagImp) FindTag(tag string) (Tag, error) {
	panic("not implemented") // TODO: Implement
}

func (t TagImp) GetTags() ([]Tag, error) {
	panic("not implemented") // TODO: Implement
}

func (t TagImp) Create() error {
	panic("not implemented") // TODO: Implement
}
