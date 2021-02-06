-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS slots (
  id SERIAL PRIMARY KEY,
  description text NOT NULL,
  banners integer[],
  total_displays integer DEFAULT 1
);

CREATE TABLE IF NOT EXISTS banners (
  id SERIAL PRIMARY KEY,
  description text NOT NULL
);

CREATE TABLE IF NOT EXISTS groups (
  id SERIAL PRIMARY KEY,
  description text NOT NULL
);

CREATE TABLE IF NOT EXISTS stats (
  id SERIAL PRIMARY KEY,
  slot_id integer REFERENCES slots (id) NOT NULL,
  banner_id integer REFERENCES banners (id) NOT NULL,
  group_id integer REFERENCES groups (id) NOT NULL,
  display integer DEFAULT 1,
  click integer DEFAULT 0
);

INSERT INTO slots (description)
VALUES ('header slot'), ('center slot'), ('footer slot');

INSERT INTO banners (description)
VALUES ('banner with car'), ('laptop banner'), ('funny banner'),
       ('banner for adults'), ('banner for youth');

INSERT INTO groups (description)
VALUES ('aged'), ('youth'), ('men'), ('women');

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS stats;
DROP TABLE IF EXISTS slots;
DROP TABLE IF EXISTS banners;
DROP TABLE IF EXISTS groups;
