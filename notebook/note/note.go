package note

import (
	models "notes-app/data"
	"strings"

	"gorm.io/gorm"
)

type Note interface {
	CreateNote() Note
	AddTag(t Tag)
	GetTags() []Tag
	GetPath() string
}

type NoteImp struct {
	Model models.Note
	DB    *gorm.DB
}

func InitNoteImp(title string, db *gorm.DB) Note {
	return NoteImp{
		DB: db,
		Model: models.Note{
			Title: title,
		},
	}
}

func (n NoteImp) CreateNote() Note {
	panic("not implemented") // TODO: Implement
}

func (n NoteImp) AddTag(t Tag) {
	panic("not implemented") // TODO: Implement
}

func (n NoteImp) GetTags() []Tag {
	panic("not implemented") // TODO: Implement
}

func (n NoteImp) GetPath() string {
	dateString := n.Model.CreatedAt.String()
	dateString = strings.ReplaceAll(dateString, " ", "_")
	return n.Model.Book.Storage.LocalPath + dateString + ".md"
}

type NoteParser interface {
	Parse(body string) (string, []Tag, error)
}

type NoteFormatter interface {
	Format(body string) (string, error)
}
