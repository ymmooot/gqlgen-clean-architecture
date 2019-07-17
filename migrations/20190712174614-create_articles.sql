
-- +migrate Up
CREATE TABLE IF NOT EXISTS articles (
    id INT NOT NULL AUTO_INCREMENT,
    title VARCHAR(255),
    body TEXT,
    published_at DATETIME,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME,
    PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE IF EXISTS articles;