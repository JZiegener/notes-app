package models

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Title      string
	Tags       []Tag `gorm:"many2many:NoteTags;"`
	NoteBookID int
	Book       NoteBook `gorm:"foreignKey:NoteBookID"`
}

type NoteBook struct {
	gorm.Model
	Name      string
	StorageID int
	Storage   Storage `gorm:"foreignKey:StorageID"`
}

type Tag struct {
	gorm.Model
	Name string
}

type Storage struct {
	gorm.Model
	Name      string
	Url       string
	LocalPath string
}

type Identity struct {
	gorm.Model
	Name string
}
