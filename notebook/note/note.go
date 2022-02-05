package note

import (
	"notes-app/data/models"
	"notes-app/editor"
	"strings"
	"time"

	"gorm.io/gorm"
)

/*
Note represents a single note composed of a file and associated tags
*/
type Note interface {
	//AddTag(t Tag)
	GetTags() []Tag
	GetPath() string
}

type noteImp struct {
	models.Note
	DB *gorm.DB
}

// CreateNote Launches the editor to get input, parses the file to extract metadata
// then persists the metadata to db, and archives document content to git
func CreateNote(title string, db *gorm.DB) Note {
	note := noteImp{
		DB: db,
		Note: models.Note{
			Model: gorm.Model{
				CreatedAt: time.Now(),
			},
			Title: title,
		},
	}
	//launch the Editor to fill in the note
	editor.EditFile(note.GetPath())
	// Run the Parser to extract tags / references meta data

	// Persist data
	if err := note.DB.Create(&note).Error; err != nil {
		//TODO: add logging
		return nil
	}
	return note
}

func InitNote(note models.Note, db *gorm.DB) Note {
	retVal := noteImp{
		DB:   db,
		Note: note,
	}
	return retVal
}

func (n noteImp) AddTag(t Tag) {
	panic("not implemented") // TODO: Implement
}

func (n noteImp) GetTags() []Tag {
	panic("not implemented") // TODO: Implement
}

func (n noteImp) GetPath() string {
	dateString := n.Model.CreatedAt.String()
	dateString = strings.ReplaceAll(dateString, " ", "_")
	return n.Book.Storage.LocalPath + dateString + ".md"
}

// type NoteParser interface {
// 	Parse(body string) (string, []Tag, error)
// }

// type NoteFormatter interface {
// 	Format(body string) (string, error)
// }
