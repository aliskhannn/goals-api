-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS goals
(
    id          SERIAL PRIMARY KEY,
    title       TEXT NOT NULL,
    description TEXT NOT NULL,
    completed   BOOLEAN   DEFAULT false,
    created_at  TIMESTAMP DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS goals
-- +goose StatementEnd
