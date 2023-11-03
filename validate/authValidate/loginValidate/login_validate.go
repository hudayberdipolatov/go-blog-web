package loginValidate

type LoginValidate struct {
	Username string `validate:"required" label:"Username"`
	Password string `validate:"required" label:"Password"`
}
