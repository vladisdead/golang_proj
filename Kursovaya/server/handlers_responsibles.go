package server

import (
	"Kursovaya/model"
	"encoding/json"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
)

func (s *Server) AddResponsibleHandler(ctx *fasthttp.RequestCtx, log *zerolog.Logger) {
	var (
		responsible model.Responsible
		err error
	)

	if err = json.Unmarshal(ctx.Request.Body(), &responsible); err != nil {
		log.Err(err).Msg("Ошибка при декодировании информации по автору")
		replyError(ctx, log, fasthttp.StatusBadRequest, "Неверный json")

		return
	}

	if err = s.books.AddResponsible(responsible); err != nil {
		log.Err(err).Msg("Ошибка при добавлении ответственного")
		replyError(ctx, log, fasthttp.StatusBadRequest, "Ошибка при добавлении ответственного")

		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}
