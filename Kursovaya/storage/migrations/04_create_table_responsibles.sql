-- +migrate Up
CREATE TABLE IF NOT EXISTS kursovaya.responsibles(
	responsible_id   SERIAL primary key,
	editor_id BIGINT not null,
	book_id BIGINT not null
);

-- +migrate Down
DROP TABLE IF EXISTS kursovaya.responsibles;
