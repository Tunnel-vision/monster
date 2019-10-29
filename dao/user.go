package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Wx_id string `gorm:"size:64" json:"wx_id" `
	Wx_name string ` gorm:"size:64" json:"wx_name"`
	Wx_avator string `gorm:"size:128" json:"wx_avator" `
}