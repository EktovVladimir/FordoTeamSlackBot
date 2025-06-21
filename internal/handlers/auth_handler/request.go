package auth_handler

// TokenRequest model
// @Description запрос на генерацию токена
type TokenRequest struct {
	// @Description Имя пользователя
	Username string `json:"username"`
	// @Description Пароль
	Password string `json:"password"`
}
