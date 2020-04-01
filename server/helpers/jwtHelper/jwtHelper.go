package jwtHelper

import (
	"time"

	"github.com/Biubiubiuuuu/warehouse/server/helpers/configHelper"
	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(configHelper.JwtSecret)

type Claims struct {
	UserName string
	Password string
	jwt.StandardClaims
}

// JWT encryption, generate and return token
// param username
// param password
// return string,error
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "golang",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// JWT parses the token and verifies
// param tokenStr
// retrun *Claims, error
func ParseToken(tokenStr string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
