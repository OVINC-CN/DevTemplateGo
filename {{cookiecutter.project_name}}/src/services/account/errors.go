package account

import "errors"

var (
	SignUpFailed       = errors.New("sign up failed")
	SessionIDNotExists = errors.New("session id not exist")
)
