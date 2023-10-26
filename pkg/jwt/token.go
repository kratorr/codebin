package jwt

import (
	"codebin/internal/model"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

type UserClaims struct {
	jwt.StandardClaims `json:"jwt_._standard_claims"`
	Username           string `json:"username,omitempty"`
	UserID             int64  `json:"user_id,omitempty"`
	Email              string `json:"email,omitempty"`
}

func GenerateToken(user *model.User, secretKey []byte, duration time.Duration) (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
		Username: user.Username,
		Email:    user.Email,
		UserID:   user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func VerifyToken(tokenStr string, secretKey []byte) (*UserClaims, error) {

	tok, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("unexpected token signing method")
		}
		return secretKey, nil
	})

	fmt.Println(tok, tok.Valid, err)
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.Errorf("unexpected token signing method")
			}

			return secretKey, nil
		},
	)
	if err != nil {
		return nil, errors.Errorf("invalid token: %s", err.Error())
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, errors.Errorf("invalid token claims")
	}

	return claims, nil
}
