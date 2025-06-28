-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
                       id BIGSERIAL PRIMARY KEY,
                       slack_name TEXT NOT NULL,
                       slack_id TEXT NOT NULL,
                       github_name TEXT NOT NULL,
                       email TEXT NOT NULL,
                       created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE deployments (
                             id BIGSERIAL PRIMARY KEY,
                             thread_ts TEXT NOT NULL,
                             pull_request_number TEXT NOT NULL,
                             workflow_run_id TEXT NOT NULL,
                             status TEXT NOT NULL,
                             created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                             updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE code_reviews (
                              id BIGSERIAL PRIMARY KEY,
                              thread_ts TEXT NOT NULL,
                              pull_request_number TEXT NOT NULL,
                              status TEXT NOT NULL,
                              created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE settings (
                          id BIGSERIAL PRIMARY KEY,
                          key TEXT NOT NULL,
                          value TEXT NOT NULL,
                          created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_settings_key ON settings(key);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS settings;
DROP TABLE IF EXISTS code_reviews;
DROP TABLE IF EXISTS deployments;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
