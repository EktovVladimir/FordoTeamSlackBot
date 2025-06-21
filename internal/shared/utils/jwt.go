package utils

import (
	"errors"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/auth"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

var (
	ErrInvalidSigningMethod = errors.New("invalid signing method")
	ErrInvalidToken         = errors.New("invalid token")
)

type JwtUtils struct {
	cfg *JwtConfig
}

type JwtConfig struct {
	Secret string
	Expiry time.Duration
	Issuer string
}

func NewJwtUtils(cfg *JwtConfig) *JwtUtils {
	return &JwtUtils{cfg: cfg}
}

func (j *JwtUtils) GenerateToken(payload auth.Payload) (string, error) {

	now := time.Now()
	expiresAt := now.Add(j.cfg.Expiry)

	claims := &auth.JwtClaims{
		Payload: payload,
		StandardClaims: jwt.StandardClaims{
			Id:        uuid.New().String(),
			ExpiresAt: expiresAt.Unix(),
			IssuedAt:  now.Unix(),
			Issuer:    j.cfg.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	bytes := []byte(j.cfg.Secret)
	return token.SignedString(bytes)
}

func (j *JwtUtils) ParseToken(tokenString string) (*auth.JwtClaims, error) {
	var claimsStruct auth.JwtClaims
	token, err := jwt.ParseWithClaims(tokenString, &claimsStruct, j.validate())

	if err != nil || !token.Valid {
		return nil, errors.Join(ErrInvalidToken, err)
	}

	if claims, ok := token.Claims.(*auth.JwtClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}

func (j *JwtUtils) validate() jwt.Keyfunc {
	return func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidSigningMethod
		}
		return []byte(j.cfg.Secret), nil
	}
}
