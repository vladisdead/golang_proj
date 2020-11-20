-- +migrate Up
CREATE TABLE IF NOT EXISTS kursovaya.translators(
	translator_id   SERIAL primary key,
	first_name  TEXT    not null,
	last_name   TEXT    not null,
	middle_name TEXT    not null
);

-- +migrate Down
DROP TABLE IF EXISTS kursovaya.translators;
