package storage

import (
	"Kursovaya/model"
	"context"
	"fmt"
)

func (p *Provider) AddBookDB (book model.Book) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultCtxTimeout)
	defer cancel()

	tx, err := p.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	var q string
	var authorId int

	q = "SELECT author_id from kursovaya.authors WHERE author_id = $1"

	if err = tx.QueryRow(q, book.AuthorId).Scan(&authorId); err != nil {
		return fmt.Errorf("Невозможно добавить книгу. Такого автора не существует")
	}

	q = `INSERT INTO kursovaya.books (tittle, place, edition, year, num_of_page, author_id) 
			VALUES ($1, $2, $3, $4, $5, $6)`

	_, err = tx.Exec(q, book.Tittle, book.Place, book.Edition, book.Year, book.NumberOfPage, book.AuthorId)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (p *Provider) GetBooksDB() ([]model.Book, error) {
	var (
		q string
		err error
		books = make([]model.Book, 0)
	)

	q = `SELECT book_id, tittle, place, edition, year, num_of_page, author_id FROM kursovaya.books`

	rows, err := p.conn.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book model.Book

		if err := rows.Scan(
			&book.Id,
			&book.Tittle,
			&book.Place,
			&book.Edition,
			&book.Year,
			&book.NumberOfPage,
			&book.AuthorId,
			); err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, rows.Err()
}

func (p *Provider) CheckBookByID(id int) error {
	var (
		q string
		err error
		bookId int
	)

	q = `SELECT book_id FROM kursovaya.books WHERE book_id = $1`

	if err = p.conn.QueryRow(q, id).Scan(&bookId); err != nil {
		p.log.Err(err).Msg("Error is")
		return fmt.Errorf("Такой книги нет. Id = %d", id)
	}

	return nil
}

func (p *Provider) DeleteBookDB(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultCtxTimeout)
	defer cancel()

	tx, err := p.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	var q string
	var bookId int

	q = `SELECT book_id FROM kursovaya.books WHERE book_id = $1`

	if err = tx.QueryRow(q, id).Scan(&bookId); err != nil {
		return fmt.Errorf("книги не существует. Id книги = %d", id)
	}

	q = `DELETE FROM kursovaya.books WHERE book_id = $1`

	_, err = tx.Exec(q, id)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (p *Provider) UpdateBookDB (id int, book model.Book) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultCtxTimeout)
	defer cancel()

	tx, err := p.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	var q string
	var bookId int

	q = `SELECT book_id FROM kursovaya.books WHERE book_id = $1`

	if err = tx.QueryRow(q, id).Scan(&bookId); err != nil {
		return fmt.Errorf("книги не существует. Id книги = %d", id)
	}

	q = `UPDATE kursovaya.books set tittle = $1, place = $2, edition = $3, year = $4, num_of_page = $5, author_id = $6 WHERE book_id = $7`

	if _, err := tx.Exec(q, book.Tittle, book.Place, book.Edition, book.Year, book.NumberOfPage, book.AuthorId, id); err != nil {
		return err
	}

	return tx.Commit()
}
