-- +goose Up
-- +goose StatementBegin
CREATE TABLE knowledge
(
    id      SERIAL  PRIMARY KEY,
    user_id INTEGER NOT NULL,
    topic   INTEGER NOT NULL,
    text    TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE knowledge;
-- +goose StatementEnd