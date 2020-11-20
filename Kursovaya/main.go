package main

import (
	"Kursovaya/books"
	"Kursovaya/cfg"
	"Kursovaya/server"
	"Kursovaya/storage"
	"github.com/rs/zerolog"
)

func main() {
	log := server.NewLogger()

	CFG, err := cfg.New()
	if err != nil {
		log.Err(err)
	}
	InitApi(CFG, log)
}

func InitApi(cfg *cfg.CFG, log *zerolog.Logger) {
	storageProv, err := storage.NewProvider(cfg.Storage.Connstring, cfg.Storage.MigrationsPath, log)
	if err != nil {
		log.Fatal().Err(err).Msg("ошибка при инициализации storage")
	}

	booksProv := books.NewProvider(storageProv, log)

	api := server.New(log, booksProv)
	api.Start()
}
