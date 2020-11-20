-- +migrate Up
CREATE TABLE IF NOT EXISTS kursovaya.books(
	book_id   SERIAL primary key,
	tittle  TEXT    not null,
	place   TEXT    not null,
	edition TEXT    not null,
	year    TEXT    not null,
	num_of_page BIGINT not null,
	author_id      BIGINT not null
);

-- +migrate Down
DROP TABLE IF EXISTS kursovaya.books;
