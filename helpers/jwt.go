package helpers

import (
	"learn-golang-gin/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(config.GetEnv("JWT_SECRET", "secrect_key"))

func GenetareToken(username string) string {

	exipirationTime := time.Now().Add(60 * time.Minute)

	claims := &jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(exipirationTime),
	}

	token, _ := jwt.NewWithClaims(jwt.SigningMethodES256, claims).SignedString(jwtKey)

	return token
}
