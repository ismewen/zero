package forms

type PasswordLoginForm struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mxmobile"`
	Password string `form:"password" json:"password" binding:"required,min=0,max=20"`
}
