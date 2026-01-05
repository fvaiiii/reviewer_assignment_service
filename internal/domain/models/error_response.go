package models

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error APIError `json:"error"`
}

const (
	ErrCodeTeamExists  = "TEAM_EXISTS"
	ErrCodePRExists    = "PR_EXISTS"
	ErrCodePRMerged    = "PR_MERGED"
	ErrCodeNotAssigned = "NOT_ASSIGNED"
	ErrCodeNoCandidate = "NO_CANDIDATE"
	ErrCodeNotFound    = "NOT_FOUND"
)
