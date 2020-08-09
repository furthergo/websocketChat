package models

import (
	"crypto/sha256"
	"github.com/futhergo/websocketChat/internal/DB"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type UserEntity struct {
	Id int64 `gorm:"primary_key" json:"id"`
	Name string `form:"email" gorm:"column:username" json:"username"`
	Password string `form:"password" gorm:"column:password" json:"password"`
	Ip string `gorm:"column:ip" json:"ip"`
	CreateTime time.Time `json:"create_time"`
	ModifyTime time.Time `json:"modify_time"`
}

func (u UserEntity)Auth() (bool, error) {
	//TODO: register/Login/JWT
	var uu UserEntity
	DB.DB.Find(&uu, "username=?", u.Name)
	err := bcrypt.CompareHashAndPassword([]byte(uu.Password), []byte(u.Password))
	if err != nil {
		return false, nil
	}
	u.Password = uu.Password
	return true, nil
}

func (u UserEntity)AddAsQuery(r *http.Request) string {
	q := r.URL.Query()
	q.Add("username", u.Name)
	return q.Encode()
}

func (u UserEntity)Save() {
	DB.DB.Create(&u)
}

func (u UserEntity)GetMap() map[string]interface{} {
	res := make(map[string]interface{})
	res["id"] = u.Id
	res["username"] = u.Name
	res["create_time"] = u.CreateTime.Unix()
	res["modify_time"] = u.ModifyTime.Unix()
	res["ip"] = u.Ip
	return res
}

func (u UserEntity)Sha256() string {
	sha := sha256.New()
	sha.Write([]byte(u.Name))
	b := sha.Sum(nil)
	return string(b)
}

func (UserEntity)TableName() string {
	return "user"
}