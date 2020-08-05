package models

import "net/http"

type UserEntity struct {
	Name string `form:"email"`
	Password string `form:"password""`
}

func (u UserEntity)Auth() (bool, error) {
	//TODO: register/Login/JWT
	if u.Name == "test" && u.Password == "test" {
		return true, nil
	}
	return true, nil
}

func (u UserEntity)AddAsQuery(r *http.Request) string {
	q := r.URL.Query()
	q.Add("username", u.Name)
	return q.Encode()
}