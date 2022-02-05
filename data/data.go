package database

import (
	"errors"
	"notes-app/data/models"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DataBase *gorm.DB

/*
InitializeLocalDB Attempts to open an existing database at the current path
and fails over to
*/
func InitializeLocalDB() (*gorm.DB, error) {
	//check working directory for db, fail over to home directory
	file_name := "./.backpack.db"
	if _, err := os.Stat(file_name); err != nil {
		file_name = "~/.notes/.backpack.db"
	}
	// open and initialize db schema
	db, err := gorm.Open(sqlite.Open(file_name), &gorm.Config{})
	if err != nil {
		return nil, errors.New("could not create backback")
	}
	db.AutoMigrate(&models.NoteBook{}, models.Note{}, models.Tag{})
	return db, nil
}

func GetDataBase() (*gorm.DB, error) {
	var err error
	if DataBase == nil {
		DataBase, err = InitializeLocalDB()
		if err != nil {
			return nil, err
		}
	}
	return DataBase, nil
}
