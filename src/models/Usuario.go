package models

import "time"

//Usuario modele de um usuario na api
type Usuario struct {
	ID       uint64    `json:id,omitempty`
	Name     string    `json:name,omitempty`
	Nick     string    `json:nick,omitempty`
	Email    string    `json:email,omitempty`
	Password string    `json:password,omitempty`
	CreateAt time.Time `json:creatAt,omitempty`
}
