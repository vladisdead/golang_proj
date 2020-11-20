-- +migrate Up
CREATE SCHEMA IF NOT EXISTS kursovaya AUTHORIZATION "api-user";

-- +migrate Down
DROP SCHEMA IF EXISTS kursovaya CASCADE;