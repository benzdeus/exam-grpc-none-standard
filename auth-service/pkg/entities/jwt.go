package entities

import "github.com/golang-jwt/jwt"

type JwtClaims struct {
	jwt.StandardClaims
	ID         int64
	Username   string
	Permission string
}
