package account

import (
	"context"
	"github.com/OrenZhang/GoDevTest/src/configs"
	"github.com/OrenZhang/GoDevTest/src/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	// 验证请求
	var form loginForm
	if err := c.ShouldBind(&form); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 获取用户
	user := User{Username: form.Username}
	result := db.DB.First(&user)
	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "user not exists"})
		return
	}
	// 校验密码
	passResult := user.CheckPassword(form.Password)
	// 不通过，报错
	if !passResult {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "username or password invalid"})
		return
	}
	// 通过，发放令牌
	sessionID := user.CreateSessionID()
	// 响应
	c.SetCookie(
		configs.SessionConfig.SessionCookieName,
		sessionID,
		configs.SessionConfig.SessionCookieAge,
		configs.SessionConfig.SessionCookiePath,
		configs.SessionConfig.SessionCookieDomain,
		configs.SessionConfig.SessionCookieSecure,
		configs.SessionConfig.SessionCookieHttpOnly,
	)
	c.JSON(http.StatusOK, gin.H{"data": gin.H{}})
}

func SignUp(c *gin.Context) {
	// 验证请求
	var form signUpForm
	if err := c.ShouldBind(&form); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 创建用户
	user := &User{Username: form.Username, NickName: form.Nickname, Enabled: true}
	err := user.SetPassword(form.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createResult := db.DB.Create(user)
	if createResult.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": createResult.Error.Error()})
		return
	}
	// 发放令牌
	sessionID := user.CreateSessionID()
	// 响应
	c.SetCookie(
		configs.SessionConfig.SessionCookieName,
		sessionID,
		configs.SessionConfig.SessionCookieAge,
		configs.SessionConfig.SessionCookiePath,
		configs.SessionConfig.SessionCookieDomain,
		configs.SessionConfig.SessionCookieSecure,
		configs.SessionConfig.SessionCookieHttpOnly,
	)
	c.JSON(http.StatusOK, gin.H{"data": gin.H{}})
}

func SignOut(c *gin.Context) {
	sessionID, err := c.Cookie(configs.SessionConfig.SessionCookieName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	db.Redis.Del(context.Background(), sessionID)
	c.SetCookie(
		configs.SessionConfig.SessionCookieName,
		sessionID,
		-1,
		configs.SessionConfig.SessionCookiePath,
		configs.SessionConfig.SessionCookieDomain,
		configs.SessionConfig.SessionCookieSecure,
		configs.SessionConfig.SessionCookieHttpOnly,
	)
	c.JSON(http.StatusOK, gin.H{"data": gin.H{}})
}
