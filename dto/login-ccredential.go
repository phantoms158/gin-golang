package dto

type LoginCredential struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
