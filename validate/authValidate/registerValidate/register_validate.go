package registerValidate

type RegisterValidate struct {
	Username        string `validate:"required,gte=3" label:"Username"`
	FullName        string `validate:"required,gte=3" label:"Full name"`
	Email           string `validate:"required,email" label:"Email"`
	Password        string `validate:"required,gte=4" label:"Password"`
	ConfirmPassword string `validate:"required,gte=4,eqfield=Password" label:"Confirm Password"`
}
