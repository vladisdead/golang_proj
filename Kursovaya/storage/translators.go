package storage

import (
	"Kursovaya/model"
	"context"
	"fmt"
)

func (p *Provider) AddTranslatorDB(translator model.Translator) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultCtxTimeout)
	defer cancel()

	tx, err := p.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	var q string

	q = `INSERT INTO kursovaya.translators (first_name, last_name, middle_name)
			VALUES ($1, $2, $3)`

	_, err = tx.Exec(q, translator.FirstName, translator.LastName, translator.MiddleName)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (p *Provider) CheckTranslatorByID(id int) error {
	var (
		err error
		q string
		translatorId int
	)

	q = `SELECT translator_id FROM kursovaya.translators WHERE translator_id = $1`

	if err = p.conn.QueryRow(q, id).Scan(&translatorId); err != nil {
		return fmt.Errorf("Такого переводчика нет. Id = %d", id)
	}

	return nil
}

func (p *Provider) GetTranslatorsDB() ([]model.Translator, error) {
	var (
		q string
		err error
		traslators = make([]model.Translator, 0)
	)

	q = `SELECT translator_id, first_name, last_name, middle_name FROM kursovaya.translators`

	rows, err := p.conn.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var translator model.Translator

		if err := rows.Scan(
			&translator.Id,
			&translator.FirstName,
			&translator.LastName,
			&translator.MiddleName,
			); err != nil {
			return nil, err
		}

		traslators = append(traslators, translator)
	}

	return traslators, rows.Err()
}

func (p *Provider) DeleteTranslatorDB(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultCtxTimeout)
	defer cancel()

	tx, err := p.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	var q string
	var translatorId int

	q = `SELECT translator_id FROM kursovaya.translators WHERE translator_id = $1`

	if err = tx.QueryRow(q, id).Scan(&translatorId); err != nil {
		return fmt.Errorf("переводчика не существует. Id переводчика %d", id)
	}

	q = `DELETE FROM kursovaya.translators WHERE translator_id = $1`

	_, err = tx.Exec(q, id)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (p *Provider) UpdateTranslatorDB(id int, translator model.Translator) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultCtxTimeout)
	defer cancel()

	tx, err := p.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	var q string
	var translatorId int

	q = `SELECT translator_id FROM kursovaya.translators WHERE translator_id = $1`

	if err = tx.QueryRow(q, id).Scan(&translatorId); err != nil {
		return fmt.Errorf("переводчика не существует. Id переводчика %d", id)
	}

	q = `UPDATE kursovaya.translators SET first_name = $1, last_name = $2, middle_name = $3 WHERE translator_id = $4`

	if _, err := tx.Exec(q, translator.FirstName, translator.LastName, translator.MiddleName, id); err != nil {
		return err
	}

	return tx.Commit()
}
