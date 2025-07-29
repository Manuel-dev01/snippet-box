package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord = errors.New("models: no matchig record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail = errors.New("models: duplicatde email")
)

type Snippet struct {
	ID 	    int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type User struct {
	ID 		       int
	Name 		   string
	Email 		   string
	HashedPassword []byte
	Created 	   time.Time
}


