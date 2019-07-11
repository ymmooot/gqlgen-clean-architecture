
-- +migrate Up
CREATE TABLE IF NOT EXISTS articles (id int, title varchar(255));

-- +migrate Down
DROP TABLE IF EXISTS articles;