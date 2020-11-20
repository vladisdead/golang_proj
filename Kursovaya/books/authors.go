package books

import "Kursovaya/model"

func (p *Provider) NewAuthor(author model.Author) error {
	return p.storage.AddAuthorDB(author)
}

func (p *Provider) GetAuthors()([]model.Author, error)  {
	return p.storage.GetAuthorsDB()
}

func (p *Provider) DeleteAuthor(id int) error {
	return p.storage.DeleteAuthorDB(id)
}

func (p *Provider) UpdateAuthor(id int, author model.Author) error {
	return p.storage.UpdateAuthorDB(id,author)
}
