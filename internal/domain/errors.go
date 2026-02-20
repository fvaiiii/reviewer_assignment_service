package domain

import "errors"

var (
	ErrTeamExists  = errors.New("team exists")
	ErrPRExists    = errors.New("pr exists")
	ErrPRMerged    = errors.New("pr merged")
	ErrNotAssigned = errors.New("not assigned")
	ErrNoCandidate = errors.New("no candidate")
	ErrNotFound    = errors.New("not found")
)
