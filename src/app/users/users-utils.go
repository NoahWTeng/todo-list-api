package users

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func passwordHashing(password *string) {
	passBytes := []byte(*password)

	hash, err := bcrypt.GenerateFromPassword(passBytes, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	*password = string(hash)
}

func comparePasswords(hashedPwd string, password string) bool {
	byteHash := []byte(hashedPwd)
	plainPwd := []byte(password)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
