package hash

import (
	"log"

	"github.com/matthewhartstonge/argon2"
)

func Generate(pwd string) (string, error) {
	argon := argon2.DefaultConfig()
	hash, err := argon.HashEncoded([]byte(pwd))
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func Compare(pwd string, hash string) bool {
	ok, err := argon2.VerifyEncoded([]byte(pwd), []byte(hash))
	if !ok || err != nil {
		log.Println(err)
		return false
	}
	return ok
}
