package model

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Password string
type EncryptedPassword string

func (p Password) Encrypt() EncryptedPassword {
	return EncryptedPassword(hashAndSalt([]byte(p)))
}

func (encrypted EncryptedPassword) Compare(password Password) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(password))
	return err == nil
}

type User struct {
	Base
	FirstName         string            `gorm:"column:first_name"`
	LastName          string            `gorm:"column:last_name"`
	Email             string            `gorm:"column:email"`
	Phone             string            `gorm:"column:phone"`
	EncryptedPassword EncryptedPassword `gorm:"column:password"`
}

func (u User) TableName() string {
	return "users"
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
