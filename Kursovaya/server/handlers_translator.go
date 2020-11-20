package server

import (
	"Kursovaya/model"
	"encoding/json"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
	"strconv"
)

func (s *Server) AddTranslatorHandler(ctx *fasthttp.RequestCtx, log *zerolog.Logger) {
	var (
		err error
		translator model.Translator
	)

	if err = json.Unmarshal(ctx.Request.Body(), &translator); err != nil {
		log.Err(err).Msg("Ошибка при декодировании информации о переводчике")
		replyError(ctx, log, fasthttp.StatusBadRequest, "Неверный josn")
	}

	if err = s.books.AddTranslator(translator); err != nil {
		log.Err(err).Msg("Ошибка при добавлении автора")
		replyError(ctx, log, fasthttp.StatusBadRequest, "Ошибка при добавлении автора")

		return
	}

	ctx.SetStatusCode(fasthttp.StatusCreated)
}

func (s *Server) GetTranslatorsHandler(ctx *fasthttp.RequestCtx, log *zerolog.Logger) {
	translators, err := s.books.GetTranslators()
	if err != nil {
		log.Err(err).Msg("Невозможно показать всех переводчиков")
		replyError(ctx, log, fasthttp.StatusInternalServerError, "Невозможно показать всех переводчиков")

		return
	}

	body, err := json.Marshal(translators)
	if err != nil {
		log.Err(err).Msg("Ошибка при кодировании информации по авторам")
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)

		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(body)
}

func (s *Server) DeleteTranslatorHandler(ctx *fasthttp.RequestCtx, log *zerolog.Logger) {
	id, err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		log.Err(err).Msg( "Невозможно считать id переводчика")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)

		return
	}

	err = s.books.DeleteTranslator(id)
	if err != nil {
		log.Err(err).Msg("Ошибка при удалении информации о переводчика")
		replyError(ctx, log, fasthttp.StatusInternalServerError, "Автор не существует")

		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}

func (s *Server) UpdateTranslatorHandler(ctx *fasthttp.RequestCtx, log *zerolog.Logger) {
	var translator model.Translator

	id ,err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		log.Err(err).Msg("Невозможно считать id переводчика")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)

		return
	}

	if err = json.Unmarshal(ctx.Request.Body(), &translator); err != nil {
		log.Err(err).Msg("Ошибка при декодировании информации по переводчику")
		replyError(ctx, log, fasthttp.StatusBadRequest, "Невозможно обновить информацию о переводчике")

		return
	}

	if err = s.books.UpdateTranslator(id, translator); err != nil {
		log.Err(err).Msg("Ошибка обновлении информации автора")
		replyError(ctx, log, fasthttp.StatusNotFound, "Невозможно обновить информацию о переводчике")

		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}