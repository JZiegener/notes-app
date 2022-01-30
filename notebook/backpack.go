package notebook

import (
	"errors"
	models "notes-app/data"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/*
BackPack represents a collection of notebooks
*/
type BackPack interface {
	GetAllNoteBooks() ([]NoteBook, error)
	CreateNotebook(name, path string) (NoteBook, error)
}

type gormBackpack struct {
	DB *gorm.DB
}

/*
InitializeBackpack Attempts to open an existing backpack collection, or create one if none exist.
*/
func InitializeBackpack() (BackPack, error) {
	db, err := gorm.Open(sqlite.Open(".backpack.db"), &gorm.Config{})
	if err != nil {
		return nil, errors.New("could not create backback")
	}
	db.AutoMigrate(&models.NoteBook{}, models.Note{}, models.Tag{})
	return gormBackpack{DB: db}, nil
}

func (g gormBackpack) GetAllNoteBooks() ([]NoteBook, error) {
	var noteBooks []models.NoteBook
	err := g.DB.Find(&noteBooks).Error

	if err != nil {
		return nil, errors.New("error finding notebooks")
	}
	output := make([]NoteBook, len(noteBooks))
	for index, element := range noteBooks {
		output[index] = gormNoteBook{DB: g.DB, Model: element}
	}

	return output, nil
}

func (g gormBackpack) CreateNotebook(name, path string) (NoteBook, error) {
	notebook, _ := InitializeNotebook(name, g.DB)

	return notebook, nil
}
