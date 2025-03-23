-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS goals
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    completed   BOOLEAN   DEFAULT false,
    created_at  TIMESTAMP DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS goals;
-- +goose StatementEnd
