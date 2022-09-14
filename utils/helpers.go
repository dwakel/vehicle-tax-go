package utils

import (
	"errors"
	"github.com/google/uuid"
	"strings"
)

//todo: modify to use HMACSHA256
func HashToken() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return strings.Replace(id.String(), "-", "", -1), nil
}

func Decryption(data string, secret string) (string, error) {
	//Decode from Base64


	return "", errors.New("An error occured")
}
