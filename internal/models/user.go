package models

import (
	"time"
)

type User struct {
	Id       uint
	Name     string
	Surname  string
	Birthday time.Time
	Email    string
	Phone    string
	Login    string
	Password []byte
}

type User_json struct {
	Name      string `json:"Name"`
	Surname   string `json:"Surname"`
	Email     string `json:"Email"`
	Birthday  string `json:"Birthday"`
	Phone     string `json:"Phone"`
	Password1 string `json:"Password1"`
	Password2 string `json:"Password2"`
	Login     string `json:"Login"`
}

// func (u User) getAllInfo() string {
// 	return fmt.Sprintf("user name: %s, age: %d , many: %d", u.Name, u.Age, u.Maney)
// }

// func (u *User) setUser(user User) {
// 	u.Age = user.Age
// }
