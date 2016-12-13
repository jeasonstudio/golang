package model

type User struct {
	Id       int `sql:"AUTO_INCREMENT"`
	Username string
	Password string
	Mail     string
}
