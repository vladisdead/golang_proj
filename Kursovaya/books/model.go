package books

import (
	"Kursovaya/model"
	"github.com/rs/zerolog"
)

// Provider - модель провайдера для работы с книгами.
type Provider struct {
	storage    storageInterface    // Интерфейс работы с хранилищем данных.
	log        *zerolog.Logger
}

// storageInterface - интерфейс работы с бд
type storageInterface interface {
	//Authors.
	AddAuthorDB (author model.Author) error
	GetAuthorsDB() ([]model.Author, error)
	DeleteAuthorDB(id int) error
	UpdateAuthorDB(id int, author model.Author) error

	//Books.
	AddBookDB (book model.Book) error
	GetBooksDB() ([]model.Book, error)
	DeleteBookDB(id int) error
	UpdateBookDB (id int, book model.Book) error
	CheckBookByID(id int) error

	//Editors.
	AddEditorDB(editor model.Editor) error
	GetEditorsDB() ([]model.Editor, error)
	DeleteEditorDB(id int) error
	UpdateEditorDB(id int, editor model.Editor) error
	CheckEditorByID(id int) error

	//Translators.
	AddTranslatorDB(translator model.Translator) error
	GetTranslatorsDB() ([]model.Translator, error)
	DeleteTranslatorDB(id int) error
	UpdateTranslatorDB(id int, translator model.Translator) error
	CheckTranslatorByID(id int) error

	//Responsibles.
	AddResponsibleDB(responsible model.Responsible) error

	//Translations.
	AddTranslationDB(translation model.Translation) error
}
