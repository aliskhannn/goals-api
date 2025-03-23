-- +goose Up
-- +goose StatementBegin
ALTER TABLE goals
ADD COLUMN user_id INT NOT NULL DEFAULT 1,
ADD CONSTRAINT fk_goals_users FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE goals
DROP CONSTRAINT fk_goals_users,
DROP COLUMN user_id;
-- +goose StatementEnd
