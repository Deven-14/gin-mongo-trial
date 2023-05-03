package controllers

type UserController interface {
}

type userController struct {
}

func NewUserController() UserController {
	return &userController{}
}
