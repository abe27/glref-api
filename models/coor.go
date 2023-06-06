package models

import (
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Coor struct {
	FCSKID string `gorm:"primaryKey;column:FCSKID;type:char;size:8;not null;index;" json:"fcskid"`
	FCCODE string `gorm:"index:idx_fccode,unique" json:"fccode"`
	FCNAME string `json:"fcname,omitempty"`
}

func (Coor) TableName() string {
	return "COOR"
}

func (obj *Coor) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}
