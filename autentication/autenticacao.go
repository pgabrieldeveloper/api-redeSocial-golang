package autentication

import (
	"api/src/config"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

func CriarToken(userID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["userID"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodES256, permissoes)
	return token.SignedString([]byte(config.SecretKey))
}
