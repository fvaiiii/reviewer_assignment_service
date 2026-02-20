package models

type APIError struct {
	Code    string
	Message string
}

type ErrorResponse struct {
	Error APIError
}
