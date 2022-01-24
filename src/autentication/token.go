package autentication

import (
	"api/src/config"
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"
)

func CriarToken(userID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["userID"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString([]byte(config.SecretKey))
}

//ValidarToken verifica se o token passado é valido
func ValidarToken(r *http.Request) error {

	tokenString := extrairToken(r)
	token, err := jwt.Parse(tokenString, retornarChaveVerificacao)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); ok && token.Valid {
		return nil
	}
	return errors.New("token invalido ")
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func retornarChaveVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Metodo de assinatura de token é invalido %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}
