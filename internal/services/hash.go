package services

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Hashing(str string) ([]byte, error) {
	str_byte := []byte(str)
	str_hashed, err := bcrypt.GenerateFromPassword(str_byte, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(str_hashed)
	fmt.Println(string(str_hashed))
	return str_hashed, err
}

func CompareHashAndPassword(str_hashed []byte, str string) error {
	str_byte := []byte(str)
	err := bcrypt.CompareHashAndPassword(str_hashed, str_byte)
	fmt.Println(err)
	return err
}
