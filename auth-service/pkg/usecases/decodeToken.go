package usecases

import (
	"auth-service/pkg/entities"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

func (u *usecase) DecodeToken(token string) (*entities.JwtClaims, error) {

	// decode jwt token
	jwtToken, err := jwt.ParseWithClaims(
		token,
		&entities.JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("abcd"), nil
		},
	)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	claims, ok := jwtToken.Claims.(*entities.JwtClaims)

	if !ok {
		log.Println("couldn't parse claims")
		return nil, errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("JWT is expired")
	}

	return claims, nil

}
