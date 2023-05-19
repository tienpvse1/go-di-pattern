-- migrate:up
CREATE TABLE IF NOT EXISTS authors (
  id   INT PRIMARY KEY,
  name text      NOT NULL,
  bio  text
);

-- migrate:down
DROP TABLE IF EXISTS authors

