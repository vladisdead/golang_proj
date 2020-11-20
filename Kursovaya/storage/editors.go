package storage

import (
	"Kursovaya/model"
	"context"
	"fmt"
)

func (p *Provider) AddEditorDB(editor model.Editor) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultCtxTimeout)
	defer cancel()

	tx, err := p.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	var q string

	q = `INSERT INTO kursovaya.editors (first_name, last_name, middle_name)
			VALUES($1, $2, $3)`

	_, err = tx.Exec(q, editor.FirstName, editor.LastName, editor.MiddleName)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (p *Provider) CheckEditorByID(id int) error {
	var (
		q string
		err error
		editorId int
	)

	q = `SELECT editor_id FROM kursovaya.editors WHERE editor_id = $1`

	if err = p.conn.QueryRow(q, id).Scan(&editorId); err != nil {
		return err
	}

	return nil
}

func (p *Provider) GetEditorsDB() ([]model.Editor, error) {
	var (
		q string
		err error
		editors = make([]model.Editor, 0)
	)

	q = `SELECT editor_id, first_name, last_name, middle_name FROM kursovaya.editors`

	rows, err := p.conn.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var editor model.Editor

		if err := rows.Scan(
			&editor.Id,
			&editor.FirstName,
			&editor.LastName,
			&editor.MiddleName,
			); err != nil {
			return nil, err
		}

		editors = append(editors, editor)
	}

	return editors, rows.Err()
}

func (p *Provider) DeleteEditorDB(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultCtxTimeout)
	defer cancel()

	tx, err := p.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	var q string
	var editorId int

	q = `SELECT editor_id FROM kursovaya.editors WHERE editor_id = $1`

	if err = tx.QueryRow(q, id).Scan(&editorId); err != nil {
		return fmt.Errorf("редактора не существует. Id редактора = %d", id)
	}

	q = `DELETE FROM kursovaya.editors WHERE editor_id = $1`

	_, err = tx.Exec(q, id)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (p *Provider) UpdateEditorDB(id int, editor model.Editor) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultCtxTimeout)
	defer cancel()

	tx, err := p.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	var q string
	var authorId int

	q = `SELECT editor_id FROM kursovaya.editors WHERE editor_id = $1`

	if err = tx.QueryRow(q, id).Scan(&authorId); err != nil {
		return fmt.Errorf("автора не существует. Id автора = %d", id)
	}

	q = `UPDATE kursovaya.editors SET first_name = $1, last_name = $2, middle_name = $3 WHERE editor_id = $4`

	if _, err := tx.Exec(q, editor.FirstName, editor.LastName, editor.MiddleName, id); err != nil {
		return err
	}

	return tx.Commit()
}
