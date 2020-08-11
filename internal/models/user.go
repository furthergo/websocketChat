package models

import (
	"crypto/sha256"
	"github.com/futhergo/websocketChat/internal/DB"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserEntity struct {
	gorm.Model
	Name string `form:"email" gorm:"column:username" json:"username"`
	Password string `form:"password" gorm:"column:password" json:"password"`
	Ip string `gorm:"column:ip" json:"ip"`
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
	res["id"] = u.ID
	res["username"] = u.Name
	res["created_at"] = u.CreatedAt.Unix()
	res["updated_at"] = u.UpdatedAt.Unix()
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