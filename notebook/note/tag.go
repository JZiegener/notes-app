package note

import (
	"errors"
	models "notes-app/data"

	"gorm.io/gorm"
)

/*
Tag used to
*/
type Tag interface {
	GetID() uint
	GetName() string

	GetTags() ([]Tag, error)
}

type tagImp struct {
	models.Tag
	DB *gorm.DB
}

/*
GetTag searches for a matching tag name and creates the tag
if it doesn't exist
*/
func GetTag(name string, db *gorm.DB) (Tag, error) {
	t := tagImp{
		DB: db,
		Tag: models.Tag{
			Name: name,
		},
	}
	result := t.DB.Where(&t, "name").First(&t)
	if result.Error != nil {
		return t, errors.New("cant create new tag")
	}
	return t, nil

}

func (t tagImp) GetID() uint {
	return t.ID
}

func (t tagImp) GetName() string {
	return t.Name
}

func (t tagImp) GetTags() ([]Tag, error) {
	panic("not implemented") // TODO: Implement
}

func (t tagImp) Create() error {
	panic("not implemented") // TODO: Implement
}
