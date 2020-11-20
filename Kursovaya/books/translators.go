package books

import "Kursovaya/model"

func (p *Provider) AddTranslator(translator model.Translator) error {
	return p.storage.AddTranslatorDB(translator)
}

func (p *Provider) GetTranslators() ([]model.Translator, error) {
	return p.storage.GetTranslatorsDB()
}

func (p *Provider) DeleteTranslator(id int) error {
	return p.storage.DeleteTranslatorDB(id)
}

func (p *Provider) UpdateTranslator(id int, translator model.Translator) error {
	return p.storage.UpdateTranslatorDB(id, translator)
}