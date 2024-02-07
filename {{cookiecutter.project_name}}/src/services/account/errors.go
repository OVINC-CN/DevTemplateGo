package account

import (
	"github.com/OVINC-CN/DevTemplateGo/src/core"
	"net/http"
)

var (
	SignInFailed       = core.NewError(http.StatusUnauthorized, "username or password invalid", nil)
	SignUpFailed       = core.NewError(http.StatusBadRequest, "sign up failed", nil)
	SessionIDNotExists = core.NewError(http.StatusBadRequest, "session id not exist", nil)
	UserNotExist       = core.NewError(http.StatusNotFound, "user not exist", nil)
)
