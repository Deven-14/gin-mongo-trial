package routes

type UserRouter interface {
}

type userRouter struct {
}

func NewUserRouter() UserRouter {
	return &userRouter{}
}
