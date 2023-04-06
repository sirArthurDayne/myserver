package myserver

import (
	"encoding/json"
	"net/http"
)

const (
	GET  = http.MethodGet
	POST = http.MethodPost
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type Metadata interface {
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (u *User) ToJson() ([]byte, error) {
	return json.Marshal(u)
}
