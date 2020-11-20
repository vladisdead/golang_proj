package server

import (
	"Kursovaya/model"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
)

type Server struct {
	api             fasthttp.Server
	log				*zerolog.Logger
	books			booksInterface
}


type booksInterface interface {
	//Author.
	NewAuthor (author model.Author) error
	GetAuthors()([]model.Author, error)
	DeleteAuthor(id int) error
	UpdateAuthor(id int, author model.Author) error

	//Book.
	AddBook (book model.Book) error
	GetBooks() ([]model.Book, error)
	DeleteBook(id int) error
	UpdateBook(id int, book model.Book) error

	//Editor.
	AddEditor(editor model.Editor) error
	GetEditors() ([]model.Editor, error)
	DeleteEditor(id int) error
	UpdateEditor(id int, editor model.Editor) error

	//Translator.
	AddTranslator(translator model.Translator) error
	GetTranslators() ([]model.Translator, error)
	DeleteTranslator(id int) error
	UpdateTranslator(id int, translator model.Translator) error

	//Responsible.
	AddResponsible(responsible model.Responsible) error

	//Translation.
	AddTranslation(translation model.Translation) error
}

