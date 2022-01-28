package notebook

import (
	"errors"
	models "notes-app/data"

	"gorm.io/gorm"
)

type NoteBook interface {
	CreateNote(title string) (Note, error)
	GetName() string
	FindNote(title string) ([]Note, error)
	GetAllNotes() ([]Note, error)
}

type GormNoteBook struct {
	DB       *gorm.DB
	Notebook models.NoteBook
}

func InitializeNotebook(db *gorm.DB) (GormNoteBook, error) {

	return GormNoteBook{DB: db}, nil
}

func (g GormNoteBook) CreateNote(name string) (Note, error) {
	note := models.Note{
		Title: name,
	}
	if err := g.DB.Create(&note).Error; err != nil {
		return note, errors.New("error")
	}
	return note, nil
}

func (g GormNoteBook) GetName() string {
	return g.Notebook.Name
}

func (g GormNoteBook) FindNote(name string) ([]Note, error) {
	var notes []Note

	result := g.DB.Where("title like ? AND note_book_id = ?", "%"+name+"%", g.Notebook.ID).Find(&notes)
	if result.Error != nil {
		return notes, errors.New("error")
	}
	return notes, nil
}

func (g GormNoteBook) GetAllNotes() ([]Note, error) {
	var notes []Note

	result := g.DB.Find(&notes)
	if result.Error != nil {
		return notes, errors.New("error")
	}
	return notes, nil
}
