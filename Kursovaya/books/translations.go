package books

import (
	"Kursovaya/model"
)

func (p *Provider) AddTranslation(translation model.Translation) error {
	if err := p.storage.CheckTranslatorByID(translation.TranslatorID); err != nil {
		return err
	}
	if err := p.storage.CheckBookByID(translation.BookID); err != nil {
		return err
	}

	return p.storage.AddTranslationDB(translation)
}
