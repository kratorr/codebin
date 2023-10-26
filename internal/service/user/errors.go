package user

import (
	"github.com/pkg/errors"
)

var (
	ErrUserNotActivated = errors.New("user not activated")
	ErrAlreadyExists    = errors.New("already exists")
)
