-- +goose Up
CREATE TABLE users (
    user_id VARCHAR(36) PRIMARY KEY,
    username VARCHAR(36) NOT NULL,
    team_name VARCHAR(36) REFERENCES teams(team_name), 
    is_active BOOLEAN NOT NULL
);
-- +goose Down