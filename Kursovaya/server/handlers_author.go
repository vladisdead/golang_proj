package server

import (
	"Kursovaya/model"
	"encoding/json"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
	"strconv"
)

func (s *Server) CreateAuthor(ctx *fasthttp.RequestCtx, log *zerolog.Logger) {
	var (
		author model.Author
		err    error
	)

	if err = json.Unmarshal(ctx.Request.Body(), &author); err != nil {
		log.Err(err).Msg("Ошибка при декодировании информации по автору")
		replyError(ctx, log, fasthttp.StatusBadRequest, "Неверный json")

		return
	}

	if err = s.books.NewAuthor(author); err != nil {
		log.Err(err).Msg("Ошибка при добавлении автора")
		replyError(ctx, log, fasthttp.StatusBadRequest, "Ошибка при добавлении автора")

		return
	}

	ctx.SetStatusCode(fasthttp.StatusCreated)
}

func (s *Server) GetAuthors(ctx *fasthttp.RequestCtx, log *zerolog.Logger) {
	authors, err := s.books.GetAuthors()
	if err != nil {
		log.Err(err).Msg("Невозможно показать всех авторов")
		replyError(ctx, log, fasthttp.StatusInternalServerError, "Невозможно показать всех авторов")

		return
	}

	body, err := json.Marshal(authors)
	if err != nil {
		log.Err(err).Msg("Ошибка при кодировании информации по авторам")
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)

		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(body)
}

func (s *Server) DeleteAuthorHandler(ctx *fasthttp.RequestCtx, log *zerolog.Logger) {
	id, err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		log.Err(err).Str("path", string(ctx.Path())).Msg("Невозможно считать id автора")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)

		return
	}

	err = s.books.DeleteAuthor(id)
	if err != nil {
		log.Err(err).Msg("Ошибка при удалении информации автора")
		replyError(ctx, log, fasthttp.StatusNotFound, "Автора не существует")

		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}

func (s *Server) UpdateAuthorHandler(ctx *fasthttp.RequestCtx, log *zerolog.Logger) {
	var author model.Author

	id, err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		log.Err(err).Str("path", string(ctx.Path())).Msg("Невозможно считать id автора")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)

		return
	}

	if err = json.Unmarshal(ctx.Request.Body(), &author); err != nil {
		log.Err(err).Msg("Ошибка при декодировании информации по автору")
		replyError(ctx, log, fasthttp.StatusBadRequest, "Неверно введен id автора")
		return
	}

	err = s.books.UpdateAuthor(id, author)
	if err != nil {
		log.Err(err).Msg("Ошибка обновлении информации автора")
		replyError(ctx, log, fasthttp.StatusNotFound, "Невозможно обновить информацию об авторе")

		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}