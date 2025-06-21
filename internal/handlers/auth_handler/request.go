package auth_handler

type TokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
