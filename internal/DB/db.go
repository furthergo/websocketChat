package DB

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open("mysql", "root:Ldj@159357@(49.232.19.17:9301)/socket_chat?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}
