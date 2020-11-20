package storage

import (
	"database/sql"
	"github.com/rs/zerolog"
	"time"
)

var (
	defaultCtxTimeout = 5 * time.Second
)

// Provider - модель провайдера для работы с базой данных.
type Provider struct {
	conn         *sql.DB         // Подключение к БД.
	log			*zerolog.Logger
}
