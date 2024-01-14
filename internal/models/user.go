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

// func (u User) getAllInfo() string {
// 	return fmt.Sprintf("user name: %s, age: %d , many: %d", u.Name, u.Age, u.Maney)
// }

// func (u *User) setUser(user User) {
// 	u.Age = user.Age
// }
