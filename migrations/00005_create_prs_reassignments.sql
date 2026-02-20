-- +goose Up
CREATE TABLE pull_request_reassignments (
    id SERIAL PRIMARY KEY,
    pull_request_id VARCHAR(36) REFERENCES pull_requests(pull_request_id),
    old_reviewer_id VARCHAR(36),
    new_reviewer_id VARCHAR(36),
    changed_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose Down