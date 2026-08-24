package user

import (
	"errors"
)

var ErrFirstNameRequired = errors.New("First name is required")
var ErrLastNameRequired = errors.New("Last name is required")

var ErrUserNotFound = errors.New("User not found")
