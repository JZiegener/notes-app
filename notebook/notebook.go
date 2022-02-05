package notebook

import (
	"errors"
	"notes-app/data/models"
	"notes-app/notebook/note"
	"time"

	"gorm.io/gorm"
)

/*
NoteBook represents a collection of related notes.
These notes are all stored in a common repository.

*/
type NoteBook interface {
	GetCreateTime() time.Time
	GetName() string

	CreateNote(title string) (note.Note, error)
	//UpdateNote(note Note) error
	FindNote(title string) ([]note.Note, error)
	GetAllNotes() ([]note.Note, error)
}

type gormNoteBook struct {
	Model models.NoteBook
	DB    *gorm.DB
}

/*
InitializeNotebook returns an initalized notebook
Loads state information from sqlite3 database if it exists
*/
func InitializeNotebook(name string, db *gorm.DB) (gormNoteBook, error) {
	notebook := gormNoteBook{
		DB: db,
		Model: models.NoteBook{
			Name: name,
		},
	}

	result := notebook.DB.Where("name = ?", name).FirstOrCreate(&notebook.Model)
	if result.Error != nil {
		return notebook, errors.New("error searching for notebook")
	}
	if result.RowsAffected == 0 {
		if err := notebook.DB.Create(&notebook.Model).Error; err != nil {
			return notebook, errors.New("could not create notebook")
		}
	}

	return notebook, nil
}

func (g gormNoteBook) GetName() string {
	return g.Model.Name
}

func (g gormNoteBook) GetCreateTime() time.Time {
	return g.Model.CreatedAt
}

func (g gormNoteBook) CreateNote(name string) (note.Note, error) {
	note := note.CreateNote(name, g.DB)
	return note, nil
}

func (g gormNoteBook) FindNote(name string) ([]note.Note, error) {
	var notes []note.Note

	result := g.DB.Where("title like ? AND note_book_id = ?", "%"+name+"%", g.Model.ID).Find(&notes)
	if result.Error != nil {
		return notes, errors.New("error")
	}
	return notes, nil
}

func (g gormNoteBook) GetAllNotes() ([]note.Note, error) {
	var notes []note.Note

	result := g.DB.Find(&notes)
	if result.Error != nil {
		return notes, errors.New("error")
	}
	return notes, nil
}
