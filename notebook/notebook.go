package notebook

import (
	"errors"
	models "notes-app/data"
	"notes-app/notebook/note"
	"time"

	"gorm.io/gorm"
)

/*
notebook implementation composed of
	editor hooks
	Formatter hooks
	Parser hooks


	database hooks
	git hooks


*/

type NoteBook interface {
	GetCreateTime() time.Time
	GetName() string

	CreateNote(title string) (note.Note, error)
	//UpdateNote(note Note) error
	FindNote(title string) ([]note.Note, error)
	GetAllNotes() ([]note.Note, error)
}

type NoteBookData interface {
	CreateNote(note note.Note) error
	UpdateNote(note note.Note) error

	FindNote(title string) ([]note.Note, error)
	GetAllNotes() ([]note.Note, error)
}

type GormNoteBook struct {
	DB       *gorm.DB
	Notebook models.NoteBook
}

func InitializeNotebook(db *gorm.DB) (GormNoteBook, error) {

	return GormNoteBook{DB: db}, nil
}

func (g GormNoteBook) GetName() string {
	return g.Notebook.Name
}

func (g GormNoteBook) GetCreateTime() time.Time {
	return g.Notebook.Model.CreatedAt
}

func (g GormNoteBook) CreateNote(name string) (note.Note, error) {
	note := note.InitNoteImp(name, g.DB)
	if err := g.DB.Create(&note).Error; err != nil {
		return note, errors.New("error")
	}
	return note, nil
}

func (g GormNoteBook) FindNote(name string) ([]note.Note, error) {
	var notes []note.Note

	result := g.DB.Where("title like ? AND note_book_id = ?", "%"+name+"%", g.Notebook.ID).Find(&notes)
	if result.Error != nil {
		return notes, errors.New("error")
	}
	return notes, nil
}

func (g GormNoteBook) GetAllNotes() ([]note.Note, error) {
	var notes []note.Note

	result := g.DB.Find(&notes)
	if result.Error != nil {
		return notes, errors.New("error")
	}
	return notes, nil
}
