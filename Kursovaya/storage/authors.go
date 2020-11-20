package storage

import (
	"Kursovaya/model"
	"context"
	"fmt"
)

func (p *Provider) AddAuthorDB (author model.Author) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultCtxTimeout)
	defer cancel()

	tx, err := p.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	var q string

	q = `INSERT INTO kursovaya.authors (first_name, last_name, middle_name)
			VALUES ($1, $2, $3)`

	_, err = tx.Exec(q, author.FirstName, author.LastName, author.MiddleName)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (p *Provider) GetAuthorsDB() ([]model.Author, error) {
	var (
		q string
		err error
		authors = make([]model.Author, 0)
	)

	q = `SELECT author_id, first_name, last_name, middle_name FROM kursovaya.authors`

	rows, err := p.conn.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var author model.Author

		if err := rows.Scan(
			&author.Id,
			&author.FirstName,
			&author.LastName,
			&author.MiddleName,
			); err != nil {
			return nil, err
		}

		authors = append(authors, author)
	}

	return authors, rows.Err()
}

func (p *Provider) DeleteAuthorDB(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultCtxTimeout)
	defer cancel()

	tx, err := p.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	var q string
	var authorId int

	q = `SELECT author_id FROM kursovaya.authors WHERE author_id = $1`

	if err = tx.QueryRow(q, id).Scan(&authorId); err != nil {
		return fmt.Errorf("автора не существует. Id автора = %d", id)
	}

	q = `DELETE FROM kursovaya.authors WHERE author_id = $1`

	_, err = tx.Exec(q, id)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (p *Provider)UpdateAuthorDB(id int, author model.Author) error  {
	ctx, cancel := context.WithTimeout(context.Background(), defaultCtxTimeout)
	defer cancel()

	tx, err := p.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	var q string
	var authorId int

	q = `SELECT author_id FROM kursovaya.authors WHERE author_id = $1`

	if err = tx.QueryRow(q, id).Scan(&authorId); err != nil {
		return fmt.Errorf("автора не существует. Id автора = %d", id)
	}

	q = `UPDATE kursovaya.authors SET first_name = $1, last_name = $2, middle_name = $3 WHERE author_id = $4`

	if _, err := tx.Exec(q, author.FirstName, author.LastName, author.MiddleName, id); err != nil {
		return err
	}

	return tx.Commit()
}
