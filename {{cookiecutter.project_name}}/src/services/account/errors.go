package account

import (
	"github.com/OVINC-CN/DevTemplateGo/src/core"
	"net/http"
)

var (
	SignInFailed       = core.NewError(http.StatusUnauthorized, "username or password invalid", &map[string]any{})
	SignUpFailed       = core.NewError(http.StatusBadRequest, "sign up failed", &map[string]any{})
	SessionIDNotExists = core.NewError(http.StatusBadRequest, "session id not exist", &map[string]any{})
	UserNotExist       = core.NewError(http.StatusNotFound, "user not exist", &map[string]any{})
)
