-- +goose Up
CREATE TABLE pull_requests (
    pull_request_id VARCHAR(36) PRIMARY KEY,   
	pull_request_name VARCHAR(255) NOT NULL,
	author_id VARCHAR(36) REFERENCES users(user_id),
	status VARCHAR(36) NOT NULL, 
	assigned_reviewers JSONB NOT NULL DEFAULT '[]'::jsonb,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	merged_at TIMESTAMP
);
-- +goose Down