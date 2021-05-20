package utils

import (
	"errors"
	"io/ioutil"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dileofrancoj/blog-app/structs"
)

func check(e error) {
	if e!=nil {
		panic(e)
	}
}

func JWT(u structs.User) (string,error) {
	keyData, err := ioutil.ReadFile("keys/private.pem")
	check(err)

	key := []byte(keyData)
	payload := jwt.MapClaims{
		"id": u.ID,
		"username" : u.Username,
		"fullname":u.Fullname,
		"exp": time.Now().Add(time.Hour*12).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,payload)
	tokenString, err := token.SignedString(key)

	if err != nil {
		return tokenString, err
	}
	return tokenString,nil

}

func ValidateJWT(token string) (*structs.Claim, bool, string, error) {
	keyData, err := ioutil.ReadFile("keys/private.pem")
	check(err)
	claims := &structs.Claim{}

	if token == "" {
		return claims,false,"", errors.New("Token invalido")
	}
	// convierto a JSON y retorno un objeto
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return keyData,nil
	})

	if err == nil {
		return claims,true, claims.Id, nil
	}

	if !tkn.Valid {
		return claims, false, "",errors.New("Token invalido")
	}
	return claims, false, "", err

}