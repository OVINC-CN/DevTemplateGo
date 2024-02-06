package main

import (
	"github.com/OrenZhang/GoDevTest/src/configs"
	"github.com/OrenZhang/GoDevTest/src/middlewares"
	"github.com/OrenZhang/GoDevTest/src/services/account"
	"github.com/OrenZhang/GoDevTest/src/services/home"
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
	engine.Use(middlewares.InitLogger(), middlewares.RequestLogger(), middlewares.Authenticate())
	// 注册校验器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("username", account.UsernameValidator)
	}
	// Home
	homeGroup := engine.Group("/home/")
	{
		homeGroup.GET("", home.Home)
	}
	// Account
	accountGroup := engine.Group("/account/")
	{
		accountGroup.POST("/signin/", account.Login)
		accountGroup.POST("/signup/", account.SignUp)
		accountGroup.POST("/signout/", middlewares.LoginRequired(), account.SignOut)
	}
	return
}
