package models

import (
	"time"
)

type Principal struct {
	AuthType string
	User string
	ExpiresIn *time.Time
	TokenType TokenType
}

type TokenType int

const (
	NoToken = iota
	AccessToken = iota
	RefreshToken = iota
)
