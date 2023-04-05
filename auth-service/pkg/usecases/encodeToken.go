package usecases

import (
	"auth-service/pkg/entities"
	"time"

	"github.com/golang-jwt/jwt"
)

func (u *usecase) EncodeToken(user entities.User) (string, error) {
	claims := &entities.JwtClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Unix() + 28800, // 8 hours
			// Issuer:    w.Issuer,
		},
		ID:         user.ID,
		Username:   user.Username,
		Permission: user.Permission,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte("abcd"))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
