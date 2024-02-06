package account

import (
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
