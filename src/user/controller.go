package user

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	// do something
}
