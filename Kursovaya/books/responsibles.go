package books

import "Kursovaya/model"

func (p *Provider) AddResponsible(responsible model.Responsible) error {
	if err := p.storage.CheckBookByID(responsible.BookID); err != nil {
		return err
	}
	if err := p.storage.CheckEditorByID(responsible.EditorID); err != nil {
		return err
	}
	return p.storage.AddResponsibleDB(responsible)
}
