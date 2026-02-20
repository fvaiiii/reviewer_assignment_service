-- +goose Up
CREATE TABLE teams (
    team_name VARCHAR(36) PRIMARY KEY
    members JSONB NOT NULL DEFAULT '[]'::jsonb
);
-- +goose Down

