-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS slots (
  id SERIAL PRIMARY KEY,
  desctiption text NOT NULL,
  banners integer[]
);

CREATE TABLE IF NOT EXISTS banners (
  id SERIAL PRIMARY KEY,
  desctiption text NOT NULL
);

CREATE TABLE IF NOT EXISTS groups (
  id SERIAL PRIMARY KEY,
  desctiption text NOT NULL
);

CREATE TABLE IF NOT EXISTS stats (
  id SERIAL PRIMARY KEY,
  slot_id integer REFERENCES slots (id) NOT NULL,
  banner_id integer REFERENCES banners (id) NOT NULL,
  group_id integer REFERENCES groups (id) NOT NULL,
  display integer DEFAULT 0,
  click integer DEFAULT 0
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS stats;
DROP TABLE IF EXISTS slots;
DROP TABLE IF EXISTS banners;
DROP TABLE IF EXISTS groups;
