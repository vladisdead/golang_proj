package books

import "Kursovaya/model"

func (p *Provider) AddEditor(editor model.Editor) error {
	return p.storage.AddEditorDB(editor)
}

func (p *Provider) GetEditors() ([]model.Editor, error) {
	return p.storage.GetEditorsDB()
}

func (p *Provider) DeleteEditor(id int) error {
	return p.storage.DeleteEditorDB(id)
}

func (p *Provider) UpdateEditor(id int, editor model.Editor) error {
	return p.storage.UpdateEditorDB(id ,editor)
}
