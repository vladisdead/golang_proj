package server

import (
	"Kursovaya/model"
	"encoding/json"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
)

func (s *Server) AddTranslationHandler(ctx *fasthttp.RequestCtx, log *zerolog.Logger) {
	var trasnaltion model.Translation

	if err := json.Unmarshal(ctx.Request.Body(), &trasnaltion); err != nil {
		log.Err(err).Msg("Ошибка при декодировании информации по переводу")
		replyError(ctx, log, fasthttp.StatusBadRequest, "Неверный json")

		return
	}

	if err := s.books.AddTranslation(trasnaltion); err != nil {
		log.Err(err).Msg("Ошибка при добавлении перевод")
		replyError(ctx, log, fasthttp.StatusBadRequest, "Ошибка при добавлении перевода")

		return
	}

	ctx.SetStatusCode(fasthttp.StatusCreated)
}
