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
