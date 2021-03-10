package dao


type CreateUserRequest struct {
	Username string `validate:"required,email"`
	Password string `validate:"required"`
}

type UserView struct {
	Id uint `json:"id"`
	Username string `json:"username"`
}
