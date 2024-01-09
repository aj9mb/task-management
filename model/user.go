package model

type User struct {
	Id       int64  `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
