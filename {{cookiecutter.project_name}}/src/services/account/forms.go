package account

type loginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type signUpForm struct {
	Username string `json:"username" binding:"required,min=4,max=16,username"`
	Nickname string `json:"nickname" binding:"required,min=4,max=16"`
	Password string `json:"password" binding:"required,min=6,max=64"`
}
