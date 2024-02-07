package account

import (
	"context"
	"github.com/OVINC-CN/DevTemplateGo/src/configs"
	"github.com/OVINC-CN/DevTemplateGo/src/db"
	"github.com/OVINC-CN/DevTemplateGo/src/utils"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Username string `json:"username" gorm:"primaryKey"`
	NickName string `json:"nick_name"`
	Password string `json:"password"`
	JoinAt   int64  `json:"join_at" gorm:"autoCreateTime:milli"`
	Enabled  bool   `json:"enabled"`
}

func (user *User) SetPassword(password string) (err error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		utils.Logger.Errorf("[EncryptPasswordFailed] %s", err)
		return
	}
	db.DB.Model(user).Updates(User{Password: string(encryptedPassword)})
	return
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

func (user *User) CreateSessionID() (sessionID string) {
	sessionID = utils.GenerateUniqID()
	db.DB.Create(&UserSession{
		Username:  user.Username,
		SessionID: sessionID,
		ExpiredAt: time.Now().Add(time.Duration(configs.Config.SessionCookieAge) * time.Second).UnixMilli(),
	})
	db.Redis.Set(
		context.Background(),
		"sessionID",
		sessionID,
		user.Username,
		time.Duration(configs.Config.SessionCookieAge)*time.Second,
	)
	return
}

func (user *User) LoadUserBySessionID(sessionID string) {
	result := db.Redis.Get(context.Background(), "sessionID", sessionID)
	user.Username = result.Val()
	db.DB.First(user)
}

type UserSession struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Username  string `json:"username" gorm:"index"`
	SessionID string `json:"session_id" gorm:"index"`
	ExpiredAt int64  `json:"expired_at"`
	CreatedAt int64  `json:"created_at" gorm:"autoCreateTime:milli"`
}
