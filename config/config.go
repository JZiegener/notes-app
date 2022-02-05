package config

import (
	"errors"
	"notes-app/notebook"

	"gorm.io/gorm"
)

/*
Config Interface for setting default application values.
*/
type Config interface {
	GetDefaultStorage() (notebook.Storage, error)
	GetDefaultNoteBook() (notebook.NoteBook, error)
	//GetDefaultTemplate(notebook notebook.NoteBook) Template
}

type DatabaseConfig struct {
	DB *gorm.DB
}

func GetConfiguration(db *gorm.DB) Config {
	return DatabaseConfig{DB: db}
}

func (c DatabaseConfig) GetDefaultStorage() (notebook.Storage, error) {
	return notebook.NoOpStorage{}, nil
}

func (c DatabaseConfig) GetDefaultNoteBook() (notebook.NoteBook, error) {
	return nil, errors.New("not yet implemented")
}
