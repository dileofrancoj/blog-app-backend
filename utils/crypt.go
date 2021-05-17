package utils

import (
	"golang.org/x/crypto/bcrypt"
)

var cost = 8

func Crypt(value string) (string,error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), cost)
	return string(bytes), err
}

func CompareHash(value string, valueToCompare string) (bool) {
	err := bcrypt.CompareHashAndPassword([]byte(valueToCompare),[]byte(value))
	if err != nil {
		return false
	}
	return true
}