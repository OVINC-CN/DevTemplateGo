package main

import (
	"github.com/OVINC-CN/DevTemplateGo/src/configs"
	"github.com/OVINC-CN/DevTemplateGo/src/middlewares"
	"github.com/OVINC-CN/DevTemplateGo/src/services/account"
	"github.com/OVINC-CN/DevTemplateGo/src/services/home"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func setupRouter() (engine *gin.Engine) {
	// 初始化
	if !configs.Config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	engine = gin.New()
	engine.RedirectTrailingSlash = false
	engine.Use(
		middlewares.CORS(),
		middlewares.Locale(),
		middlewares.InitLogger(),
		middlewares.RequestLogger(),
		middlewares.Timeout(),
		middlewares.Authenticate(),
	)

	// 注册校验器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("username", account.UsernameValidator)
	}

	// Home
	homeGroup := engine.Group("/")
	{
		homeGroup.GET("", home.Home)
		homeGroup.GET("/rum_config/", home.RumConfig)
	}

	// Account
	accountGroup := engine.Group("/account/")
	{
		accountGroup.POST("/sign_in/", account.Login)
		accountGroup.POST("/sign_up/", account.SignUp)
		accountGroup.POST("/sign_out/", middlewares.LoginRequired(), account.SignOut)
		accountGroup.POST("/user_info/", middlewares.LoginRequired(), account.UserInfo)
	}
	return
}
