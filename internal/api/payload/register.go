package payload

type RegisterUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
