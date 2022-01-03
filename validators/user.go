package validators

type SignUpValidator struct {
	Email     string `form:"email" json:"email" binding:"required,email"`
	Password1 string `form:"password1" json:"password1" binding:"required,min=8,max=32,alphanum"`
	Password2 string `form:"password2" json:"password2" binding:"eqfield=Password1,required"`
}
