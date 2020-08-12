package models

import (
	"fmt"
	"github.com/futhergo/websocketChat/internal/pkg/DB"
	"github.com/jinzhu/gorm"
	"time"
)

type Message struct {
	gorm.Model
	FromUid uint
	FromName string
	ToUid uint
	ToName string
	Content string
	SendTime time.Time
	SendStatus int // 0:success -1:fail 1: sending
	ExtraJson string
}

func (m Message)Save() {
	DB.DB.Create(&m)
}

func (m Message)FormatString() string {
	return fmt.Sprintf("[%v](%v): %v", m.FromName, m.SendTime.Format("15:04:05"), m.Content)
}

func AllMessages() []Message {
	var msgs []Message
	DB.DB.Find(&msgs)
	return msgs
}

func GetMessageById(id uint) (Message, error) {
	var m Message
	err := DB.DB.Where("id = ?", id).First(&m).Error
	return m, err
}

func (Message)TableName() string {
	return "messages"
}
