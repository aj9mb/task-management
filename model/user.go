package model

type User struct {
	Id                int64   `json:"id,omitempty"`
	UserName          string  `json:"userName,omitempty"`
	Password          string  `json:"password,omitempty"`
	Name              string  `json:"name,omitempty"`
	BasicAuthUserName *string `json:"basicAuthUserName,omitempty"`
	BasicAuthPwd      *string `json:"basicAuthPwd,omitempty"`
}
