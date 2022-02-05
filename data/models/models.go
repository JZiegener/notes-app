package models

import (
	"gorm.io/gorm"
)

// Note provides db model for storing metadata about how to access an individual
// note
type Note struct {
	gorm.Model
	Title      string
	Tags       []Tag `gorm:"many2many:NoteTags;"`
	NoteBookID int
	Book       NoteBook `gorm:"foreignKey:NoteBookID"`
}

// NoteBook provides a db model for a collection of grouped notes
type NoteBook struct {
	gorm.Model
	Name      string `gorm:"Index:uniqueIndex"`
	StorageID int
	Storage   Storage `gorm:"foreignKey:StorageID"`
}

// Tag provides a db model for metadata that can be added to notes
type Tag struct {
	gorm.Model
	Name string `gorm:"Index:uniqueIndex"`
}

// Storage is used to hold data about where a particular note book is stored on disk
// and credentials for the remote repository
type Storage struct {
	gorm.Model
	Name      string `gorm:"Index:uniqueIndex"`
	URL       string
	LocalPath string
}

// Identity is used to represent what credentials should be used for a
// Storage given
type Identity struct {
	gorm.Model
	Name string
}

// Template used
type Template struct {
	gorm.Model
	Name string
}
