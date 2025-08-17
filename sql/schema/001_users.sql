
DROP TABLE IF EXISTS users CASCADE;
-- +goose Up
CREATE TABLE users (
	id  UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	name TEXT UNIQUE NOT NULL
);

