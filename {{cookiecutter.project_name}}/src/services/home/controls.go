package home

import (
	"github.com/OVINC-CN/DevTemplateGo/src/configs"
	"github.com/OVINC-CN/DevTemplateGo/src/services/account"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context) {
	user := account.GetContextUser(c)
	c.JSON(
		http.StatusOK,
		gin.H{
			"data": gin.H{
				"username": user.Username,
				"nickname": user.NickName,
			},
		},
	)
}

func RumConfig(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"data": gin.H{
				"id":               configs.Config.RUMID,
				"reportApiSpeed":   true,
				"reportAssetSpeed": true,
				"spa":              true,
				"hostUrl":          configs.Config.RUMHost,
			},
		},
	)
}
