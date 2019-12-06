package helper

import (
	"fmt"
	"tbox_backend/src/entity"

	jwt "github.com/dgrijalva/jwt-go"
)

type IAuthHelper interface {
	GenerateToken(string, string) (string, error)
	ParseToken(string) (bool, string)
}

type authHelper struct {
	config entity.JWTConfig
}

func NewAuthHelper(config entity.JWTConfig) IAuthHelper {
	return &authHelper{
		config: config,
	}
}

func (ah *authHelper) GenerateToken(phoneNumber, userID string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	// Create the Claims
	claims := token.Claims.(jwt.MapClaims)

	// set some claims
	claims["phone_number"] = phoneNumber
	claims["userId"] = userID
	// Not need
	// claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	return token.SignedString([]byte(ah.config.SecretKey))
}

func (ah *authHelper) ParseToken(unparsedToken string) (bool, string) {
	token, err := jwt.Parse(unparsedToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return false, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(ah.config.SecretKey), nil
	})

	if err == nil && token.Valid {
		return true, unparsedToken
	}
	return false, ""

}
