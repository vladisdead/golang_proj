package server

import (
	"github.com/fasthttp/router"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
	"time"
)

func (s *Server) newRouter() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "Accept,Origin,Accept-Encoding,DNT,User-Agent,Content-Type")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,PUT,DELETE,PATCH")

		var (
			path = string(ctx.Path())
			method = string(ctx.Method())
			start = time.Now()
		)

		handlerLogger := s.log.With().Time("received_time", start).
			Str("method", method).
			Str("url", path).
			Str("agent", string(ctx.Request.Header.UserAgent())).
			Str("server_ip", ctx.LocalAddr().String()).
			Logger()

		r := router.New()
		s.newRouterAPI(r, &handlerLogger)

		r.Handler(ctx)
	}
}

func (s *Server) newRouterAPI(r *router.Router, log *zerolog.Logger) {
	api := r.Group("/api")
	v1 := api.Group("/v1")

	//Authors.
	v1.GET("/authors", fasthttp.CompressHandler(func(ctx *fasthttp.RequestCtx) {
		s.GetAuthors(ctx, log)
	}))
	v1.POST("/author", fasthttp.CompressHandler(func(ctx *fasthttp.RequestCtx) {
		s.CreateAuthor(ctx, log)
	}))
	v1.DELETE("/author/{id}", fasthttp.CompressHandler(func(ctx *fasthttp.RequestCtx) {
		s.DeleteAuthorHandler(ctx, log)
	}))
	v1.PUT("/author/{id}", fasthttp.CompressHandler(func(ctx *fasthttp.RequestCtx) {
		s.UpdateAuthorHandler(ctx, log)
	}))

	//Books.
	v1.POST("/book", fasthttp.CompressHandler(func(ctx *fasthttp.RequestCtx) {
		s.AddNewBookHandler(ctx, log)
	}))
	v1.GET("/books", fasthttp.CompressHandler(func(ctx *fasthttp.RequestCtx) {
		s.GetBooksHandler(ctx, log)
	}))
	v1.DELETE("/book/{id}", fasthttp.CompressHandler(func(ctx *fasthttp.RequestCtx) {
		s.DeleteBookHandler(ctx, log)
	}))
	v1.PUT("/book/{id}", fasthttp.CompressHandler(func(ctx *fasthttp.RequestCtx) {
		s.UpdateBookHandler(ctx, log)
	}))

	//Editor.
	v1.POST("/editor", fasthttp.CompressHandler(func(ctx *fasthttp.RequestCtx) {
		s.AddEditorHandler(ctx, log)
	}))
	v1.GET("/editors", fasthttp.CompressHandler(func(ctx *fasthttp.RequestCtx) {
		s.GetEditorsHandler(ctx, log)
	}))
	v1.DELETE("/editor/{id}", fasthttp.CompressHandler(func(ctx *fasthttp.RequestCtx) {
		s.DeleteEditorHandler(ctx, log)
	}))
	v1.PUT("/editor/{id}", fasthttp.CompressHandler(func(ctx *fasthttp.RequestCtx) {
		s.UpdateEditorHandler(ctx, log)
	}))

	//Translator.
	v1.POST("/translator", fasthttp.CompressHandler(func(ctx *fasthttp.RequestCtx) {
		s.AddTranslatorHandler(ctx, log)
	}))
	v1.GET("/translators", fasthttp.CompressHandler(func(ctx *fasthttp.RequestCtx) {
		s.GetTranslatorsHandler(ctx, log)
	}))
	v1.DELETE("/translator/{id}", fasthttp.CompressHandler(func(ctx *fasthttp.RequestCtx) {
		s.DeleteTranslatorHandler(ctx, log)
	}))
	v1.PUT("/translator/{id}", fasthttp.CompressHandler(func(ctx *fasthttp.RequestCtx) {
		s.UpdateTranslatorHandler(ctx, log)
	}))

	//Responsible.
	v1.POST("/responsible", fasthttp.CompressHandler(func(ctx *fasthttp.RequestCtx) {
		s.AddResponsibleHandler(ctx, log)
	}))

	//Translation.
	v1.POST("/translation", fasthttp.CompressHandler(func(ctx *fasthttp.RequestCtx) {
		s.AddTranslationHandler(ctx, log)
	}))

	r.NotFound = func(ctx *fasthttp.RequestCtx) {
		ctx.Response.SetBodyString(`"` + string(ctx.Path()) + `"`)
		ctx.Response.SetStatusCode(fasthttp.StatusMethodNotAllowed)
	}
}
