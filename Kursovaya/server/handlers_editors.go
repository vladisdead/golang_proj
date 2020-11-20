package server

import (
	"Kursovaya/model"
	"encoding/json"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
	"strconv"
)

func (s *Server) AddEditorHandler(ctx *fasthttp.RequestCtx, log *zerolog.Logger) {
	var (
		editor model.Editor
		err error
	)

	if err = json.Unmarshal(ctx.Request.Body(), &editor); err != nil {
		log.Err(err).Msg("Ошибка при декодировании информации по редактору")
		replyError(ctx, log, fasthttp.StatusBadRequest, "Неверно введен id редактора")
		return
	}

	if err = s.books.AddEditor(editor); err != nil {
		log.Err(err).Msg("Ошибка при добавлении редактора")
		replyError(ctx, log, fasthttp.StatusBadRequest, "Ошибка при добавлении редактора")

		return
	}

	ctx.SetStatusCode(fasthttp.StatusCreated)
}

func (s *Server) GetEditorsHandler(ctx *fasthttp.RequestCtx, log *zerolog.Logger) {
	editors, err := s.books.GetEditors()
	if err != nil {
		log.Err(err).Msg("Невозможно показать всех редакторов")
		replyError(ctx, log, fasthttp.StatusInternalServerError, "Невозможно показать всех редакторов")

		return
	}

	body, err := json.Marshal(editors)
	if err != nil {
		log.Err(err).Msg("Ошибка при кодировании информации по редакторам")
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)

		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(body)
}

func (s *Server) DeleteEditorHandler(ctx *fasthttp.RequestCtx, log *zerolog.Logger) {
	id, err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		log.Err(err).Str("path", string(ctx.Path())).Msg("Невозможно считать id редактора")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)

		return
	}

	err = s.books.DeleteEditor(id)
	if err != nil {
		log.Err(err).Msg("Ошибка при удалении информации редактора")
		replyError(ctx, log, fasthttp.StatusNotFound, "Редактора не существует")

		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}

func (s *Server) UpdateEditorHandler(ctx *fasthttp.RequestCtx, log *zerolog.Logger) {
	var editor model.Editor

	id, err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		log.Err(err).Str("path", string(ctx.Path())).Msg("Невозможно считать id редактора")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)

		return
	}

	if err = json.Unmarshal(ctx.Request.Body(), &editor); err != nil {
		log.Err(err).Msg("Ошибка при декодировании информации по редактору")
		replyError(ctx, log, fasthttp.StatusBadRequest, "Неверно введен id редактора")
		return
	}

	err = s.books.UpdateEditor(id, editor)
	if err != nil {
		log.Err(err).Msg("Ошибка обновлении информации редактора")
		replyError(ctx, log, fasthttp.StatusNotFound, "Невозможно обновить информацию об редакторе")

		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}
