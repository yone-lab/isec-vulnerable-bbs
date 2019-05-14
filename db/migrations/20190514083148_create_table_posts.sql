
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE posts (
    id INTEGER NOT NULL,
    uid VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL,
    content VARCHAR(4095) NOT NULL,
    PRIMARY KEY(id),
    FOREIGN KEY(uid) REFERENCES users(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE IF EXISTS posts;

