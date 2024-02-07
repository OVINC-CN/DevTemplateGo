package account

import (
	"context"
	"github.com/OVINC-CN/DevTemplateGo/src/configs"
	"github.com/OVINC-CN/DevTemplateGo/src/core"
	"github.com/OVINC-CN/DevTemplateGo/src/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignIn(c *gin.Context) {
	// 验证请求
	var form loginForm
	if err := c.ShouldBind(&form); err != nil {
		panic(core.NewError(http.StatusBadRequest, err.Error(), nil))
	}
	// 获取用户
	user := User{Username: form.Username}
	result := db.DB.First(&user)
	if result.RowsAffected == 0 {
		panic(UserNotExist)
	}
	// 校验密码
	passResult := user.CheckPassword(form.Password)
	// 不通过，报错
	if !passResult {
		panic(SignInFailed)
	}
	// 通过，发放令牌
	sessionID := user.CreateSessionID()
	// 响应
	c.SetCookie(
		configs.Config.SessionCookieName,
		sessionID,
		configs.Config.SessionCookieAge,
		configs.Config.SessionCookiePath,
		configs.Config.SessionCookieDomain,
		configs.Config.SessionCookieSecure,
		configs.Config.SessionCookieHttpOnly,
	)
	c.JSON(http.StatusOK, gin.H{"data": gin.H{}})
}

func SignUp(c *gin.Context) {
	// 验证请求
	var form signUpForm
	if err := c.ShouldBind(&form); err != nil {
		panic(core.NewError(http.StatusBadRequest, err.Error(), nil))
	}
	// 创建用户
	user := &User{Username: form.Username, NickName: form.Nickname, Enabled: true}
	err := user.SetPassword(form.Password)
	if err != nil {
		panic(SignUpFailed)
	}
	createResult := db.DB.Create(user)
	if createResult.Error != nil {
		panic(SignUpFailed)
	}
	// 发放令牌
	sessionID := user.CreateSessionID()
	// 响应
	c.SetCookie(
		configs.Config.SessionCookieName,
		sessionID,
		configs.Config.SessionCookieAge,
		configs.Config.SessionCookiePath,
		configs.Config.SessionCookieDomain,
		configs.Config.SessionCookieSecure,
		configs.Config.SessionCookieHttpOnly,
	)
	c.JSON(http.StatusOK, gin.H{"data": gin.H{}})
}

func SignOut(c *gin.Context) {
	sessionID, err := c.Cookie(configs.Config.SessionCookieName)
	if err != nil {
		panic(SessionIDNotExists)

	}
	db.Redis.Del(context.Background(), "sessionID", sessionID)
	c.SetCookie(
		configs.Config.SessionCookieName,
		sessionID,
		-1,
		configs.Config.SessionCookiePath,
		configs.Config.SessionCookieDomain,
		configs.Config.SessionCookieSecure,
		configs.Config.SessionCookieHttpOnly,
	)
	c.JSON(http.StatusOK, gin.H{"data": gin.H{}})
}

func UserInfo(c *gin.Context) {
	user := GetContextUser(c)
	c.JSON(
		http.StatusOK,
		gin.H{"data": gin.H{"username": user.Username, "nick_name": user.NickName}},
	)
}
