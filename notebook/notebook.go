package notebook

import (
	"errors"
	"fmt"
	models "notes-app/data"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Tag interface {
	//Name string
}
type Note interface {
	//AddTag(t Tag)
	//GetTags() []Tag
	//GetPath() string
}

type NoteBook interface {
	CreateNote(title string) (Note, error)
	GetName() string
	//FindNote(title string) (Note, error)
	//EditNote(id int) (Note, error)
	//GetAllNotes() []Note
}

type BackPack interface {
	GetAllNoteBooks() ([]NoteBook, error)
	CreateNotebook(name string) (NoteBook, error)
}

type GormBackpack struct {
	DB *gorm.DB
}

func InitializeBackpack() (BackPack, error) {
	db, err := gorm.Open(sqlite.Open("notes.db"), &gorm.Config{})
	if err != nil {
		return nil, errors.New("could not create backback")
	}
	db.AutoMigrate(&models.NoteBook{}, models.Note{}, models.Tag{})
	return GormBackpack{DB: db}, nil
}

type GormNoteBook struct {
	DB       *gorm.DB
	Notebook models.NoteBook
}

func InitializeNotebook(db *gorm.DB) (GormNoteBook, error) {

	return GormNoteBook{DB: db}, nil
}
func (g GormBackpack) GetAllNoteBooks() ([]NoteBook, error) {
	var noteBooks []models.NoteBook
	err := g.DB.Find(&noteBooks).Error

	if err != nil {
		return nil, errors.New("error finding notebooks")
	}
	output := make([]NoteBook, len(noteBooks))
	for index, element := range noteBooks {
		output[index] = GormNoteBook{DB: g.DB, Notebook: element}
	}

	return output, nil
}

func (g GormBackpack) CreateNotebook(name string) (NoteBook, error) {
	notebook := models.NoteBook{
		Name: name,
	}
	if err := g.DB.Create(&notebook).Error; err != nil {
		return nil, errors.New("could not create notebook")
	}
	fmt.Println("Notebook ID: ", notebook.ID)
	return GormNoteBook{DB: g.DB, Notebook: notebook}, nil
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

// func (g *GormBackpack) GetAllNoteBooks() ([]NoteBook, error){
// 	var books []NoteBook
// 	g.DB.

// }
