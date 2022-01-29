package notebook

import (
	"errors"
	models "notes-app/data"
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

// type NoteBookData interface {
// 	CreateNote(note note.Note) error
// 	UpdateNote(note note.Note) error

// 	FindNote(title string) ([]note.Note, error)
// 	GetAllNotes() ([]note.Note, error)
// }

type gormNoteBook struct {
	DB       *gorm.DB
	Notebook models.NoteBook
}

/*
InitializeNotebook returns an initalized notebook
Loads state information from sqlite3 database if it exists
*/
func InitializeNotebook(db *gorm.DB) (NoteBook, error) {

	return gormNoteBook{DB: db}, nil
}

func (g gormNoteBook) GetName() string {
	return g.Notebook.Name
}

func (g gormNoteBook) GetCreateTime() time.Time {
	return g.Notebook.Model.CreatedAt
}

func (g gormNoteBook) CreateNote(name string) (note.Note, error) {
	note := note.CreateNote(name, g.DB)
	return note, nil
}

func (g gormNoteBook) FindNote(name string) ([]note.Note, error) {
	var notes []note.Note

	result := g.DB.Where("title like ? AND note_book_id = ?", "%"+name+"%", g.Notebook.ID).Find(&notes)
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
