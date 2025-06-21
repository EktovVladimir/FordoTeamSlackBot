package auth

import "github.com/dgrijalva/jwt-go"

type JwtClaims struct {
	Payload Payload `json:"payload"`
	jwt.StandardClaims
}

type Payload struct {
	Username string `json:"username"`
}
