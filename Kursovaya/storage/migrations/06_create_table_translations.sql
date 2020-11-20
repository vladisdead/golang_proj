-- +migrate Up
CREATE TABLE IF NOT EXISTS kursovaya.translations(
	translations_id   SERIAL primary key,
	translator_id BIGINT not null,
	book_id BIGINT not null
);

-- +migrate Down
DROP TABLE IF EXISTS kursovaya.translations;
