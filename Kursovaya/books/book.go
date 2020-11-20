package books

import "Kursovaya/model"

func (p *Provider) AddBook (book model.Book) error {
	return p.storage.AddBookDB(book)
}

func (p *Provider) GetBooks() ([]model.Book, error) {
	return p.storage.GetBooksDB()
}

func (p *Provider) DeleteBook(id int) error {
	return p.storage.DeleteBookDB(id)
}

func (p *Provider) UpdateBook(id int, book model.Book) error {
	return p.storage.UpdateBookDB(id, book)
}
