package books

import (
	"github.com/rs/zerolog"
)

func NewProvider(
	storageProvider storageInterface, log *zerolog.Logger) *Provider {
	p := Provider{
		storage:    storageProvider,
		log: log,
	}
	return &p
}