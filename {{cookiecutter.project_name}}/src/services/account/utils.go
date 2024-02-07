package account

import (
	"context"
	"github.com/OVINC-CN/DevTemplateGo/src/db"
	"github.com/gin-gonic/gin"
)

func GetContextUser(c *gin.Context) *User {
	if val, ok := c.Get("User"); ok {
		if user, ok := val.(*User); ok {
			return user
		}
	}
	return &User{}
}

func LoadUserBySessionID(sessionID string) *User {
	result := db.Redis.Get(context.Background(), "sessionID", sessionID)
	user := &User{Username: result.Val()}
	dbResult := db.DB.First(user)
	if dbResult.Error == nil {
		return user
	}
	return &User{}
}
