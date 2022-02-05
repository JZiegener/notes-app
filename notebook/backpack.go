package notebook

import (
	"errors"
	"notes-app/data/models"

	"gorm.io/gorm"
)

/*
BackPack represents a collection of notebooks
*/
type BackPack interface {
	GetAllNoteBooks() ([]NoteBook, error)
	CreateNotebook(name string) (NoteBook, error)
	//GetNotebook(name string) (NoteBook, error)
}

type gormBackpack struct {
	DB *gorm.DB
}

func InitializeBackpack(db *gorm.DB) BackPack {
	return gormBackpack{DB: db}
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

func GetNotebook(name string, db *gorm.DB) (gormNoteBook, error) {
	notebook := gormNoteBook{
		DB: db,
		Model: models.NoteBook{
			Name: name,
		},
	}
	// check if the name already exists
	result := notebook.DB.Where("name = ?", name).FirstOrCreate(&notebook.Model)
	if result.Error != nil {
		return notebook, errors.New("error searching for notebook")
	}
	// insert it if
	if result.RowsAffected == 0 {
		if err := notebook.DB.Create(&notebook.Model).Error; err != nil {
			return notebook, errors.New("could not create notebook")
		}
	}

	return notebook, nil
}

func (g gormBackpack) CreateNotebook(name string) (NoteBook, error) {
	notebook, err := InitializeNotebook(name, g.DB)

	if err != nil {
		return nil, errors.New("error finding notebooks")
	}
	return notebook, nil
}
