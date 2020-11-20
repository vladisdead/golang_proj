package storage

import (
	"Kursovaya/model"
	"context"
)

func (p *Provider) AddResponsibleDB(responsible model.Responsible) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultCtxTimeout)
	defer cancel()

	tx, err := p.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	var q string

	q = `INSERT INTO kursovaya.responsibles (editor_id, book_id)
			VALUES($1, $2)`

	_, err = tx.Exec(q, responsible.EditorID, responsible.BookID)
	if err != nil {
		return err
	}

	return tx.Commit()
}
