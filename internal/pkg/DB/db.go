package DB

import (
	"fmt"
	"github.com/futhergo/websocketChat/internal/pkg/settings"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dbSetting := settings.DBSetting
	DB, err = gorm.Open(dbSetting.Type, fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbSetting.User,
		dbSetting.Password,
		dbSetting.Host,
		dbSetting.Name))
	if err != nil {
		panic(err)
	}
}
