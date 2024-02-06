package account

type loginForm struct {
	Username string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type signUpForm struct {
	Username string `form:"user" binding:"required,min=4,max=16,username"`
	Nickname string `form:"nickname" binding:"required,min=4,max=16"`
	Password string `form:"password" binding:"required,min=6,max=64"`
}
