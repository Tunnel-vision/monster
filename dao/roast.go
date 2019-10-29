package models

import "github.com/jinzhu/gorm"

type Roast struct {
	gorm.Model
	Content         string `gorm:"type:text" json:"content"`
	ParentRoastID uint64  `json:"parent_roast_id"` // ID of replied comment
	Uid             uint64 `sql:"index" json:"uid"`
	Likes           uint32 `json:"likes"`
}
