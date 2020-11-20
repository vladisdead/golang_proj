package server

import (
	"Kursovaya/model"
	"encoding/json"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
	"strconv"
)

func (s *Server) AddNewBookHandler(ctx *fasthttp.RequestCtx, log *zerolog.Logger) {
	var (
		err error
		book model.Book
	)

	if err = json.Unmarshal(ctx.Request.Body(), &book); err != nil {
		log.Err(err).Msg("Ошибка при декодировании информации по книги")
		replyError(ctx, log, fasthttp.StatusBadRequest, "Неверно введен json")
		return
	}

	if err = s.books.AddBook(book); err != nil {
		log.Err(err).Msg("Ошибка при добавлении книги")
		replyError(ctx, log, fasthttp.StatusBadRequest, "Ошибка при добавлении книги")

		return
	}

	ctx.SetStatusCode(fasthttp.StatusCreated)
}

func (s *Server) GetBooksHandler(ctx *fasthttp.RequestCtx, log *zerolog.Logger) {
	books, err := s.books.GetBooks()
	if err != nil {
		log.Err(err).Msg("Невозможно показать все книги")
		replyError(ctx, log, fasthttp.StatusInternalServerError, "Невозможно показать все книги")

		return
	}

	body, err := json.Marshal(books)
	if err != nil {
		log.Err(err).Msg("Ошибка при кодировании информации по книгам")
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)

		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(body)
}

func (s *Server) DeleteBookHandler(ctx *fasthttp.RequestCtx, log *zerolog.Logger) {
	id, err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		log.Err(err).Str("path", string(ctx.Path())).Msg("Невозможно считать id книги")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)

		return
	}

	err = s.books.DeleteBook(id)
	if err != nil {
		log.Err(err).Msg("Ошибка при удалении книги")
		replyError(ctx, log, fasthttp.StatusNotFound, "Книги не существует")

		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}

func (s *Server) UpdateBookHandler(ctx *fasthttp.RequestCtx, log *zerolog.Logger) {
	var book model.Book

	id, err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		log.Err(err).Str("path", string(ctx.Path())).Msg("Невозможно считать id книги")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)

		return
	}

	if err = json.Unmarshal(ctx.Request.Body(), &book); err != nil {
		log.Err(err).Msg("Ошибка при декодировании информации по книги")
		replyError(ctx, log, fasthttp.StatusBadRequest, "Неверно введен id книг")
		return
	}

	err = s.books.UpdateBook(id, book)
	if err != nil {
		log.Err(err).Msg("Ошибка обновлении информации книги")
		replyError(ctx, log, fasthttp.StatusNotFound, "Невозможно обновить информацию о книге")

		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}
