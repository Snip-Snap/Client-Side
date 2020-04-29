package jwtoken

import (
	"fmt"
	"time"

	"github.com/dchest/uniuri"
	"github.com/dgrijalva/jwt-go"
)

var secret = []byte(uniuri.NewLen(64))

type Claims struct {
	username string
	jwt.StandardClaims
}

func GenerateToken(username string) (string, error) {

	payload := jwt.MapClaims{}
	payload["username"] = username
	payload["exp"] = time.Now().Add(time.Minute * 15).Unix()

	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := claim.SignedString(secret)

	if err != nil {
		return "", err
	}
	return token, nil
}

func VerifyToken(tokenstring string) (string, error) {
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims["username"].(string), nil
	}
	return "", err

}
