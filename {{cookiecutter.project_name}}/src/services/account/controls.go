package account

import (
	"context"
	"github.com/OVINC-CN/DevTemplateGo/src/configs"
	"github.com/OVINC-CN/DevTemplateGo/src/db"
	"github.com/OVINC-CN/DevTemplateGo/src/utils"
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignIn(c *gin.Context) {
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
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": ginI18n.MustGetMessage(c, "user not exists")})
		return
	}
	// 校验密码
	passResult := user.CheckPassword(form.Password)
	// 不通过，报错
	if !passResult {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": ginI18n.MustGetMessage(c, "username or password invalid")})
		return
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 创建用户
	user := &User{Username: form.Username, NickName: form.Nickname, Enabled: true}
	err := user.SetPassword(form.Password)
	if err != nil {
		utils.ContextErrorf(c, "[SignUpFailed] %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ginI18n.MustGetMessage(c, SignUpFailed.Error())})
		return
	}
	createResult := db.DB.Create(user)
	if createResult.Error != nil {
		utils.ContextErrorf(c, "[SignUpFailed] %s", createResult.Error.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ginI18n.MustGetMessage(c, SignUpFailed.Error())})
		return
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
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"error": ginI18n.MustGetMessage(c, SessionIDNotExists.Error())})
		return
	}
	db.Redis.Del(context.Background(), sessionID)
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
