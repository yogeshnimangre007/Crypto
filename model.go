package model

import "time"

type User struct {
	Uid int 'json:"uid"'
	Username string 'json:"username"'
	FirstName string 'json:"firstname"'
	LastName string 'json:"lastname"'
	Email string 'json:"email"'
	Password string 'json:"password"'
	Assets []string 'json:"assets"'
	Token string 'json:"token"'
	isActive bool 'json:"is_active"'
	CreationDate time.Time 'json:"creation_date"'
	UpdationDate time.Time 'json:"updation_date"'

}

type ResponseResult struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}

