package user

import "errors"

var ErrUserNotFound error = errors.New("user not found")
var ErrUserAlreadyExist error = errors.New("user already exist")
